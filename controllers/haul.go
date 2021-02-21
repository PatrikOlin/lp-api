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
