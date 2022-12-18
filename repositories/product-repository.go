package repositories

import (
	"log"
	"standart-lib-rest-api/database"
	"standart-lib-rest-api/entities"
)

type ProductRepository interface {
	FindAll() ([]entities.Product, error)
	FindByID(uint) (entities.Product, error)
	Save(entities.Product) (entities.Product, error)
	Update(uint, entities.Product) (entities.Product, error)
	Delete(uint) (bool, error)
}

type repository struct{}

func NewProductRepository() ProductRepository {
	return &repository{}
}

func (repository) FindAll() ([]entities.Product, error) {
	var products []entities.Product
	result := database.Instance.Find(&products)
	return products, result.Error
}

func (repository) FindByID(id uint) (entities.Product, error) {
	var product entities.Product
	result := database.Instance.First(&product, id)
	return product, result.Error
}

func (repository) Save(newProduct entities.Product) (entities.Product, error) {
	result := database.Instance.Create(&newProduct)
	return newProduct, result.Error
}

func (repository) Update(id uint, newProduct entities.Product) (entities.Product, error) {
	var oldProduct entities.Product
	var err error
	if result := database.Instance.First(&oldProduct, id); result.Error != nil {
		log.Println("Finding error", result.Error)
		err = result.Error
	}
	if result := database.Instance.Model(&oldProduct).Updates(&newProduct); result.Error != nil {
		log.Println("Update error", result.Error)
		err = result.Error
	}
	newProduct.ID = oldProduct.ID
	return newProduct, err
}

func (repository) Delete(id uint) (bool, error) {
	if result := database.Instance.Delete(&entities.Product{}, id); result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
