package models

import (
	"errors"

	"github.com/dhairyajoshi/go-rest-api/helpers/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
}

func migrateUsers() {

	conn := database.GetDbConnection()

	conn.AutoMigrate(&User{})
}

func (user *User) Save() {
	conn := database.GetDbConnection()

	password := []byte(user.Password)

	hashed, err := bcrypt.GenerateFromPassword(password, 14)

	if err != nil {
		return
	}

	user.Password = string(hashed)

	conn.Create(&user)
}

func (User) Authenticate(email, password string) (bool, error) {
	conn := database.GetDbConnection()

	var user User

	result := conn.First(&user, "email = ?", email)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, gorm.ErrRecordNotFound
	}

	hashed := []byte(user.Password)
	pass := []byte(password)

	err := bcrypt.CompareHashAndPassword(hashed, pass)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (User) Exists(email string) bool {
	conn := database.GetDbConnection()

	var user User

	result := conn.First(&user, "email = ?", email)

	return !errors.Is(result.Error, gorm.ErrRecordNotFound)

}
