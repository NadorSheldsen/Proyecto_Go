package main

import (
	"escuela/database"
	"escuela/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.NewDatabaseDriver()
	if err != nil {
		fmt.Println("Error al conectar a la base de datos: ", err)
		return
	}
	db.AutoMigrate(&database.Alumnos{})
	AlumnoService := services.NewAlumnoService(db)

	// Inicialización del router
	router := gin.Default()

	// Sirve archivos estáticos como CSS o JS desde la carpeta 'static'
	router.Static("/static", "./static")

	// Carga las plantillas HTML
	router.LoadHTMLGlob("vistas/*")

	// Rutas
	router.GET("/students/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "nuevo_estudiante.html", nil)
	})

	router.POST("/bd/students", func(c *gin.Context) {
		nombre := c.PostForm("nombre")
		grupo := c.PostForm("grupo")
		email := c.PostForm("email")

		var count int64
		if err := db.Model(&database.Alumnos{}).Count(&count).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al contar los alumnos"})
			return
		}

		nuevoAlumno := database.Alumnos{
			Student_id: count + 1,
			Name:       nombre,
			Group:      grupo,
			Email:      email,
		}

		err := AlumnoService.CreateAlumno(nuevoAlumno)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/students")
	})

	router.GET("/students", func(c *gin.Context) {
		alumnos, err := AlumnoService.GetAllAlumnos()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.HTML(http.StatusOK, "estudiantes.html", gin.H{"alumnos": alumnos})
	})

	router.GET("/students/:student_id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("student_id"))
		alumno, err := AlumnoService.GetAlumnoByID(uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.HTML(http.StatusOK, "ver_estudiante.html", gin.H{"alumno": alumno})
	})

	router.POST("/bd/students/delete/:student_id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("student_id"))

		err := AlumnoService.DeleteAlumno(uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Redirect(http.StatusMovedPermanently, "/students")
	})

	router.GET("/students/act/:student_id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("student_id"))
		alumno, err := AlumnoService.GetAlumnoByID(uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.HTML(http.StatusOK, "actualizar_estudiante.html", gin.H{"alumno": alumno, "id": id})
	})

	router.POST("/bd/students/update/:student_id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("student_id"))
		alumno, err := AlumnoService.GetAlumnoByID(uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		alumno.Name = c.PostForm("nombre")
		alumno.Group = c.PostForm("grupo")
		alumno.Email = c.PostForm("email")

		err = AlumnoService.UpdateAlumno(alumno)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/students")
	})

	// Inicia el servidor en el puerto 8000
	router.Run(":8000")
}
