package routes

import (
	"github.com/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.ShowAllStudents)
	r.GET("/:name", controllers.Greeting)
	r.Run()
}
