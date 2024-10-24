package database

import (
	"escuela/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabaseDriver() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.BaseDeDatos), &gorm.Config{})
	if err != nil {
		fmt.Println("Error al conectar a la base de datos: ", err)
		return nil, err
	}
	return db, nil
}
