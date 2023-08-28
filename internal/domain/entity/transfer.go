package entity

import (
	"time"
)

type Transfer struct {
	ID                   uint      `gorm:"primarykey" json:"id"`
	AccountOriginID      uint      `gorm:"column:account_origin_id;NOT NULL" json:"account_origin_id"`
	AccountDestinationID uint      `gorm:"column:account_destination_id;NOT NULL" json:"account_destination_id"`
	Amount               float64   `gorm:"column:amount;NOT NULL" json:"amount"`
	CreatedAt            time.Time `gorm:"column:createdAt" json:"createdAt"`
	AccountOrigin        *Account  `gorm:"foreignKey:AccountOriginID" json:"-"`
	AccountDestination   *Account  `gorm:"foreignKey:AccountDestinationID" json:"-"`
}
