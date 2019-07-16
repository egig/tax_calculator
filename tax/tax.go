package tax

import (
	"log"
)

// Tax code constants
const (
	TaxCodeFood    = 1
	TaxCodeTobacco = 2
	TaxCodeEnt     = 3
)

// TaxObject tax.TaxObject
type TaxObject struct {
	ID      int64   `json:"id"`
	Name    string  `json:"name"`
	TaxCode int8    `json:"tax_code"`
	Price   float64 `json:"price"`
}

// Tax tax
type Tax struct {
	TaxObject
	TypeName   string  `json:"type_name"`
	Refundable bool    `json:"refundable"`
	TaxAmount  float64 `json:"tax"`
	Amount     float64 `json:"amount"`
}


// NewFoodTax create new Tax with type food
func NewFoodTax(t TaxObject) Tax {

	taxAmount := float64((10 * t.Price) / 100)

	ft := Tax{
		TaxObject:  t,
		TypeName:   "Food & Beverages",
		TaxAmount:  taxAmount,
		Refundable: true,
	}

	return ft
}

// NewTobaccoTax create new Tax with type tobacco
func NewTobaccoTax(t TaxObject) Tax {
	taxAmount := float64(10 + ((2 * t.Price) / 100))

	tt := Tax{
		TaxObject:  t,
		TypeName:   "Tobacco",
		TaxAmount:  taxAmount,
		Refundable: false,
	}

	return tt
}

// NewEntTax create new Tax with type entertainment
func NewEntTax(t TaxObject) Tax {

	var taxAmount float64
	if t.Price < 100 {
		taxAmount = 0
	} else {
		taxAmount = float64(float64(t.Price-100) / float64(100))
	}

	et := Tax{
		TaxObject:  t,
		TypeName:   "Entertainment",
		TaxAmount:  taxAmount,
		Refundable: false,
	}

	return et
}

// NewTax create new Tax based on TaxObject
func NewTax(o TaxObject) Tax {

	var t Tax
	if o.TaxCode == TaxCodeFood {
		t = NewFoodTax(o)
	} else if o.TaxCode == TaxCodeTobacco {
		t = NewTobaccoTax(o)
	} else if o.TaxCode == TaxCodeEnt {
		t = NewEntTax(o)
	} else {
		log.Fatal("Unknown Tax Code: ", o.TaxCode)
	}

	t.Amount = float64(t.Price) + t.TaxAmount

	return t
}
