package withgorm

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"

	"vikishptra/domain_pensfess/model/entity"
	"vikishptra/shared/gogen"
	"vikishptra/shared/infrastructure/config"
	"vikishptra/shared/infrastructure/logger"
)

type gateway struct {
	log     logger.Logger
	appData gogen.ApplicationData
	config  *config.Config
	Db      *gorm.DB
}

func NewGateway(log logger.Logger, appData gogen.ApplicationData, cfg *config.Config) *gateway {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	Db, err := gorm.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	err = Db.AutoMigrate(entity.User{}).Error

	if err != nil {
		panic(err)
	}

	return &gateway{
		log:     log,
		appData: appData,
		config:  cfg,
		Db:      Db,
	}
}
