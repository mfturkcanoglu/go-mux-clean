package services

import (
	"strconv"

	"github.com/mfturkcanoglu/go-mux-clean/cache"
	"github.com/mfturkcanoglu/go-mux-clean/entities"
	"github.com/mfturkcanoglu/go-mux-clean/repositories"

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
	productCache      cache.ProductCache             = cache.NewRedisCache("localhost:6379", 0, 30)
)

func (productService) GetAll() ([]entities.Product, error) {
	return productRepository.FindAll()
}

func (productService) GetByID(id string) (entities.Product, error) {
	if product := productCache.Get(id); product != nil {
		return *product, nil
	}
	product, err := productRepository.FindByID(convertStringToUnit(id))
	productCache.Set(id, &product)
	return product, err
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
