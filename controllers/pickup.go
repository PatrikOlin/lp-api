package controllers

import (
	"log"
	"net/http"

	"github.com/PatrikOlin/lp-api/db"
	"github.com/PatrikOlin/lp-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllPickups(c *gin.Context) {
	var pickups []models.Pickup

	res := db.GlobalDB.Preload("Haul").Find(&pickups)

	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not get pickups",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, pickups)

	return
}

func GetPickupByID(c *gin.Context) {
	var pickup models.Pickup

	id := c.Params.ByName("id")

	res := db.GlobalDB.Preload("Haul").Where("id = ?", id).First(&pickup)

	if res.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "pickup not found",
		})
		c.Abort()
		return
	}

	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not get pickup",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": pickup,
	})

	return
}

func GetPickupsByUserID(c *gin.Context) {
	var pickup []models.Pickup

	userID := c.Params.ByName("id")

	res := db.GlobalDB.Debug().Preload("Haul").Where("user_id = ?", userID).Find(&pickup)

	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not get users pickups",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": pickup,
	})
}

func CreatePickup(c *gin.Context) {
	var pickup models.Pickup
	var user models.User

	email, _ := c.Get("email")

	res := db.GlobalDB.Where("email = ?", email).First(&user)

	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		c.Abort()
		return
	}

	pickup.UserID = user.ID

	err := c.ShouldBindJSON(&pickup)
	if err != nil {
		log.Println(err)

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json",
		})
		c.Abort()
		return
	}

	err = pickup.CreatePickupRecord()
	if err != nil {
		log.Println(err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error creating pickup",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pickup})

	return
}

func UpdatePickupByID(c *gin.Context) {
	var pickup models.Pickup
	id := c.Params.ByName("id")

	res := db.GlobalDB.Where("id = ?", id).First(&pickup)

	if res.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "pickup not found",
		})
		c.Abort()
		return
	}

	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not get pickup",
		})
		c.Abort()
		return
	}

	if err := c.ShouldBindJSON(&pickup); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	db.GlobalDB.Save(&pickup)

	c.JSON(http.StatusOK, gin.H{"data": pickup})

	return
}

func DeletePickupByID(c *gin.Context) {
	var pickup models.Pickup
	id := c.Params.ByName("id")

	res := db.GlobalDB.Where("id = ?", id).Delete(&pickup)

	if res.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "pickup not found",
		})
		c.Abort()
		return
	}

	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not remove pickup",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "pickup with id " + string(id) + " removed"})
}
