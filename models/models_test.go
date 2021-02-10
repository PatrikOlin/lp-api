package models

import (
	"github.com/PatrikOlin/lp-api/db"
	"gorm.io/gorm"

	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	user := User{
		Password: "secret",
	}

	err := user.HashPassword(user.Password)
	assert.NoError(t, err)

	os.Setenv("passwordHash", user.Password)
}

func TestCreateUserRecord(t *testing.T) {
	var userResult User

	err := db.InitDatabase()
	if err != nil {
		t.Error(err)
	}

	err = db.GlobalDB.AutoMigrate(&User{})
	assert.NoError(t, err)

	user := User{
		Model:    gorm.Model{},
		Name:     "Test user",
		Email:    "test@email.com",
		Password: os.Getenv("passwordHash"),
	}

	err = user.CreateUserRecord()
	assert.NoError(t, err)

	db.GlobalDB.Where("email = ?", user.Email).Find(&userResult)

	db.GlobalDB.Unscoped().Delete(&user)

	assert.Equal(t, "Test user", userResult.Name)
	assert.Equal(t, "test@email.com", userResult.Email)
}

func TestCheckPassword(t *testing.T) {
	hash := os.Getenv("passwordHash")

	user := User{
		Password: hash,
	}

	err := user.CheckPassword("secret")
	assert.NoError(t, err)
}
