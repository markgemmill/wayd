package database

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DatabaseSuite struct {
	suite.Suite
	Service *DatabaseService
}

func (s *DatabaseSuite) SetupTest() {
	fmt.Printf("DatabaseSuite.SetUpTest...\n")
	logger := slog.New(slog.Default().Handler())
	service, err := NewDatabaseService("", logger)
	if err != nil {
		panic("failed to initialize new db")
	}
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), service)
	s.Service = service
}

func (s *DatabaseSuite) TearDownTest() {
	fmt.Printf("DatabaseSuite.TeardownTest...\n")
	err := s.Service.Close()
	if err != nil {
		fmt.Printf("Error closing :memory: db: %s\n", err)
	}
}

func (s *DatabaseSuite) TestInitialization() {

	// initial row count
	assert.Equal(
		s.T(),
		int64(1),
		s.Service.ProjectCount(),
		fmt.Sprintf("expecting a single row in the project table, but got %d", s.Service.ProjectCount()),
	)

	assert.Equal(
		s.T(),
		int64(1),
		s.Service.EntryCount(),
		fmt.Sprintf("expecting a single row in the entry table, but got %d", s.Service.EntryCount()),
	)

	var cat *Category
	s.Service.Db.Model(&Category{}).First(cat)

	// create new project
	prj := s.Service.NewProject("Good to go!", cat)
	assert.Equal(
		s.T(),
		uint(2),
		prj.ID,
		fmt.Sprintf("expecting an intial project id of 1, but got %d", prj.ID),
	)

	entry := s.Service.NewEntry(prj)

	// create new entry
	s.Service.Db.Create(&entry)
	assert.Equal(
		s.T(),
		uint(2),
		entry.ID,
		fmt.Sprintf("expecting an intial entry id of 2, but got %d", prj.ID),
	)

	// how many rows
	assert.Equal(
		s.T(),
		int64(2),
		s.Service.ProjectCount(),
		fmt.Sprintf("expecting a two rows in the project table, but got %d", s.Service.ProjectCount()),
	)
	assert.Equal(
		s.T(),
		int64(2),
		s.Service.EntryCount(),
		fmt.Sprintf("expecting a two rows in the entry table, but got %d", s.Service.EntryCount()),
	)
}

func (s *DatabaseSuite) TestThis() {
	project := s.Service.GetProjectById(1)
	assert.Equal(s.T(), uint(1), project.ID)
	assert.Equal(s.T(), "Demo Project", project.Name)

	entry := s.Service.NewEntry(&project)

	assert.Equal(s.T(), uint(2), entry.ID)
	assert.Equal(s.T(), int64(2), s.Service.EntryCount())

}

func (s *DatabaseSuite) TestGetActiveEntry() {
	project := s.Service.GetProjectById(1)
	assert.Equal(s.T(), uint(1), project.ID)
	assert.Equal(s.T(), "Demo Project", project.Name)

	entry1 := s.Service.GetEntryById(1)
	entry2 := s.Service.NewEntry(&project)
	entry3 := s.Service.NewEntry(&project)
	entry4 := s.Service.NewEntry(&project)
	entry5 := s.Service.NewEntry(&project)

	assert.Equal(s.T(), uint(1), entry1.ID)
	assert.Equal(s.T(), uint(2), entry2.ID)
	assert.Equal(s.T(), uint(3), entry3.ID)
	assert.Equal(s.T(), uint(4), entry4.ID)
	assert.Equal(s.T(), uint(5), entry5.ID)

	assert.Equal(s.T(), int64(5), s.Service.EntryCount())
	assert.Equal(s.T(), int64(5), s.Service.ActiveEntryCount())

	entry1.StopEntry()
	entry2.StopEntry()
	entry4.StopEntry()

	s.Service.SaveEntry(entry1)
	s.Service.SaveEntry(entry2)
	s.Service.SaveEntry(entry4)

	assert.Equal(s.T(), int64(5), s.Service.EntryCount())
	assert.Equal(s.T(), int64(2), s.Service.ActiveEntryCount())

	active := s.Service.GetActiveEntry()
	assert.Equal(s.T(), uint(5), active.ID)

	entry3.StopEntry()
	entry5.StopEntry()
	s.Service.SaveEntry(entry3)
	s.Service.SaveEntry(entry5)

	assert.Equal(s.T(), int64(0), s.Service.ActiveEntryCount())

	noActive := s.Service.GetActiveEntry()
	assert.Equal(s.T(), uint(0), noActive.ID)
}

func (s *DatabaseSuite) TestStatement() {
	// assert.Equal(s.T(), "nnnnnno!", s.Service.ProjectTimes(time.Now(), time.Now()))
}

func TestDatabaseSuite(t *testing.T) {
	suite.Run(t, new(DatabaseSuite))
}
