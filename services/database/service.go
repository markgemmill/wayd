package database

import (
	"context"
	_ "embed"
	"fmt"
	"log/slog"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/wailsapp/wails/v3/pkg/application"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//go:embed resources/project_durations.sql
var projectDurationSql string

//go:embed resources/category_durations.sql
var categoryDurationSql string

type DatabaseService struct {
	location string
	log      *slog.Logger
	Db       *gorm.DB
	Ctx      context.Context
}

func NewDatabaseService(location string, logger *slog.Logger) (*DatabaseService, error) {
	service := DatabaseService{
		location: location,
		log:      logger,
		Ctx:      context.Background(),
	}
	err := service.Connect()
	if err != nil {
		return nil, err
	}

	err = service.Initialize()
	if err != nil {
		return nil, err
	}

	err = service.CreateDefaults()
	if err != nil {
		return nil, err
	}

	return &service, nil
}

func (s *DatabaseService) Name() string {
	return "DatabaseService"
}

func (s *DatabaseService) OnStartup(ctx context.Context, options application.ServiceOptions) error {
	s.log.Debug(fmt.Sprintf("DatabaseService.OnStartup... %s", options.Name))
	return nil
}

func (s *DatabaseService) OnShutdown() error {
	s.log.Debug("DatabaseService.OnShutdown...")
	return nil
}

func (dbs *DatabaseService) Close() error {
	root, err := dbs.Db.DB()
	if err != nil {
		return err
	}
	err = root.Close()
	if err != nil {
		return err
	}
	return nil
}

func (dbs *DatabaseService) Connect() error {
	db, err := gorm.Open(sqlite.Open(dbs.location), &gorm.Config{})
	if err != nil {
		return err
	}
	dbs.Db = db
	return nil
}

func (dbs *DatabaseService) Initialize() error {
	err := dbs.Db.AutoMigrate(&Project{}, &Entry{})
	if err != nil {
		return err
	}
	return nil
}

func (dbs *DatabaseService) CreateDefaults() error {
	catCount := dbs.CategoryCount()
	prjCount := dbs.ProjectCount()

	dbs.log.Debug(fmt.Sprintf("DatabaseService has %d projects\n", prjCount))

	var defaultCategory *Category
	dbs.Db.Model(&Category{}).First(defaultCategory)

	if catCount < 1 {
		dbs.log.Debug("Creating default categories...")
		defaultCategory = dbs.NewCategory("DEMO")
	}

	if prjCount < 1 {
		dbs.log.Debug("Creating default data...")
		defaultProject := dbs.NewProject("Demo Project", defaultCategory)
		dbs.NewEntry(defaultProject)
	}

	var nonCategoriedProjects []Project
	dbs.Db.Where("category_id is null").Find(&nonCategoriedProjects)
	for _, prj := range nonCategoriedProjects {
		prj.Category = defaultCategory
		dbs.Db.Save(prj)
	}

	return nil
}

func (dbs *DatabaseService) CategoryCount() int64 {
	var count int64
	dbs.Db.Model(&Category{}).Count(&count)
	return count
}

func (dbs *DatabaseService) NewCategory(name string) *Category {
	category := NewCategory(name)
	dbs.Db.Create(category)
	return category
}

func (dbs *DatabaseService) ProjectCount() int64 {
	var count int64
	dbs.Db.Model(&Project{}).Count(&count)
	return count
}

func (dbs *DatabaseService) NewProject(name string, category *Category) *Project {
	project := NewProject(name)
	project.Category = category
	dbs.Db.Create(project)
	return project
}

func (dbs *DatabaseService) GetProjectById(projectId int) Project {
	var project Project
	dbs.Db.First(&project, projectId)
	return project
}

func (dbs *DatabaseService) GetAllActiveProjects() []Project {
	var projects []Project
	dbs.Db.Where("active = ?", 1).Find(&projects)
	return projects
}

func (dbs *DatabaseService) GetAllActiveCategories() []Category {
	var categories []Category
	dbs.Db.Where("active = ?", 1).Find(&categories)
	return categories
}

func (dbs *DatabaseService) EntryCount() int64 {
	var count int64
	dbs.Db.Model(&Entry{}).Count(&count)
	return count
}

func (dbs *DatabaseService) NewEntry(project *Project) *Entry {
	entry := NewEntry(project)

	dbs.Db.Create(entry)

	return entry
}

func (dbs *DatabaseService) GetEntryById(entryId int) *Entry {
	var entry Entry
	dbs.Db.Preload(clause.Associations).First(&entry, entryId)
	return &entry
}

func (dbs *DatabaseService) GetActiveEntry() *Entry {
	var entry Entry
	qry := dbs.Db.Preload(clause.Associations).Where("end = '0001-01-01 00:00:00+00:00'")
	qry = qry.Order("created_at DESC")
	qry = qry.Limit(1)
	qry.Find(&entry)

	// dbs.log.Debug(fmt.Sprintf("DatabaseService.GetActiveEntry -> ID -> %d", entry.ID))
	// dbs.log.Debug(fmt.Sprintf("DatabaseService.GetActiveEntry -> ProjectID -> %d", entry.ProjectId))
	// dbs.log.Debug(fmt.Sprintf("DatabaseService.GetActiveEntry -> Project -> %v", entry.Project))
	// dbs.log.Debug(fmt.Sprintf("DatabaseService.GetActiveEntry -> ID -> %d", entry.ID))
	// dbs.log.Debug(fmt.Sprintf("DatabaseService.GetActiveEntry -> ID -> %d", entry.ID))

	return &entry
}

func (dbs *DatabaseService) ActiveEntryCount() int64 {
	var count int64 = 0
	qry := dbs.Db.Model(&Entry{})
	qry = qry.Where("end = '0001-01-01 00:00:00+00:00'")
	qry.Count(&count)

	return count
}

func (dbs *DatabaseService) SaveEntry(entry *Entry) {
	dbs.Db.Save(entry)
}

func (dbs *DatabaseService) StopEntry(entry *Entry) {
	entry.StopEntry()
	dbs.Db.Save(entry)
}

type ProjectDuration struct {
	Name       string
	EntryCount int64
	Seconds    int64
	Duration   string
}

func (dbs *DatabaseService) ProjectDurationTimes(groupOn string, startDate time.Time, endDate time.Time) ([]ProjectDuration, error) {
	var rows []ProjectDuration

	dbs.log.Debug(projectDurationSql)

	sDate := startDate.Format("2006-01-02 15:04:05")
	eDate := endDate.Format("2006-01-02 15:04:05")

	dbs.log.Debug(sDate)
	dbs.log.Debug(eDate)

	statement := projectDurationSql
	if groupOn == "CATEGORY" {
		statement = categoryDurationSql
	}

	qry := dbs.Db.Raw(statement, sDate, eDate)
	result := qry.Find(&rows)
	if result.Error != nil {
		return rows, result.Error
	}

	return rows, nil
}
