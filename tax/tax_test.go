package tax

import (
	"gotest.tools/assert"
	"testing"
)

func TestNewFoodTax(t *testing.T) {
	o := TaxObject{
		0,
		"Big Mac",
		1,
		1000,
	}

	ft := NewFoodTax(o)

	assert.Equal(t, ft.TaxAmount, float64(100))
}

func TestNewTobaccoTax(t *testing.T) {

	o := TaxObject{
		0,
		"Lucky Strike",
		2,
		1000,
	}

	ft := NewTobaccoTax(o)

	assert.Equal(t, ft.TaxAmount, float64(30))
}

func TestNewEntTax(t *testing.T) {
	o := TaxObject{
		0,
		"Lucky Strike",
		3,
		150,
	}

	ft := NewEntTax(o)
	assert.Equal(t, ft.TaxAmount, 0.5)
}

func TestNewTax(t *testing.T) {
	o := TaxObject{
		0,
		"Big Mac",
		1,
		1000,
	}

	nt := NewTax(o)

	assert.Equal(t, nt.Amount, float64(1100))
}
