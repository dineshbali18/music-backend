package main

import (
	musicDelivery "musicApp/music/delivery/http"
	musicRepository "musicApp/music/usecase"
	musicUsecase "musicApp/music/repository/mysql"
)

var (
	e *echo.Echo
)

func init() {
	//Initialize config
	config.InitializeConfig()
	e = echo.New()
}

func main() {
	//Load Database config from config.yml
	err := config.GetDatabaseConfig()
	if err != nil {
		log.Println(err.Error())
	}

	// Establish data base connection
	db, err := gorm.Open(mysql.Open(config.DatabaseConfig.DatabaseURL), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		log.Println(err.Error())
	}

	// Specifying DB Reader and Writer
	err = db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(config.DatabaseConfig.DatabaseWriteURL)},
		Replicas: []gorm.Dialector{mysql.Open(config.DatabaseConfig.DatabaseReadURL)},
		Policy:   dbresolver.RandomPolicy{},
	}))

	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println("DATABASE CONNECTED SUCCESSFULLY")

	rdb := cacheServices.InitRedisCacheService()
	cacheService := cacheServices.NewRedisCacheService(rdb)

	res, err := cacheService.CheckRedisConnection()

	if err != nil {
		fmt.Println("Redis not connected properly", err)
		return
	} else {
		fmt.Println("Redis connected succesfully....", res)
	}

	musicDelivery.NewBBHandler(e, musicUsecase.NewUser(musicRepository.NewUser(db), cacheService))
	log.Fatal(e.Start(":" + "80"))

}
