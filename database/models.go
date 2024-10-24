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

func (Alumnos) TableName() string {
	return "Alumnos"
}
