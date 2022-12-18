package entities

type Product struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name" validate:"required,min=5,max=20"`
	Price       float64 `json:"price" validate:"required,min=5.0"`
	Description string  `json:"description" validate:"required,min=10,max=50"`
}
