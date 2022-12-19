package cache

import "github.com/mfturkcanoglu/go-mux-clean/entities"

type ProductCache interface {
	Set(string, *entities.Product)
	Get(string) *entities.Product
}
