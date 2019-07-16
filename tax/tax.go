package tax

import (
	"log"
)

const (
	TaxCodeFood = 1
	TaxCodeTobacco = 2
	TaxCodeEnt = 3
)

type Object struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	TaxCode int8 `json:"tax_code"`
	Price float64 `json:"price"`
}

type Tax struct {
	Object
	TypeName string `json:"type_name"`
	Refundable bool `json:"refundable"`
	TaxAmount float64  `json:"tax"`
	Amount float64 `json:"amount"`
}

func NewFoodTax(t Object) Tax {

	taxAmount := float64((10 * t.Price)/100)

	ft := Tax{
		Object: t,
		TypeName: "Food & Beverages",
		TaxAmount: taxAmount,
		Refundable: true,
	}

	return ft
}

func NewTobaccoTax(t Object) Tax {
	taxAmount := float64(10 + ((2 * t.Price)/100) )

	tt := Tax{
		Object: t,
		TypeName: "Tobacco",
		TaxAmount: taxAmount,
		Refundable: false,
	}

	return tt
}

func NewEntTax(t Object) Tax {

	var taxAmount float64
	if t.Price < 100 {
		taxAmount = 0
	} else {
		taxAmount = float64(float64(t.Price - 100)/float64(100))
	}

	et := Tax{
		Object: t,
		TypeName: "Entertainment",
		TaxAmount: taxAmount,
		Refundable: false,
	}

	return et
}

func NewTax(o Object) Tax {

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