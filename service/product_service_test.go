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

const (
	UID                       = "1233456"
	UNAME                     = "UPDATE PRODUCT NAME"
	UDESCRIPTION              = "UPDATE PRODUCT DESCRIPTION"
	USTATUS                   = "UPDATE PRODUCT STATUS"
	UCREATION_DATE            = "2022-12-27T20:36:13"
	UUPDATE_DATE              = "2022-12-27T20:36:13"
	UACCOUNT_ID               = "4"
	UFORMAT_PRODUCT           = "UPDATE FORMAT PRODUCT"
	UVALUE_UNIT       float32 = 4000.00
	UUNIT_NAME                = "UPDATE PRODUCT UNIT NAME"
	UUNIT_DESCRIPTION         = "UPDATE PRODUCT UNIT DESCRIPTION"
	USTOCK            int32   = 200
)

var testUpdatedProduct entity.Product = entity.Product{
	ID:               UID,
	Name:             UNAME,
	Description:      UDESCRIPTION,
	Status:           USTATUS,
	Creation_date:    UCREATION_DATE,
	Update_date:      UUPDATE_DATE,
	Account_id:       UACCOUNT_ID,
	Format_product:   UFORMAT_PRODUCT,
	Value_unit:       UVALUE_UNIT,
	Unit_name:        UUNIT_NAME,
	Unit_description: UUNIT_DESCRIPTION,
	Stock:            USTOCK,
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

	Describe("Adding new products", func() {

		Context("If a new product is added to the database", func() {

			It("Should return nil", func() {
				product := productService.Save(testProduct)

				Ω(product).Should(BeNil())
			})

			AfterEach(func() {
				product := productService.FindAll()[0]
				productService.Delete(product)
			})

		})

		Context("If a repeating product is added to the database", func() {

			BeforeEach(func() {
				productService.Save(testProduct)
			})

			It("Should not return nil", func() {
				product := productService.Save(testProduct)

				Ω(product).ShouldNot(BeNil())
			})

			AfterEach(func() {
				product := productService.FindAll()[0]
				productService.Delete(product)
			})

		})

	})

	Describe("Update products", func() {

		Context("If a product is updated in the database", func() {

			BeforeEach(func() {
				productService.Save(testProduct)
			})

			It("Should update the element", func() {
				// Element updated
				productService.Update(testUpdatedProduct)

				// Find element
				updatedProduct := productService.FindOne(ID)

				// Check if item was updated
				Ω(updatedProduct.ID).Should(Equal(UID))
				Ω(updatedProduct.Name).Should(Equal(UNAME))
				Ω(updatedProduct.Description).Should(Equal(UDESCRIPTION))
				Ω(updatedProduct.Status).Should(Equal(USTATUS))
				Ω(updatedProduct.Creation_date).Should(Equal(UCREATION_DATE))
				Ω(updatedProduct.Update_date).Should(Equal(UUPDATE_DATE))
				Ω(updatedProduct.Account_id).Should(Equal(UACCOUNT_ID))
				Ω(updatedProduct.Format_product).Should(Equal(UFORMAT_PRODUCT))
				Ω(updatedProduct.Value_unit).Should(Equal(UVALUE_UNIT))
				Ω(updatedProduct.Unit_name).Should(Equal(UUNIT_NAME))
				Ω(updatedProduct.Unit_description).Should(Equal(UUNIT_DESCRIPTION))
				Ω(updatedProduct.Stock).Should(Equal(USTOCK))

			})

			AfterEach(func() {
				product := productService.FindAll()[0]
				productService.Delete(product)
			})
		})

	})

	Describe("Deleting products", func() {

		Context("If a repeating product is added to the database", func() {

			BeforeEach(func() {
				productService.Save(testProduct)
			})

			It("Should return nil", func() {
				product := productService.Delete(testProduct)

				Ω(product).Should(BeNil())
			})

		})

	})

	Describe("Fetching all existing products", func() {

		Context("If there is a product in the database", func() {

			BeforeEach(func() {
				productService.Save(testProduct)
			})

			It("Should return at least one element", func() {
				productList := productService.FindAll()

				Ω(productList).ShouldNot(BeEmpty())
			})

			It("Should map the fields correctly", func() {
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

			It("Should return an empty list", func() {
				products := productService.FindAll()

				Ω(products).Should(BeEmpty())
			})

			It("Should not find the element", func() {
				product := productService.AlreadyExist(ID)

				Ω(product).Should(Equal(false))
			})

		})

		Context("If there are no data in the JSON file", func() {

			It("Should return nil", func() {
				volumes := productService.Volumes()

				Ω(volumes).Should(BeNil())
			})

		})
	})

	AfterSuite(func() {
		productRepository.CloseDB()
	})

})
