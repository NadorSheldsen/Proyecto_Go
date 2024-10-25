package database

import (
	"gorm.io/gorm"
)

type Alumnos struct {
	gorm.Model
	Student_id int64
	Name       string
	Group      string
	Email      string
}
type Grades struct {
	gorm.Model
	GradeID   int64   `gorm:"primaryKey"`
	StudentID int64   `gorm:"not null"`
	SubjectID int64   `gorm:"not null"`
	Grade     float64 `gorm:"not null"`
}
type Subjects struct {
	gorm.Model
	SubjectID int64  `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
}

func (Alumnos) TableName() string {
	return "Alumnos"
}
func (Grades) TableName() string {
	return "Calificaciones"
}

func (Subjects) TableName() string {
	return "Materias"
}
