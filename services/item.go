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

// Materias
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

func (m *materiaService) GetAllMaterias() ([]database.Materias, error) {
	var materias []database.Materias
	err := m.db.Find(&materias).Error
	return materias, err
}

// Calificaciones
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

/*
	func (c *calificacionService) GetCalificacionesByStudentID(studentID uint) ([]database.Calificaciones, error) {
		var calificaciones []database.Calificaciones
		err := c.db.Where("student_id = ?", studentID).Find(&calificaciones)
		return calificaciones, err.Error
	}
*/
func (c *calificacionService) GetCalificacionesByStudentID(studentID uint) ([]database.CalificacionConDetalles, error) {
	var calificaciones []database.CalificacionConDetalles
	err := c.db.Table("grades").
		Select("grades.grade_id, grades.student_id, grades.subject_id, grades.grade, subjects.name as subject_name, Alumnos.name as student_name").
		Joins("left join subjects on grades.subject_id = subjects.subject_id").
		Joins("left join Alumnos on grades.student_id = Alumnos.student_id").
		Where("grades.student_id = ?", studentID).
		Find(&calificaciones).Error
	return calificaciones, err
}
