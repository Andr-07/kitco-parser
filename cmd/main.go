package main

import (
	"kitco-parser/configs"
	"kitco-parser/internal/repository"
	"kitco-parser/internal/service"
	"kitco-parser/pkg/db"
	"log"
)

func main() {
	conf := configs.LoadConfig()
	database := db.NewDb(&conf.Db)

	// Repositories
	newsMetaRepository := repository.NewNewsMetaRepository(database)
	// Services
	service := service.NewNewsMetaService(newsMetaRepository)

	if err := service.IngestKitcoNews(); err != nil {
		log.Fatal(err)
	}
}
