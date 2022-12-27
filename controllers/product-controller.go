// Package Product API
//
// Schemes: http, https
// Host: localhost
// BasePath: v2
// Version: 0.0.1
//
// Consumes:
//   - application/json
//
// Produces:
// - application/json
//
//swagger:meta
package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mfturkcanoglu/go-mux-clean/entities"
	"github.com/mfturkcanoglu/go-mux-clean/services"

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

// temp
// swagger:response
type badRequest struct{}

// swagger:route GET /products products saveProduct
//
// This will show all available products by default.
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	200: productModel
//	400: badRequest
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
		http.Error(w, err.Error(), http.StatusBadRequest)
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(newProduct)
}
