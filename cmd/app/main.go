package main

import (
	"os"

	"github.com/fms85/desafio-tecnico-go-stone/internal/delivery/api"
	"github.com/fms85/desafio-tecnico-go-stone/internal/domain/common"
	"github.com/fms85/desafio-tecnico-go-stone/internal/domain/migration"
	gormDriver "github.com/fms85/desafio-tecnico-go-stone/internal/driver/gorm"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	if os.Getenv("APP_ENV") == "" {
		if err := godotenv.Load(".env"); err != nil {
			panic(err)
		}
	}

	db := map[string]*gorm.DB{
		"wr": gormDriver.Setup(os.Getenv("DB_CONNECTION_WRITE")),
		"rd": gormDriver.Setup(os.Getenv("DB_CONNECTION_READ")),
	}

	migration.Run(db["wr"])

	app := common.App{
		DB: db,
		Env: common.Env{
			HTTP_ADDR:  os.Getenv("HTTP_ADDR"),
			JWT_SECRET: os.Getenv("JWT_SECRET"),
		},
	}

	api.Init(app)
}
