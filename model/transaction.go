package model

import (
	"time"
)

type Transaction struct {
	Id               int `gorm:"primaryKey"`
	Account_id       string
	Bank_id          string
	Amount           int
	Transaction_date *time.Time `gorm:"type:timestamp with time zone"`
}

func (Transaction) TableName() string {
	return "transaction"
}
