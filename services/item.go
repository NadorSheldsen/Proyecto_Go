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

<<<<<<< HEAD
// Materias
type MateriaService interface {
	CreateMateria(materia database.Materias) error
	UpdateMateria(materia database.Materias) error
	DeleteMateria(id uint) error
	GetMateriaByID(id uint) (database.Materias, error)
}

type materiaService struct {
	db *gorm.DB
}

func NewMateriaService(db *gorm.DB) MateriaService {
	return &materiaService{db}
}

func (m *materiaService) CreateMateria(materia database.Materias) error {
	err := m.db.Create(&materia)
	return err.Error
}

func (m *materiaService) UpdateMateria(materia database.Materias) error {
	err := m.db.Save(&materia)
	return err.Error
}

func (m *materiaService) DeleteMateria(id uint) error {
	err := m.db.Delete(&database.Materias{}, id)
	return err.Error
}

func (m *materiaService) GetMateriaByID(id uint) (database.Materias, error) {
	var materia database.Materias
	err := m.db.First(&materia, id)
	return materia, err.Error
}

// Calificaciones
type CalificacionService interface {
	CreateCalificacion(calificacion database.Calificaciones) error
	UpdateCalificacion(calificacion database.Calificaciones) error
	DeleteCalificacion(id uint) error
	GetCalificacionByID(id uint) (database.Calificaciones, error)
	GetCalificacionByGradeIDAndStudentID(gradeID, studentID uint) (database.Calificaciones, error)
	GetCalificacionesByStudentID(studentID uint) ([]database.Calificaciones, error)
}

type calificacionService struct {
	db *gorm.DB
}

func NewCalificacionService(db *gorm.DB) CalificacionService {
	return &calificacionService{db}
}

func (c *calificacionService) CreateCalificacion(calificacion database.Calificaciones) error {
	err := c.db.Create(&calificacion)
	return err.Error
}

func (c *calificacionService) UpdateCalificacion(calificacion database.Calificaciones) error {
	err := c.db.Save(&calificacion)
	return err.Error
}

func (c *calificacionService) DeleteCalificacion(id uint) error {
	err := c.db.Delete(&database.Calificaciones{}, id)
	return err.Error
}

func (c *calificacionService) GetCalificacionByID(id uint) (database.Calificaciones, error) {
	var calificacion database.Calificaciones
	err := c.db.First(&calificacion, id)
	return calificacion, err.Error
}

func (c *calificacionService) GetCalificacionByGradeIDAndStudentID(gradeID, studentID uint) (database.Calificaciones, error) {
	var calificacion database.Calificaciones
	err := c.db.Where("student_id = ?", studentID).First(&calificacion)
	return calificacion, err.Error
}

func (c *calificacionService) GetCalificacionesByStudentID(studentID uint) ([]database.Calificaciones, error) {
	var calificaciones []database.Calificaciones
	err := c.db.Where("student_id = ?", studentID).Find(&calificaciones)
	return calificaciones, err.Error

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
