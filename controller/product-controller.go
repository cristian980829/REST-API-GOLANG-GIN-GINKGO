package controller

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/cristian980829/PRUEBA-TECNICA-KUASAR/entity"
	"github.com/cristian980829/PRUEBA-TECNICA-KUASAR/service"
)

type ProductController interface {
	FindAll() []entity.Product
	FindOne(ctx *gin.Context) entity.Product
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	AlreadyExist(ctx *gin.Context) bool
	GetVolumes() interface{}
}

type controller struct {
	service service.ProductService
}

func New(service service.ProductService) ProductController {
	return &controller{
		service: service,
	}
}

// POST /products
// Create new product
func (c *controller) Save(ctx *gin.Context) error {
	var product entity.Product

	// Get the product from the request
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		return err
	}

	// Get current date
	t := time.Now()
	date := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// Assign the current date
	product.Creation_date = date
	product.Update_date = date

	// Product saved in database
	errorr := c.service.Save(product)
	return errorr
}

// PUT /products/:id
// Update a product
func (c *controller) Update(ctx *gin.Context) error {

	// Get id param
	id := ctx.Param("id")

	var product entity.Product

	// Get the product from the request
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		return err
	}

	// Find the record and get the creation date
	prod := c.service.FindOne(id)
	product.Creation_date = prod.Creation_date

	// Assign the id param
	product.ID = id

	// Get current date
	t := time.Now()
	date := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// Assign the current date
	product.Update_date = date

	// Product updated in database
	errorr := c.service.Update(product)
	return errorr
}

// DELETE /products/:id
// Delete a product
func (c *controller) Delete(ctx *gin.Context) error {
	var product entity.Product

	// Get id param
	id := ctx.Param("id")

	// Assign the id param
	product.ID = id

	// Product deleted in database
	c.service.Delete(product)
	return nil
}

// Validate if the record already exists
func (c *controller) AlreadyExist(ctx *gin.Context) bool {
	// Get id param
	id := ctx.Param("id")
	return c.service.AlreadyExist(id)
}

// GET /products
// Find all products
func (c *controller) FindAll() []entity.Product {
	// Return all databse products
	return c.service.FindAll()
}

// GET /products/:id
// Find a product
func (c *controller) FindOne(ctx *gin.Context) entity.Product {
	// Get id param
	id := ctx.Param("id")
	// Return one database product
	return c.service.FindOne(id)
}

// GET /volumes
// Find volumes
func (c *controller) GetVolumes() interface{} {
	return c.service.Volumes()
}
