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
<<<<<<< HEAD

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
=======
func (Grades) TableName() string {
	return "Calificaciones"
}

func (Subjects) TableName() string {
	return "Materias"
>>>>>>> 1935322cc44571609ccfb0750421356905541b28
}
