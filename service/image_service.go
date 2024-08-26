package service

import (
	"lashes/db"
	"lashes/models"
)

func GetAllImages() ([]models.Image, error) {
	var images []models.Image
	result := db.DB.Find(&images)
	return images, result.Error
}

func GetImageByID(id uint) (models.Image, error) {
	var image models.Image
	result := db.DB.First(&image, id)
	return image, result.Error
}

func CreateImage(image *models.Image) error {
	result := db.DB.Create(image)
	return result.Error
}

func UpdateImage(id uint, updatedImage models.Image) error {
	var image models.Image
	if err := db.DB.First(&image, id).Error; err != nil {
		return err
	}
	result := db.DB.Model(&image).Updates(updatedImage)
	return result.Error
}

func DeleteImage(id uint) error {
	var image models.Image
	result := db.DB.Delete(&image, id)
	return result.Error
}
