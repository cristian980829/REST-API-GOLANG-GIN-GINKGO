package main

import (
	"github.com/cristian980829/PRUEBA-TECNICA-KUASAR/api"
	"github.com/cristian980829/PRUEBA-TECNICA-KUASAR/controller"
	"github.com/cristian980829/PRUEBA-TECNICA-KUASAR/helper"
	"github.com/cristian980829/PRUEBA-TECNICA-KUASAR/repository"
	"github.com/cristian980829/PRUEBA-TECNICA-KUASAR/service"

	"github.com/gin-gonic/gin"
)

var (
	productRepository repository.ProductRepository = repository.NewProductRepository()
	productService    service.ProductService       = service.New(productRepository)
	productController controller.ProductController = controller.New(productService)
)

func main() {

	router := gin.Default()

	productAPI := api.NewProductAPI(productController)

	// Routes
	router.POST("/api/products", productAPI.CreateProduct)
	router.GET("/api/products", productAPI.GetProducts)
	router.GET("/api/products/:id", productAPI.GetProduct)
	router.PUT("/api/products/:id", productAPI.UpdateProduct)
	router.DELETE("/api/products/:id", productAPI.DeleteProduct)

	// Check if is an authorized user
	authorized := router.Group("/", gin.BasicAuth(helper.ValidUser()))
	authorized.GET("/api/volumes", productAPI.GetVolumes)

	// Custom port
	listenPort := "9098"

	// Run port
	router.Run(":" + listenPort)
}
