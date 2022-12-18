package services

import (
	"standart-lib-rest-api/entities"
	"standart-lib-rest-api/repositories"
	"strconv"

	"github.com/go-playground/validator/v10"
)

type ProductService interface {
	GetAll() ([]entities.Product, error)
	GetByID(string) (entities.Product, error)
	Create(entities.Product) (entities.Product, error)
	Update(string, entities.Product) (entities.Product, error)
	Delete(string) (bool, error)
}

type productService struct{}

func NewProductService() ProductService {
	return &productService{}
}

var (
	productRepository repositories.ProductRepository = repositories.NewProductRepository()
)

func (productService) GetAll() ([]entities.Product, error) {
	return productRepository.FindAll()
}

func (productService) GetByID(id string) (entities.Product, error) {
	return productRepository.FindByID(convertStringToUnit(id))
}

func (productService) Create(product entities.Product) (entities.Product, error) {
	if err := Validate(product); err != nil {
		return product, err
	}
	return productRepository.Save(product)
}

func (productService) Update(id string, newProduct entities.Product) (entities.Product, error) {
	if err := Validate(newProduct); err != nil {
		return newProduct, err
	}
	return productRepository.Update(convertStringToUnit(id), newProduct)
}

func (productService) Delete(id string) (bool, error) {
	return productRepository.Delete(convertStringToUnit(id))
}

func Validate(product entities.Product) error {
	validate := validator.New()
	return validate.Struct(product)
}

func convertStringToUnit(str string) uint {
	var ui, _ = strconv.Atoi(str)
	return uint(ui)
}
