package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Address string
	Phone   string
	PageID  uint
}