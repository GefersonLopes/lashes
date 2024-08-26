package main

import (
	"lashes/db"
	"lashes/routes"
	"lashes/models"
	"log"
)

func main() {
	db.Connect()

	err := db.DB.AutoMigrate(&models.Page{}, &models.Image{}, &models.Contact{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Migration complete.")

	r := routes.SetupRouter()
	r.Run()
}
