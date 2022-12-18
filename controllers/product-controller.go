package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"standart-lib-rest-api/entities"
	"standart-lib-rest-api/services"

	"github.com/gorilla/mux"
)

type ProductController interface {
	CreateProduct(w http.ResponseWriter, r *http.Request)
	GetProduct(w http.ResponseWriter, r *http.Request)
	GetAllProducts(w http.ResponseWriter, r *http.Request)
	DeleteProduct(w http.ResponseWriter, r *http.Request)
	UpdateProduct(w http.ResponseWriter, r *http.Request)
}

type productController struct{}

func NewProductController() ProductController {
	return &productController{}
}

var (
	productService services.ProductService = services.NewProductService()
)

func (productController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var product entities.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	product, err = productService.Create(product)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (productController) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	product, err := productService.GetByID(id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (productController) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	products, err := productService.GetAll()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (productController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var id = mux.Vars(r)["id"]
	result, err := productService.Delete(id)
	if err != nil {
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(fmt.Sprintf(`{"success":%v}`, result))
}

func (productController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var newProduct entities.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		log.Println("Decoding error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	newProduct, err = productService.Update(mux.Vars(r)["id"], newProduct)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(newProduct)
}
