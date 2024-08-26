package routes

import (
	"lashes/controllers"
	"lashes/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.CORSMiddleware())

	// Rotas Pages
	r.GET("/pages", controllers.GetPages)
	r.GET("/pages/:slug", controllers.GetPage)
	r.POST("/pages", controllers.CreatePage)
	r.PATCH("/pages/:id", controllers.UpdatePage)
	r.DELETE("/pages/:id", controllers.DeletePage)

	// Rotas Images
	r.GET("/images", controllers.GetImages)
	r.GET("/images/:id", controllers.GetImage)
	r.POST("/images", controllers.CreateImage)
	r.PATCH("/images/:id", controllers.UpdateImage)
	r.DELETE("/images/:id", controllers.DeleteImage)

	// Rotas Contacts
	r.GET("/contacts", controllers.GetContacts)
	r.GET("/contacts/:id", controllers.GetContact)
	r.POST("/contacts", controllers.CreateContact)
	r.PATCH("/contacts/:id", controllers.UpdateContact)
	r.DELETE("/contacts/:id", controllers.DeleteContact)

	return r
}
