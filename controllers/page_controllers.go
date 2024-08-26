package controllers

import (
	"lashes/models"
	"lashes/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPages(c *gin.Context) {
	pages, err := service.GetAllPages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pages)
}

func GetPage(c *gin.Context) {
	slug := c.Param("slug")
	page, err := service.GetPageBySlug(slug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
		return
	}
	c.JSON(http.StatusOK, page)
}

func CreatePage(c *gin.Context) {
	var page models.Page
	if err := c.ShouldBindJSON(&page); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.CreatePage(&page); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, page)
}

func UpdatePage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedPage models.Page
	if err := c.ShouldBindJSON(&updatedPage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.UpdatePage(uint(id), updatedPage); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Page updated successfully"})
}

func DeletePage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := service.DeletePage(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Page deleted successfully"})
}
