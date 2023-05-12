package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/softclub-go-0-0/crm-service/pkg/database"
	"github.com/softclub-go-0-0/crm-service/pkg/handlers"
	"github.com/softclub-go-0-0/crm-service/pkg/middlewares"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.DBInit(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	if err != nil {
		log.Fatal("db connection error:", err)
	}

	log.Println("successfully connected to DB")

	h := handlers.NewHandler(db)

	router := gin.Default()

	router.Use(middlewares.AuthMiddleware())

	router.GET("/teachers", h.GetAllTeachers)
	router.POST("/teachers", h.CreateTeacher)
	router.GET("/teachers/:teacherID", h.GetOneTeacher)
	router.PUT("/teachers/:teacherID", h.UpdateTeacher)
	router.DELETE("/teachers/:teacherID", h.DeleteTeacher)

	router.GET("/courses", h.GetAllCourses)
	router.POST("/courses", h.CreateCourse)
	router.GET("/courses/:courseID", h.GetOneCourse)
	router.PUT("/courses/:courseID", h.UpdateCourse)
	router.DELETE("/courses/:courseID", h.DeleteCourse)

	router.GET("/timetables", h.GetAllTimetables)
	router.POST("/timetables", h.CreateTimetable)
	router.GET("/timetables/:timetableID", h.GetOneTimetable)
	router.PUT("/timetables/:timetableID", h.UpdateTimetable)
	router.DELETE("/timetables/:timetableID", h.DeleteTimetable)

	groups := router.Group("/groups")
	{
		groups.GET("/", h.GetAllGroups)
		groups.POST("/", h.CreateGroup)
		groups.GET("/:groupID", h.GetOneGroup)
		groups.PUT("/:groupID", h.UpdateGroup)
		groups.DELETE("/:groupID", h.DeleteGroup)

		students := groups.Group("/:groupID/students")
		{
			students.GET("/", h.GetAllStudents)
			students.POST("/", h.CreateStudent)
			students.GET("/:studentID", h.GetOneStudent)
			students.PUT("/:studentID", h.UpdateStudent)
			students.DELETE("/:studentID", h.DeleteStudent)
		}
	}

	router.Run(":" + os.Getenv("APP_PORT"))
}
