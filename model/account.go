package model

type Account struct {
	ID       string `gorm:"primaryKey" json:"account_id"`
	Username string `gorm:"not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Name     string `json:"name"`
	Balance  float64
}

func (Account) TableName() string {
	return "account"
}
