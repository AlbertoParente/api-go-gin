package controllers

import (
	"github.com/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

func ShowAllStudents(c *gin.Context) {
	c.JSON(200, models.Students)
}

func Greeting(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API:": "It's all OK " + name + "...?",
	})
}
