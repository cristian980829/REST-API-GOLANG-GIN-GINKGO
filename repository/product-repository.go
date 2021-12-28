package repository

import (
	"encoding/json"
	"io/ioutil"

	"github.com/cristian980829/PRUEBA-TECNICA-KUASAR/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type ProductRepository interface {
	Save(product entity.Product) error
	Update(product entity.Product) error
	Delete(product entity.Product) error
	FindAll() []entity.Product
	FindOne(string) entity.Product
	AlreadyExist(string) bool
	Volumes() interface{}
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewProductRepository() ProductRepository {
	db, err := gorm.Open("sqlite3", "store.db")
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&entity.Product{})
	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("Failed to close database")
	}
}

func (db *database) Save(product entity.Product) error {
	err := db.connection.Create(&product).Error
	return err
}

func (db *database) Update(product entity.Product) error {
	err := db.connection.Save(&product).Error
	return err
}

func (db *database) Delete(product entity.Product) error {
	err := db.connection.Delete(&product).Error
	return err
}

func (db *database) FindAll() []entity.Product {
	var products []entity.Product
	db.connection.Set("gorm:auto_preload", true).Find(&products)
	return products
}

func (db *database) FindOne(id string) entity.Product {
	var product entity.Product
	db.connection.Where("id = ?", id).First(&product)
	return product
}

func (db *database) AlreadyExist(id string) bool {
	var product entity.Product
	if err := db.connection.Where("id = ?", id).First(&product).Error; err != nil {
		return false
	}
	return true
}

func (db *database) Volumes() interface{} {
	var data interface{}

	// Read file
	volumen_list, _ := ioutil.ReadFile("./helper/volumen_list.json")

	// JSON is converted
	json.Unmarshal(volumen_list, &data)

	return data
}
