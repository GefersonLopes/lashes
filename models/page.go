package models

import "gorm.io/gorm"

type Page struct {
    gorm.Model
    Title    string
    Slug     string
    Content  string
    Images   []Image   `gorm:"foreignKey:PageID"`
    Contacts []Contact `gorm:"foreignKey:PageID"`
}
