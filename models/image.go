package models

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	ImageURL string
	PageID   uint
}