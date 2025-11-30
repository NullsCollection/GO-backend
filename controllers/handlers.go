package controllers

import (
	"backend/connection"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//READ ALL
func GetProjects(c *gin.Context) {
	var items []models.Projects
	
	if err := connection.DB.Find(&items).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No items found"})
		return
	}

	c.JSON(http.StatusOK, items)
}

//CREATE
func CreateProjects(c *gin.Context) {

	var item models.Projects
	
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := connection.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item"})
		return
	}

	c.JSON(http.StatusCreated, item)
}

//READ ONE
func GetProjectID(c *gin.Context) {
	id := c.Param("id")
	var item models.Projects
	
	if err := connection.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, item)
}


//UPDATE
func UpdateProjects(c *gin.Context) {
	id := c.Param("id")
	var item models.Projects
	
	if err := connection.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	var input models.Projects
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	connection.DB.Model(&item).Updates(input)
	c.JSON(http.StatusOK, item)
}

//DELETE
func DeleteProjects(c *gin.Context) {
	id := c.Param("id")
	var item models.Projects
	
	if err := connection.DB.Delete(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Delete Failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
}
