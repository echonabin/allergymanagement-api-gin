package main

import (
	"github.com/echonabin/allergymanagement-api/configs"
	"github.com/echonabin/allergymanagement-api/db"
	"github.com/echonabin/allergymanagement-api/graph/model"
	"gorm.io/gorm"
)

func init() {
	configs.LoadEnvVariables()
	db.ConnectDatabase()
}

func main() {
	type Allergy struct {
		model.Allergy
		gorm.Model
	}

	type User struct {
		model.User
		gorm.Model
	}

	user := User{}
	allergy := Allergy{}

	db.Model.AutoMigrate(&allergy)
	db.Model.AutoMigrate(&user)
}
