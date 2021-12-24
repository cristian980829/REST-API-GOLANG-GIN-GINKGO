package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/cristian980829/PRUEBA-TECNICA-KUASAR/models"

)

type CreateProductInput struct {
	Name  				string 	`json:"name" binding:"required"`
	Description 		string 	`json:"description" binding:"required"`
	Status				string 	`json:"status" binding:"required"`
	Creation_date		string 	`json:"creation_date" binding:"required"`
	Update_date			string 	`json:"update_date" binding:"required"`
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


	// Create product
	product := models.Product{
		Name: input.Name,
		Description: input.Description,
		Status: input. Status,
		Creation_date: input.Creation_date,
		Update_date: input.Update_date,
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
