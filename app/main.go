package main

import (
	"fmt"
	"log"
	"musicApp/config"

	musicDelivery "musicApp/music/delivery/http"
	musicRepository "musicApp/music/repository/mysql"
	musicUsecase "musicApp/music/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

var (
	e *echo.Echo
)

func init() {
	config.InitializeConfig()
	e = echo.New()
	e.Use(middleware.CORS())
}

func main() {

	err := config.GetDatabaseConfig()
	if err != nil {
		log.Println(err.Error())
	}

	db, err := gorm.Open(mysql.Open(config.DatabaseConfig.DatabaseURL), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		log.Println(err.Error())
	}

	err = db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(config.DatabaseConfig.DatabaseWriteURL)},
		Replicas: []gorm.Dialector{mysql.Open(config.DatabaseConfig.DatabaseReadURL)},
		Policy:   dbresolver.RandomPolicy{},
	}))

	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println("DATABASE CONNECTED SUCCESSFULLY")
	musicDelivery.NewMusicHandler(e, musicUsecase.NewMusicUsecase(musicRepository.NewMusicRepository(db)))
	log.Fatal(e.Start(":" + "80"))

}
