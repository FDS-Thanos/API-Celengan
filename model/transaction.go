package model

import (
	"time"
)

type Transaction struct {
	Id                   int `gorm:"primaryKey"`
	AccountId            string
	TransactionReference string
	TransactionAmount    int
	TransactionDate      *time.Time `gorm:"type:timestamp with time zone"`
}

func (Transaction) TableName() string {
	return "transaction"
}
