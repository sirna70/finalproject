package models

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username    string        `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Your Username is required"`
	Email       string        `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your Email is required,email~Invalid Email format"`
	Password    string        `gorm:"not null" json:"password" form:"password" valid:"required~Your Password is required,minstringlength(6)~Your Password must be at least 6 characters"`
	Age         int           `gorm:"not null" json:"age" form:"age" valid:"required~Your Age is required"`
	Photo       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
	Comment     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	SocialMedia []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"social_media"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if u.Age < 8 {
		err = errors.New("your age must be at least 8")
		return
	}

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	// check email null or not
	if u.Email == "" {
		err = errors.New("your Email is required")
		return
	}

	// check email format
	if !govalidator.IsEmail(u.Email) {
		err = errors.New("invalid Email format")
		return
	}

	// check username null or not
	if u.Username == "" {
		err = errors.New("your Username is required")
		return
	}

	err = nil
	return
}
