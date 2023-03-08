package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IfIGetAnErrorIfIDIsBlank(t *testing.T) {
	order := Order{}
	assert.Error(t, order.Validate(), "invalid id")
}

func Test_IfIGetAnErrorIfPriceIsBlank(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.Validate(), "invalid price")
}

func Test_IfIGetAnErrorIfTaxIsBlank(t *testing.T) {
	order := Order{ID: "123", Price: 10.0}
	assert.Error(t, order.Validate(), "invalid tax")
}

func Test_WithAllValidParams(t *testing.T) {
	order := Order{ID: "123", Price: 10.0, Tax: 2.0}
	assert.NoError(t, order.Validate())
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
	order.CalculateFinalPrice()
	assert.Equal(t, 20.0, order.FinalPrice)
}
