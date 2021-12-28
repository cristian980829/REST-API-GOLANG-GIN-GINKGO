package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cristian980829/PRUEBA-TECNICA-KUASAR/controller"
	"github.com/cristian980829/PRUEBA-TECNICA-KUASAR/helper"
)

type ProductApi struct {
	productController controller.ProductController
}

func NewProductAPI(productController controller.ProductController) *ProductApi {
	return &ProductApi{
		productController: productController,
	}
}

func (api *ProductApi) CreateProduct(ctx *gin.Context) {

	err := api.productController.Save(ctx)
	if err != nil {
		// Response to request
		ctx.JSON(http.StatusBadRequest, &helper.Response{
			Message: err.Error(),
			Ok:      false,
		})
	} else {
		// Response to request
		ctx.JSON(http.StatusOK, &helper.Response{
			Message: "Success!",
			Ok:      true,
		})
	}
}

func (api *ProductApi) GetProducts(ctx *gin.Context) {
	// Response to request
	ctx.JSON(200, api.productController.FindAll())
}

func (api *ProductApi) GetProduct(ctx *gin.Context) {
	exist := api.productController.AlreadyExist(ctx)
	if exist {
		// Response to request
		ctx.JSON(200, api.productController.FindOne(ctx))
	} else {
		// Response to request
		ctx.JSON(http.StatusBadRequest, &helper.Response{
			Message: "Record does not exist",
			Ok:      false,
		})
	}
}

func (api *ProductApi) UpdateProduct(ctx *gin.Context) {
	exist := api.productController.AlreadyExist(ctx)
	if !exist {
		// Response to request
		ctx.JSON(http.StatusBadRequest, &helper.Response{
			Message: "Record does not exist",
			Ok:      false,
		})
	} else {
		err := api.productController.Update(ctx)
		if err != nil {
			// Response to request
			ctx.JSON(http.StatusBadRequest, &helper.Response{
				Message: err.Error(),
				Ok:      false,
			})
		} else {
			// Response to request
			ctx.JSON(http.StatusOK, &helper.Response{
				Message: "Success!",
				Ok:      true,
			})
		}
	}
}

func (api *ProductApi) DeleteProduct(ctx *gin.Context) {
	exist := api.productController.AlreadyExist(ctx)
	if !exist {
		// Response to request
		ctx.JSON(http.StatusBadRequest, &helper.Response{
			Message: "Record does not exist",
			Ok:      false,
		})
	} else {
		err := api.productController.Delete(ctx)
		if err != nil {
			// Response to request
			ctx.JSON(http.StatusBadRequest, &helper.Response{
				Message: err.Error(),
				Ok:      false,
			})
		} else {
			// Response to request
			ctx.JSON(http.StatusOK, &helper.Response{
				Message: "Success!",
				Ok:      true,
			})
		}
	}
}

func (api *ProductApi) GetVolumes(ctx *gin.Context) {
	// Response to request
	ctx.JSON(200, api.productController.GetVolumes())
}
