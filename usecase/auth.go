package usecase

import (
	"api_gateway/model"
	"api_gateway/utils"
	"log"

	"gorm.io/gorm"
)

type Login struct{}

type LoginInterface interface {
	Auth(username, password string) bool
}

func TaskLogin() LoginInterface {
	return &Login{}
}

func (pi *Login) Auth(username, password string) bool {
	if username == "ivialva" && password == "123456" {
		return true
	}
	return false
}

type Logindb struct{}

type LogindbInterface interface {
	Autentikasi(Username, Password string) bool
}

func NewLogin() LogindbInterface {
	return &Login{}
}

func (masuk *Login) Autentikasi(Username string, Password string) bool {
	bodyPayloadAuth := model.Account{}

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()
	// Mencari akun berdasarkan username
	result := orm.Where("username = ?", Username).First(&bodyPayloadAuth)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// Tidak ditemukan
			log.Printf("Username not found: %v", Username)
			return false
		}
		log.Printf("Error querying database: %v", result.Error)
		return false
	}

	// Verifikasi kata sandi
	if bodyPayloadAuth.Password != Password {
		// Kata sandi tidak cocok
		log.Printf("Password does not match for username: %v", Username)
		return false
	}

	// Sukses
	return true
}
