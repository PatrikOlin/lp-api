package controllers

import (
	"net/http"

	"github.com/PatrikOlin/lp-api/db"
	"github.com/PatrikOlin/lp-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllHauls(c *gin.Context) {
	var hauls []models.Haul

	result := db.GlobalDB.Find(&hauls)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, hauls)

	return
}

func GetHaulByID(c *gin.Context) {
	var haul models.Haul
	id := c.Params.ByName("id")

	result := db.GlobalDB.Where("id = ?", id).First(&haul)

	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "haul not found",
		})
		c.Abort()
		return
	}

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could get get haul",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, haul)

	return
}

func UpdateHaulByID(c *gin.Context) {
	var haul models.Haul
	id := c.Params.ByName("id")

	res := db.GlobalDB.Where("id = ?", id).First(&haul)

	if res.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "haul not found",
		})
		c.Abort()
		return
	}

	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not get haul",
		})
		c.Abort()
		return
	}

	if err := c.ShouldBindJSON(&haul); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	db.GlobalDB.Save(&haul)

	c.JSON(http.StatusOK, gin.H{"data": haul})

	return
}

func DeleteHaulByID(c *gin.Context) {
	var haul models.Haul

	id := c.Params.ByName("id")

	res := db.GlobalDB.Where("id = ?", id).Delete(&haul)

	if res.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "haul not found",
		})
		c.Abort()
		return
	}

	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not remove haul",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "haul with id " + string(id) + " removed."})
}
