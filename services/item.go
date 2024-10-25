package services

import (
	"escuela/database"

	"gorm.io/gorm"
)

// Servicio de Alumnos
type AlumnoService interface {
	CreateAlumno(alumno database.Alumnos) error
	UpdateAlumno(alumno database.Alumnos) error
	DeleteAlumno(id int64) error
	GetAlumnoByID(id int64) (database.Alumnos, error)
	GetAllAlumnos() ([]database.Alumnos, error)
}

type alumnoService struct {
	db *gorm.DB
}

func NewAlumnoService(db *gorm.DB) AlumnoService {
	return &alumnoService{db}
}

func (a *alumnoService) CreateAlumno(alumno database.Alumnos) error {
	err := a.db.Create(&alumno)
	return err.Error
}

func (a *alumnoService) UpdateAlumno(alumno database.Alumnos) error {
	err := a.db.Save(&alumno)
	return err.Error
}

func (a *alumnoService) DeleteAlumno(id int64) error {
	err := a.db.Delete(&database.Alumnos{}, id)
	return err.Error
}

func (a *alumnoService) GetAlumnoByID(id int64) (database.Alumnos, error) {
	var alumno database.Alumnos
	err := a.db.First(&alumno, id)
	return alumno, err.Error
}

func (a *alumnoService) GetAllAlumnos() ([]database.Alumnos, error) {
	var alumnos []database.Alumnos
	err := a.db.Find(&alumnos)
	return alumnos, err.Error
}

// Servicio de Materias
type SubjectService interface {
	CreateSubject(subject database.Subjects) error
	UpdateSubject(subject database.Subjects) error
	DeleteSubject(id int64) error
	GetSubjectByID(id int64) (database.Subjects, error)
	GetAllSubjects() ([]database.Subjects, error)
}

type subjectService struct {
	db *gorm.DB
}

func NewSubjectService(db *gorm.DB) SubjectService {
	return &subjectService{db}
}

func (s *subjectService) CreateSubject(subject database.Subjects) error {
	err := s.db.Create(&subject)
	return err.Error
}

func (s *subjectService) UpdateSubject(subject database.Subjects) error {
	err := s.db.Save(&subject)
	return err.Error
}

func (s *subjectService) DeleteSubject(id int64) error {
	err := s.db.Delete(&database.Subjects{}, id)
	return err.Error
}

func (s *subjectService) GetSubjectByID(id int64) (database.Subjects, error) {
	var subject database.Subjects
	err := s.db.First(&subject, id)
	return subject, err.Error
}

func (s *subjectService) GetAllSubjects() ([]database.Subjects, error) {
	var subjects []database.Subjects
	err := s.db.Find(&subjects)
	return subjects, err.Error
}

// Servicio de Calificaciones
type GradeService interface {
	CreateGrade(grade database.Grades) error
	UpdateGrade(grade database.Grades) error
	DeleteGrade(id int64) error
	GetGradeByID(id int64) (database.Grades, error)
	GetGradesByStudentID(studentID int64) ([]database.Grades, error)
}

type gradeService struct {
	db *gorm.DB
}

func NewGradeService(db *gorm.DB) GradeService {
	return &gradeService{db}
}

func (g *gradeService) CreateGrade(grade database.Grades) error {
	err := g.db.Create(&grade)
	return err.Error
}

func (g *gradeService) UpdateGrade(grade database.Grades) error {
	err := g.db.Save(&grade)
	return err.Error
}

func (g *gradeService) DeleteGrade(id int64) error {
	err := g.db.Delete(&database.Grades{}, id)
	return err.Error
}

func (g *gradeService) GetGradeByID(id int64) (database.Grades, error) {
	var grade database.Grades
	err := g.db.First(&grade, id)
	return grade, err.Error
}

func (g *gradeService) GetGradesByStudentID(studentID int64) ([]database.Grades, error) {
	var grades []database.Grades
	err := g.db.Where("student_id = ?", studentID).Find(&grades)
	return grades, err.Error
}
