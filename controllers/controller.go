package controllers

import "github.com/gin-gonic/gin"

func ShowAllStudents(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"name": "Alberto Parente",
	})
}

func Greeting(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API:": "It's all OK " + name + "...?",
	})
}
