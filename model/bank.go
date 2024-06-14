package model

type Bank struct {
	BankCode string `gorm:"primaryKey" json:"bank_code"`
	Name     string `gorm:"not null" json:"name"`
	Address  string `json:"address"`
}

func (Bank) TableName() string {
	return "bank"
}
