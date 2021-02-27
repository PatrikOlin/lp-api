package models

import (
	"github.com/PatrikOlin/lp-api/db"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string        `json:"email" gorm:"unique"`
	Password     string        `json:"password"`
	Profile      Profile       `json:"profile" gorm:"foreignKey:UserID"`
	Pickups      []Pickup      `json:"pickups" gorm:"foreignKey:UserID"`
	Propositions []Proposition `json:"propositions" gorm:"foreignKey:UserID"`
}

type Profile struct {
	gorm.Model
	UserID   uint    `json:"userID"`
	Name     string  `json:"name"`
	Address  string  `json:"address"`
	Recycler bool    `json:"recycler"`
	Rating   float64 `json:"rating"`
	Phone    string  `json:"phone"`
}

func (user *User) CreateUserRecord() error {
	result := db.GlobalDB.Session(&gorm.Session{FullSaveAssociations: true}).Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	user.Password = string(bytes)

	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}
