package db

import (
	"github.com/echonabin/allergymanagement-api/configs"
	logUtils "github.com/echonabin/allergymanagement-api/utils/logger"
	stringUtil "github.com/echonabin/allergymanagement-api/utils/stringify"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Model *gorm.DB

func ConnectDatabase() {
	var err error
	DbEnvVariables := configs.DbEnvConfig

	dsn := stringUtil.CreateFormattedString("host=%s user=%s password=%s dbname=%s port=%s", DbEnvVariables.DbHost, DbEnvVariables.DbUser, DbEnvVariables.DbPassword, DbEnvVariables.DbName, DbEnvVariables.DbPort)
	Model, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logUtils.Error(err, "Error connecting to database: ")
	}
}
