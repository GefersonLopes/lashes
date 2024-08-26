package service

import (
	"lashes/db"
	"lashes/models"
)

func GetAllPages() ([]models.Page, error) {
	var pages []models.Page
	result := db.DB.Find(&pages)
	return pages, result.Error
}

func GetPageBySlug(slug string) (models.Page, error) {
	var page models.Page
	result := db.DB.Where("slug = ?", slug).First(&page)
	return page, result.Error
}

func CreatePage(page *models.Page) error {
	result := db.DB.Create(page)
	return result.Error
}

func UpdatePage(id uint, updatedPage models.Page) error {
	var page models.Page
	if err := db.DB.First(&page, id).Error; err != nil {
		return err
	}
	result := db.DB.Model(&page).Updates(updatedPage)
	return result.Error
}

func DeletePage(id uint) error {
	var page models.Page
	result := db.DB.Delete(&page, id)
	return result.Error
}
