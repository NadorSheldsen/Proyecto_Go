package database

// Alumnos
type Alumnos struct {
	Student_id int64  `gorm:"column:student_id;primaryKey;autoIncrement"`
	Name       string `gorm:"type:varchar(255);not null"`
	Group      string `gorm:"type:varchar(255);not null"`
	Email      string `gorm:"type:varchar(255);not null"`
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

type CalificacionConDetalles struct {
	GradeID     int64   `gorm:"column:grade_id"`
	StudentID   int64   `gorm:"column:student_id"`
	SubjectID   int64   `gorm:"column:subject_id"`
	Grade       float32 `gorm:"type:float"`
	SubjectName string  `gorm:"column:subject_name"`
	StudentName string  `gorm:"column:student_name"`
}
