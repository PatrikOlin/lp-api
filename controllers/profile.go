package controllers

import (
	"github.com/PatrikOlin/lp-api/db"
	"github.com/PatrikOlin/lp-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Profile(c *gin.Context) {
	var user models.User

	email, _ := c.Get("email")

	result := db.GlobalDB.Where("email = ?", email.(string)).First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(404, gin.H{
			"msg": "user not found",
		})
		c.Abort()
		return
	}

	if result.Error != nil {
		c.JSON(500, gin.H{
			"msg": "could not get user profile",
		})
		c.Abort()
		return
	}

	user.Password = ""

	c.JSON(200, user)

	return
}
