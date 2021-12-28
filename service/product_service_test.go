package service_test

import (
	"github.com/cristian980829/PRUEBA-TECNICA-KUASAR/entity"
	"github.com/cristian980829/PRUEBA-TECNICA-KUASAR/repository"
	"github.com/cristian980829/PRUEBA-TECNICA-KUASAR/service"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	ID                       = "1233456"
	NAME                     = "PRODUCT NAME"
	DESCRIPTION              = "PRODUCT DESCRIPTION"
	STATUS                   = "PRODUCT STATUS"
	CREATION_DATE            = "2021-12-27T20:36:13"
	UPDATE_DATE              = "2021-12-27T20:36:13"
	ACCOUNT_ID               = "2"
	FORMAT_PRODUCT           = "FORMAT PRODUCT"
	VALUE_UNIT       float32 = 2000.00
	UNIT_NAME                = "PRODUCT UNIT NAME"
	UNIT_DESCRIPTION         = "PRODUCT UNIT DESCRIPTION"
	STOCK            int32   = 100
)

var testProduct entity.Product = entity.Product{
	ID:               ID,
	Name:             NAME,
	Description:      DESCRIPTION,
	Status:           STATUS,
	Creation_date:    CREATION_DATE,
	Update_date:      UPDATE_DATE,
	Account_id:       ACCOUNT_ID,
	Format_product:   FORMAT_PRODUCT,
	Value_unit:       VALUE_UNIT,
	Unit_name:        UNIT_NAME,
	Unit_description: UNIT_DESCRIPTION,
	Stock:            STOCK,
}

var _ = Describe("ProductService", func() {

	var (
		productRepository repository.ProductRepository
		productService    service.ProductService
	)

	BeforeSuite(func() {
		productRepository = repository.NewProductRepository()
		productService = service.New(productRepository)
	})

	Describe("Fetching all existing products", func() {

		Context("If there is a product in the database", func() {

			BeforeEach(func() {
				productService.Save(testProduct)
			})

			It("should return at least one element", func() {
				productList := productService.FindAll()

				Ω(productList).ShouldNot(BeEmpty())
			})

			It("should map the fields correctly", func() {
				firstProduct := productService.FindAll()[0]

				Ω(firstProduct.ID).Should(Equal(ID))
				Ω(firstProduct.Name).Should(Equal(NAME))
				Ω(firstProduct.Description).Should(Equal(DESCRIPTION))
				Ω(firstProduct.Status).Should(Equal(STATUS))
				Ω(firstProduct.Creation_date).Should(Equal(CREATION_DATE))
				Ω(firstProduct.Update_date).Should(Equal(UPDATE_DATE))
				Ω(firstProduct.Account_id).Should(Equal(ACCOUNT_ID))
				Ω(firstProduct.Format_product).Should(Equal(FORMAT_PRODUCT))
				Ω(firstProduct.Value_unit).Should(Equal(VALUE_UNIT))
				Ω(firstProduct.Unit_name).Should(Equal(UNIT_NAME))
				Ω(firstProduct.Unit_description).Should(Equal(UNIT_DESCRIPTION))
				Ω(firstProduct.Stock).Should(Equal(STOCK))
			})

			AfterEach(func() {
				product := productService.FindAll()[0]
				productService.Delete(product)
			})

		})

		Context("If there are no products in the database", func() {

			It("should return an empty list", func() {
				products := productService.FindAll()

				Ω(products).Should(BeEmpty())
			})

		})
	})

	AfterSuite(func() {
		productRepository.CloseDB()
	})

})
