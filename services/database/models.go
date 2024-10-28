package database

import (
	"fmt"
	"math"
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name   string
	Active bool
}

func NewCategory(name string) *Category {
	category := Category{
		Name:   name,
		Active: true,
	}
	return &category
}

type Project struct {
	gorm.Model
	Name       string
	Active     bool
	CategoryId uint
	Category   *Category `gorm:"foreignKey:CategoryId"`
}

func (prj Project) String() string {
	return fmt.Sprintf("%s::%s", prj.Category.Name, prj.Name)
}

func NewProject(name string) *Project {
	project := Project{
		Name:   name,
		Active: true,
	}
	return &project
}

type Entry struct {
	gorm.Model
	Start     time.Time
	End       time.Time
	Duration  int64
	ProjectId uint
	Project   *Project `gorm:"foreignKey:ProjectId"`
	Note      string
}

func NewEntry(project *Project) *Entry {
	entry := Entry{
		Start:   time.Now(),
		Project: project,
	}
	return &entry
}

func (e *Entry) CalculateDuration() {
	if e.Start.IsZero() || e.End.IsZero() {
		e.Duration = 0
	}
	d := e.End.Sub(e.Start)
	seconds := math.Floor(d.Seconds())
	e.Duration = int64(seconds)
}

func (e *Entry) StopEntry() {
	e.End = time.Now()

	e.CalculateDuration()
}

func (e *Entry) IsZero() bool {
	return e.ID == 0 && e.End.IsZero()
}

func (e *Entry) IsActive() bool {
	return e.ID > 0 && !e.Start.IsZero() && e.End.IsZero()
}
