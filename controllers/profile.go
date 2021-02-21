package controllers

import (
	"net/http"

	"github.com/PatrikOlin/lp-api/db"
	"github.com/PatrikOlin/lp-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProfileByID(c *gin.Context) {
	var user models.Profile
	id := c.Params.ByName("id")

	result := db.GlobalDB.Where("user_id = ?", id).First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user profile not found",
		})
		c.Abort()
		return
	}

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not get user profile",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, user)

	return
}

func GetAllProfiles(c *gin.Context) {
	var users []models.User

	res := db.GlobalDB.Preload("Profile").Find(&users)

	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": res.Error,
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, users)

	return
}

func UpdateProfile(c *gin.Context) {
	var user models.User

	email, _ := c.Get("email")

	result := db.GlobalDB.Where("email = ?", email).First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		c.Abort()
		return
	}

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not get user profile",
		})
		c.Abort()
		return
	}

	var profile models.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	db.GlobalDB.Model(&user).Association("Profile").Append(&profile)

	c.JSON(http.StatusOK, gin.H{"data": user.Profile})

	return
}

func ToggleRecycler(c *gin.Context) {
	var user models.User

	email, _ := c.Get("email")

	result := db.GlobalDB.Where("email = ?", email).First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		c.Abort()
		return
	}

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not get user profile",
		})
		c.Abort()
		return
	}

	var recycler models.Profile
	if err := c.ShouldBindJSON(&recycler); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	db.GlobalDB.Model(&user).Association("Profile").Append(&recycler)

	c.JSON(http.StatusOK, gin.H{"data": user.Profile})

	return
}
