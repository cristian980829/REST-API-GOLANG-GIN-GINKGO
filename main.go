package main

import (
	"github.com/gin-gonic/gin"
	"github.com/cristian980829/PRUEBA-TECNICA-KUASAR/controllers"
	"github.com/cristian980829/PRUEBA-TECNICA-KUASAR/models"
	"github.com/cristian980829/PRUEBA-TECNICA-KUASAR/helper"
)

func main() {

	router := gin.Default()

	// Connected to database
	models.ConnectDatabase()

	// Routes
	router.POST("/api/products", controllers.CreateProduct)
	router.GET("/api/products", controllers.FindProducts)
	router.GET("/api/products/:id", controllers.FindOneProduct)
	router.PUT("/api/products/:id", controllers.UpdateProduct)
	router.DELETE("/api/products/:id", controllers.DeleteProduct)

	// Check if is an authorized user
	authorized := router.Group("/", gin.BasicAuth(helper.ValidUser()))
	authorized.GET("/api/volumes", controllers.GetVolumes)
	
	// Custom port
	listenPort := "9098"
	
	// Run port
	router.Run(":"+listenPort)
}