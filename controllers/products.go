package controllers

import (
	"net/http"
	"time"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/cristian980829/PRUEBA-TECNICA-KUASAR/models"

)

type CreateProductInput struct {
	Name  				string 	`json:"name" binding:"required"`
	Description 		string 	`json:"description" binding:"required"`
	Status				string 	`json:"status" binding:"required"`
	Account_id			string 	`json:"account_id" binding:"required"`
	Format_product		string 	`json:"format_product" binding:"required"`
	Value_unit			float32 `json:"value_unit" binding:"required"`
	Unit_name			string 	`json:"unit_name" binding:"required"`
	Unit_description 	string 	`json:"unit_description" binding:"required"`
	Stock				int32 	`json:"stock" binding:"required"`
}


// POST /products
// Create new product
func CreateProduct(c *gin.Context) {

	// Validate input
	var input CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get current date
	t := time.Now()
	date := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())


	// Create product
	product := models.Product{
		Name: input.Name,
		Description: input.Description,
		Status: input. Status,
		Creation_date: date,
		Update_date: date,
		Account_id: input.Account_id,
		Format_product: input.Format_product,
		Value_unit: input.Value_unit,
		Unit_name: input.Unit_name,
		Unit_description: input.Unit_description,
		Stock: input.Stock,
	}

	// Save product in database
	models.DB.Create(&product)

	// Response to request
	c.JSON(http.StatusOK, gin.H{"data": product})
}

// GET /products
// Find all products
func FindProducts(c *gin.Context) {

	var products []models.Product
	//Get all database products
	models.DB.Find(&products)

	// Response to request
	c.JSON(http.StatusOK, gin.H{"data": products})
}

// GET /products/:id
// Find a product
func FindOneProduct(c *gin.Context) {

	var product models.Product

	// Get product if exist
	if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Response to request
	c.JSON(http.StatusOK, gin.H{"data": product})
}

// PATCH /products/:id
// Update a product
func UpdateProduct(c *gin.Context) {

	var product models.Product
	// Get product if exist
	if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get current date
	t := time.Now()
	date := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// Update_date will be updated
	product.Update_date = date

	// Product updated
	models.DB.Save(&product)

	// Response to request
	c.JSON(http.StatusOK, gin.H{"data": product})

}