package main

import (
	"github.com/api-go-gin/database"
	"github.com/api-go-gin/models"
	"github.com/api-go-gin/routes"
)

func main() {
	database.ConnectDatabase()
	models.Students = []models.Student{
		{Name: "Alberto Parente", CPF: "00000000000", RG: "9900000000000"},
		{Name: "Juliana Cavalcante", CPF: "11111111111", RG: "8800000000000"},
	}
	routes.HandleRequests()
}
