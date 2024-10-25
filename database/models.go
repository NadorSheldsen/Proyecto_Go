package database

import "gorm.io/gorm"

// Alumnos
type Alumnos struct {
	gorm.Model
	Student_id int64 `gorm:"column:student_id;primaryKey"`
	Name       string
	Group      string
	Email      string
}

func (Alumnos) TableName() string {
	return "Alumnos"
}

// Materias
type Materias struct {
	SubjectID int64  `gorm:"column:subject_id;primaryKey;autoIncrement"`
	Name      string `gorm:"type:varchar(255);not null"`
}

func (Materias) TableName() string {
	return "subjects"
}

// Calificaciones
type Calificaciones struct {
	GradeID   int64   `gorm:"column:grade_id;primaryKey;autoIncrement"`
	StudentID int64   `gorm:"column:student_id"`
	SubjectID int64   `gorm:"column:subject_id"`
	Grade     float32 `gorm:"type:float"`
}

func (Calificaciones) TableName() string {
	return "grades"
}
