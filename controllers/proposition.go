package controllers

import (
	"log"
	"net/http"

	"github.com/PatrikOlin/lp-api/db"
	"github.com/PatrikOlin/lp-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllPropositions(c *gin.Context) {
	var propo []models.Proposition

	res := db.GlobalDB.Find(&propo)

	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not get pickups",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": propo,
	})
}

func GetPropositionsByUserID(c *gin.Context) {
	var propos []models.Proposition

	userID := c.Params.ByName("id")

	res := db.GlobalDB.Where("user_id = ?", userID).Find(&propos)

	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not get users propositions",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": propos,
	})
	return
}

func CreateProposition(c *gin.Context) {
	var propo models.Proposition
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

	propo.UserID = user.ID

	err := c.ShouldBindJSON(&propo)
	if err != nil {
		log.Println(err)

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json",
		})
		c.Abort()
		return
	}

	err = propo.CreatePropositionRecord()
	if err != nil {
		log.Println(err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error creating pickup",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": propo,
	})
}

func UpdatePropositionByID(c *gin.Context) {
	var propos models.Proposition
	id := c.Params.ByName("id")

	res := db.GlobalDB.Where("id = ?", id).First(&propos)

	if res.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "proposition not found",
		})
		c.Abort()
		return
	}

	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not get proposition",
		})
		c.Abort()
		return
	}

	if err := c.ShouldBindJSON(&propos); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	db.GlobalDB.Save(&propos)

	c.JSON(http.StatusOK, gin.H{"data": propos})

	return
}

func DeletePropositionByID(c *gin.Context) {
	var propos models.Proposition

	id := c.Params.ByName("id")

	res := db.GlobalDB.Where("id = ?", id).Delete(&propos)

	if res.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "proposition not found",
		})
		c.Abort()
		return
	}

	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not remove proposition",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "proposition with id " + string(id) + " removed."})
}
