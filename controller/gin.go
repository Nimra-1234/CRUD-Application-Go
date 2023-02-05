package controllers

import (
	"errors"
	"net/http"

	"example.com/go-crud-api/database"
	"github.com/gin-gonic/gin"
)

func ReadMed(c *gin.Context) {
	var med database.Medicine
	id := c.Param("id")
	res := database.DB.Find(&med, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Medicine not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"med": med,
	})
	return
}

func ReadMeds(c *gin.Context) {
	var meds []database.Medicine
	res := database.DB.Find(&meds)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": errors.New("company not found"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"meds": meds,
	})
	return
}

func UpdateMed(c *gin.Context) {
	var med database.Medicine
	id := c.Param("id")
	err := c.ShouldBind(&med)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	var updateMed database.Medicine
	res := database.DB.Model(&updateMed).Where("id = ?", id).Updates(med)

	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "medicine not updated",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"med": med,
	})
	return
}

func DeleteMed(c *gin.Context) {
	var med database.Medicine
	id := c.Param("id")
	res := database.DB.Find(&med, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "medicine not found",
		})
		return
	}
	database.DB.Delete(&med)
	c.JSON(http.StatusOK, gin.H{
		"message": "medicine deleted successfully",
	})
	return
}
func CreateMed(c *gin.Context) {
	var med *database.Medicine
	err := c.ShouldBind(&med)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res := database.DB.Create(med)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error creating a medicine",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"med": med,
	})
	return
}
