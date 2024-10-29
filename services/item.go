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
	return a.db.Create(&alumno).Error
}

func (a *alumnoService) UpdateAlumno(alumno database.Alumnos) error {
	return a.db.Save(&alumno).Error
}

func (a *alumnoService) DeleteAlumno(id int64) error {
	return a.db.Delete(&database.Alumnos{}, id).Error
}

func (a *alumnoService) GetAlumnoByID(id int64) (database.Alumnos, error) {
	var alumno database.Alumnos
	err := a.db.First(&alumno, id).Error
	return alumno, err
}

func (a *alumnoService) GetAllAlumnos() ([]database.Alumnos, error) {
	var alumnos []database.Alumnos
	err := a.db.Find(&alumnos).Error
	return alumnos, err
}

// Servicio de Materias
type MateriaService interface {
	CreateMateria(materia database.Materias) error
	UpdateMateria(materia database.Materias) error
	DeleteMateria(id uint) error
	GetMateriaByID(id uint) (database.Materias, error)
	GetAllMaterias() ([]database.Materias, error)
}

type materiaService struct {
	db *gorm.DB
}

func NewMateriaService(db *gorm.DB) MateriaService {
	return &materiaService{db}
}

func (m *materiaService) CreateMateria(materia database.Materias) error {
	return m.db.Create(&materia).Error
}

func (m *materiaService) UpdateMateria(materia database.Materias) error {
	return m.db.Save(&materia).Error
}

func (m *materiaService) DeleteMateria(id uint) error {
	return m.db.Delete(&database.Materias{}, id).Error
}

func (m *materiaService) GetMateriaByID(id uint) (database.Materias, error) {
	var materia database.Materias
	err := m.db.First(&materia, id).Error
	return materia, err
}

// Servicio de Calificaciones
type CalificacionService interface {
	CreateCalificacion(calificacion database.Calificaciones) error
	UpdateCalificacion(calificacion database.Calificaciones) error
	DeleteCalificacion(id uint) error
	GetCalificacionByID(id uint) (database.Calificaciones, error)
	GetCalificacionByGradeIDAndStudentID(gradeID, studentID uint) (database.Calificaciones, error)
	GetCalificacionesByStudentID(studentID uint) ([]database.CalificacionConDetalles, error)
}

type calificacionService struct {
	db *gorm.DB
}

func NewCalificacionService(db *gorm.DB) CalificacionService {
	return &calificacionService{db}
}

func (c *calificacionService) CreateCalificacion(calificacion database.Calificaciones) error {
	return c.db.Create(&calificacion).Error
}

func (c *calificacionService) UpdateCalificacion(calificacion database.Calificaciones) error {
	return c.db.Save(&calificacion).Error
}

func (c *calificacionService) DeleteCalificacion(id uint) error {
	return c.db.Delete(&database.Calificaciones{}, id).Error
}

func (c *calificacionService) GetCalificacionByID(id uint) (database.Calificaciones, error) {
	var calificacion database.Calificaciones
	err := c.db.First(&calificacion, id).Error
	return calificacion, err
}

func (c *calificacionService) GetCalificacionByGradeIDAndStudentID(gradeID, studentID uint) (database.Calificaciones, error) {
	var calificacion database.Calificaciones
	err := c.db.Where("student_id = ?", studentID).First(&calificacion, gradeID).Error
	return calificacion, err
}

func (c *calificacionService) GetCalificacionesByStudentID(studentID uint) ([]database.Calificaciones, error) {
	var calificaciones []database.Calificaciones
	err := c.db.Where("student_id = ?", studentID).Find(&calificaciones).Error
	return calificaciones, err
}
