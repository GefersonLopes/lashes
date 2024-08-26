package controllers

import (
	"lashes/service"
	"lashes/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetImages(c *gin.Context) {
	images, err := service.GetAllImages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, images)
}

func GetImage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	image, err := service.GetImageByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}
	c.JSON(http.StatusOK, image)
}

func CreateImage(c *gin.Context) {
	var image models.Image
	if err := c.ShouldBindJSON(&image); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.CreateImage(&image); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, image)
}

func UpdateImage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedImage models.Image
	if err := c.ShouldBindJSON(&updatedImage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.UpdateImage(uint(id), updatedImage); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image updated successfully"})
}

func DeleteImage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := service.DeleteImage(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image deleted successfully"})
}
