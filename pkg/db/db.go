package db

import (
	"kitco-parser/configs"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(conf *configs.DbConfig) *Db {
	db, err := gorm.Open(postgres.Open(conf.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	return &Db{db}
}