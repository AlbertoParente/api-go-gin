package database

import (
	"log"

	"github.com/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	stringToConnection := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(stringToConnection))
	if err != nil {
		log.Panic("Error connecting to database...!")
	}
	DB.AutoMigrate(&models.Student{
		Model: gorm.Model{},
		Name:  "",
		CPF:   "",
		RG:    "",
	})
}
