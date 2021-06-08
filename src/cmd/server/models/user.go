package models

import (
	"errors"
	"shoeguard-main-backend/cmd/server/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	PhoneNumber        string `gorm:"unique;not null" json:"phone_number"` // username field
	Password           string `gorm:"not null"        json:"password"`     // password field
	IsStudent          bool   `gorm:"not null"        json:"is_student"`
	PartnerPhoneNumber string `                       json:"partner_phone_number"`
	Nickname           string `gorm:"not null"        json:"nickname"`
}

func (user *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func (user *User) Create() error {
	db := utils.GetDB()
	if result := db.Where("phone_number = ?", user.PhoneNumber).Find(&user); result.RowsAffected != 0 {
		return errors.New("phone number duplicates")
	}
	user.HashPassword()
	return db.Create(&user).Error
}

func (user *User) SetUser(phoneNumber string) error {
	db := utils.GetDB()
	if result := db.Find(&user, "phone_number = ?", phoneNumber); result.Error != nil {
		return result.Error
	}
	return nil
}

func (user *User) IsPasswordCorrect(password string) bool {
	hashedPassword := user.Password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
