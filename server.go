package main

import (
	"standart-lib-rest-api/controllers"
	"standart-lib-rest-api/database"
	router "standart-lib-rest-api/http"
)

var (
	httpRouter        router.Router                 = router.NewMuxRouter()
	productController controllers.ProductController = controllers.NewProductController()
)

func main() {
	// Load configs from config.json with Viper
	LoadAppConfig()

	// Initialize and migrate db
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	RegisterProductRoutes(httpRouter)
	httpRouter.SERVE(AppConfig.Port)
}

func RegisterProductRoutes(router router.Router) {
	router.GET("/products/{id}", productController.GetProduct)
	router.GET("/products", productController.GetAllProducts)
	router.POST("/products", productController.CreateProduct)
	router.PATCH("/products/{id}", productController.UpdateProduct)
	router.DELETE("/products/{id}", productController.DeleteProduct)
}
