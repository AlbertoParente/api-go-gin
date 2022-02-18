package routes

import (
	"github.com/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.ShowAllStudents)
	r.GET("/:name", controllers.Greeting)
	r.POST("/students", controllers.CreateNewStudent)
	r.GET("/students/:id", controllers.GetById)
	r.DELETE("/student/:id", controllers.DeleteStudent)
	r.PATCH("/student/:id", controllers.EditStudent)
	r.GET("/student/cpf/:cpf", controllers.GetByCPF)
	r.Run()
}
