package service

import (
	"lashes/db"
	"lashes/models"
)

func GetAllContacts() ([]models.Contact, error) {
	var contacts []models.Contact
	result := db.DB.Find(&contacts)
	return contacts, result.Error
}

func GetContactByID(id uint) (models.Contact, error) {
	var contact models.Contact
	result := db.DB.First(&contact, id)
	return contact, result.Error
}

func CreateContact(contact *models.Contact) error {
	result := db.DB.Create(contact)
	return result.Error
}

func UpdateContact(id uint, updatedContact models.Contact) error {
	var contact models.Contact
	if err := db.DB.First(&contact, id).Error; err != nil {
		return err
	}
	result := db.DB.Model(&contact).Updates(updatedContact)
	return result.Error
}

func DeleteContact(id uint) error {
	var contact models.Contact
	result := db.DB.Delete(&contact, id)
	return result.Error
}
