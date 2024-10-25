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
	db.AutoMigrate(&database.Materias{})
	db.AutoMigrate(&database.Calificaciones{})
	AlumnoService := services.NewAlumnoService(db)
	MateriaService := services.NewMateriaService(db)
	CalificacionService := services.NewCalificacionService(db)

	router := gin.Default()

	router.LoadHTMLGlob("vistas/*")

	router.GET("/students/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "nuevo_estudiante.html", nil)
	})

	router.POST("/bd/students", func(c *gin.Context) {
		nombre := c.PostForm("nombre")
		grupo := c.PostForm("grupo")
		email := c.PostForm("email")

		/*var count int64
		if err := db.Model(&database.Alumnos{}).Count(&count).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al contar los alumnos"})
			return
		}*/

		nuevoAlumno := database.Alumnos{
			//Student_id: count + 1,
			Name:  nombre,
			Group: grupo,
			Email: email,
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
		alumno.Group = c.PostForm("grupo") // Asegúrate de que el nombre del campo coincida con el formulario HTML
		alumno.Email = c.PostForm("email")

		err = AlumnoService.UpdateAlumno(alumno)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/students")
	})

	//Materias

	router.GET("/subjects/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "crear_materia.html", nil)
	})
	router.POST("/escuela/subjects", func(c *gin.Context) {
		nombre := c.PostForm("nombre")
		nuevaMateria := database.Materias{Name: nombre}

		err := MateriaService.CreateMateria(nuevaMateria)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Materia creada exitosamente"})
	})

	router.PUT("/escuela/subjects/:subject_id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("subject_id"))
		materia, err := MateriaService.GetMateriaByID(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Materia no encontrada"})
			return
		}

		materia.Name = c.PostForm("nombre")
		err = MateriaService.UpdateMateria(materia)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Materia actualizada exitosamente"})
	})

	router.GET("/subjects/:subject_id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("subject_id"))
		materia, err := MateriaService.GetMateriaByID(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Materia no encontrada"})
			return
		}

		c.HTML(http.StatusOK, "ver_materia.html", gin.H{"materia": materia})
	})

	router.DELETE("/escuela/subjects/:subject_id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("subject_id"))
		err := MateriaService.DeleteMateria(uint(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Materia eliminada exitosamente"})
	})

	//Calificaciones
	router.GET("/grades/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "crear_calificacion.html", nil)
	})
	router.POST("/escuela/grades", func(c *gin.Context) {
		studentID, _ := strconv.Atoi(c.PostForm("student_id"))
		subjectID, _ := strconv.Atoi(c.PostForm("subject_id"))
		grade, _ := strconv.ParseFloat(c.PostForm("grade"), 32)

		nuevaCalificacion := database.Calificaciones{
			StudentID: int64(studentID),
			SubjectID: int64(subjectID),
			Grade:     float32(grade),
		}

		err := CalificacionService.CreateCalificacion(nuevaCalificacion)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Calificación creada exitosamente"})
	})

	router.PUT("/grades/act/:grade_id", func(c *gin.Context) {
		gradeID, _ := strconv.Atoi(c.Param("grade_id"))
		calificacion, err := CalificacionService.GetCalificacionByID(uint(gradeID))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Calificación no encontrada"})
			return
		}

		gradeValue, _ := strconv.ParseFloat(c.PostForm("grade"), 32)
		calificacion.Grade = float32(gradeValue)
		err = CalificacionService.UpdateCalificacion(calificacion)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Calificación actualizada exitosamente"})
	})

	router.DELETE("/grades/delete/:grade_id", func(c *gin.Context) {
		gradeID, _ := strconv.Atoi(c.Param("grade_id"))
		err := CalificacionService.DeleteCalificacion(uint(gradeID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Calificación eliminada exitosamente"})
	})

	router.GET("/grades/:grade_id", func(c *gin.Context) {
		gradeID, _ := strconv.Atoi(c.Param("grade_id"))
		studentID, _ := strconv.Atoi(c.Param("student_id"))
		calificacion, err := CalificacionService.GetCalificacionByGradeIDAndStudentID(uint(gradeID), uint(studentID))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Calificación no encontrada"})
			return
		}

		c.HTML(http.StatusOK, "ver_calificacion.html", gin.H{"calificacion": calificacion})
	})

	router.GET("/grades/student/:student_id", func(c *gin.Context) {
		studentID, _ := strconv.Atoi(c.Param("student_id"))
		calificaciones, err := CalificacionService.GetCalificacionesByStudentID(uint(studentID))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.HTML(http.StatusOK, "ista_calificaciones_estudiante.html", gin.H{
			"student_id": studentID,
			"grades":     calificaciones,
		})
		fmt.Println("AQUI")
		fmt.Println(calificaciones)
	})

	err = router.Run(":8000")
	if err != nil {
		fmt.Println("Error al iniciar el servidor: ", err)
		return
	}
}
