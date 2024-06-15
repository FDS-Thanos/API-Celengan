package usecase

import (
	"api_gateway/model"
	"api_gateway/utils"
	"errors"
	"log"

	"gorm.io/gorm"
)

// type Login struct{}

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
	db, err := orm.DB()
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return false, errors.New("gagal menghubungi database")
	}
	defer db.Close()

	// Mencari akun berdasarkan username
	result := orm.Where("username = ?", Username).First(&bodyPayloadAuth)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("Username tidak ditemukan: %v", Username)
			return false, errors.New("username tidak ditemukan")
		}
		log.Printf("Error querying database: %v", result.Error)
		return false, errors.New("gagal melakukan query ke database")
	}

	// Verifikasi kata sandi
	if bodyPayloadAuth.Password != Password {
		log.Printf("Password tidak cocok untuk username: %v", Username)
		return false, errors.New("password tidak cocok")
	}

	// Sukses
	return true, nil
}
