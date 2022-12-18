package services

import (
	"strings"
	"testing"

	"github.com/mfturkcanoglu/go-mux-clean/entities"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestValidateEmptyProduct(t *testing.T) {
	var product entities.Product = entities.Product{}
	err := Validate(product)
	var expected = validator.New().Struct(product)
	assert.NotNil(t, err)
	assert.Equal(t, expected.Error(), err.Error())
}

func TestValidateShortCharName(t *testing.T) {
	var product entities.Product = entities.Product{
		ID:          0,
		Name:        "..",
		Price:       15.0,
		Description: strings.Repeat(".", 25),
	}
	err := Validate(product)
	var expected = validator.New().Struct(product)
	assert.NotNil(t, err)
	assert.Equal(t, expected.Error(), err.Error())
}

func TestValidateSmallAmountPrice(t *testing.T) {
	var product entities.Product = entities.Product{
		ID:          0,
		Name:        strings.Repeat(".", 15),
		Price:       3.0,
		Description: strings.Repeat(".", 25),
	}
	err := Validate(product)
	var expected = validator.New().Struct(product)
	assert.NotNil(t, err)
	assert.Equal(t, expected.Error(), err.Error())
}
