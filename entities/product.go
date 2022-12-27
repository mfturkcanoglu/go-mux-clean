package entities

// Product structure
// swagger:response productModel
type Product struct {
	// the id for this product
	//
	// required: true
	// min: 1
	ID uint `json:"id"`
	// the name for this product
	// required: true
	// min length: 5
	// max length: 20
	Name string `json:"name" validate:"required,min=5,max=20"`
	// the price of this product
	// required: true
	// min: 5.0
	Price float64 `json:"price" validate:"required,min=5.0"`
	// the name for this product
	// required: true
	// min length: 10
	// max length: 50
	Description string `json:"description" validate:"required,min=10,max=50"`
}
