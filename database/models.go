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
type Calificacion struct {
	gorm.Model
	GradeID   int64   `gorm:"primaryKey"`
	StudentID int64   `gorm:"not null"`
	SubjectID int64   `gorm:"not null"`
	Grade     float64 `gorm:"not null"`
}
type Materia struct {
	gorm.Model
	SubjectID int64  `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
}

func (Alumnos) TableName() string {
	return "Alumnos"
}
func (Calificacion) TableName() string {
	return "Calificaciones"
}

func (Materia) TableName() string {
	return "Materias"
}
