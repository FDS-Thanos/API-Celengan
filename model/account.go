package model

type Account struct {
	AccountID string `gorm:"primaryKey" json:"account_id"`
	Username  string `gorm:"not null" json:"username"`
	Password  string `gorm:"not null" json:"password"`
	Name      string `json:"name"`
}

func (Account) TableName() string {
	return "account"
}
