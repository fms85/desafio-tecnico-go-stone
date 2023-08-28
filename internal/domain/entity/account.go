package entity

import (
	"time"
)

const ENTITY_BALANCE_DEFAULT = 100

type Account struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Name      string    `gorm:"column:name;NOT NULL" json:"name"`
	CPF       string    `gorm:"column:cpf;NOT NULL;unique" json:"cpf"`
	Secret    string    `gorm:"column:secret;NOT NULL" json:"-"`
	Balance   float64   `gorm:"column:balance;NOT NULL" json:"balance"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
}
