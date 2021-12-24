package main

import (
	"github.com/gin-gonic/gin"
	"github.com/cristian980829/PRUEBA-TECNICA-KUASAR/controllers"
	"github.com/cristian980829/PRUEBA-TECNICA-KUASAR/models"
)

func main() {

	router := gin.Default()

	// Connected to database
	models.ConnectDatabase()

	// Routes
	router.POST("/api/products", controllers.CreateProduct)
	router.GET("/api/products", controllers.FindProducts)
	router.GET("/api/products/:id", controllers.FindOneProduct)
	
	// Custom port
	listenPort := "9098"
	
	// Run port
	router.Run(":"+listenPort)
}