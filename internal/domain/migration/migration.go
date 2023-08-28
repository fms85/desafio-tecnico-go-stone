package migration

import (
	"log"

	"github.com/fms85/desafio-tecnico-go-stone/internal/domain/entity"
	"gorm.io/gorm"
)

func Run(wr *gorm.DB) {
	if err := wr.AutoMigrate(
		&entity.Account{},
		&entity.Transfer{},
	); err != nil {
		log.Fatal(err)
	}
}
