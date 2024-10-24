package services

import (
	"escuela/database"

	"gorm.io/gorm"
)

type AlumnoService interface {
	CreateAlumno(alumno database.Alumnos) error
	UpdateAlumno(alumno database.Alumnos) error
	DeleteAlumno(id uint) error
	GetAlumnoByID(id uint) (database.Alumnos, error)
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

func (a *alumnoService) DeleteAlumno(id uint) error {
	err := a.db.Delete(&database.Alumnos{}, id)
	return err.Error
}

func (a *alumnoService) GetAlumnoByID(id uint) (database.Alumnos, error) {
	var alumno database.Alumnos
	err := a.db.First(&alumno, id)
	return alumno, err.Error
}

func (a *alumnoService) GetAllAlumnos() ([]database.Alumnos, error) {
	var alumnos []database.Alumnos
	err := a.db.Find(&alumnos)
	return alumnos, err.Error
}
