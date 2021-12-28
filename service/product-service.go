package service

import (
	"github.com/cristian980829/PRUEBA-TECNICA-KUASAR/entity"
	"github.com/cristian980829/PRUEBA-TECNICA-KUASAR/repository"
)

type ProductService interface {
	Save(entity.Product) error
	Update(entity.Product) error
	Delete(entity.Product) error
	FindAll() []entity.Product
	FindOne(string) entity.Product
	AlreadyExist(string) bool
	Volumes() interface{}
}

type productService struct {
	repository repository.ProductRepository
}

func New(productRepository repository.ProductRepository) ProductService {
	return &productService{
		repository: productRepository,
	}
}

func (service *productService) Save(product entity.Product) error {
	error := service.repository.Save(product)
	return error
}

func (service *productService) Update(product entity.Product) error {
	error := service.repository.Update(product)
	return error
}

func (service *productService) Delete(product entity.Product) error {
	error := service.repository.Delete(product)
	return error
}

func (service *productService) FindOne(id string) entity.Product {
	return service.repository.FindOne(id)
}

func (service *productService) FindAll() []entity.Product {
	return service.repository.FindAll()
}

func (service *productService) AlreadyExist(id string) bool {
	return service.repository.AlreadyExist(id)
}

func (service *productService) Volumes() interface{} {
	return service.repository.Volumes()
}
