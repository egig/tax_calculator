package tax

const (
	TaxCodeFood = 1
	TaxCodeTobacco = 2
	TaxCodeEnt = 3
)

type Object struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	TaxCode int8 `json:"tax_code"`
	Price int64 `json:"price"`
}

type tax struct {
	Object
	TypeName string `json:"type_name"`
	Refundable bool `json:"refundable"`
	TaxAmount float32  `json:"tax"`
	Amount float32 `json:"amount"`
}

func newFoodTax(t Object) tax {

	taxAmount := float32(10/100 * t.Price)

	ft := tax{
		Object: t,
		TypeName: "Food & Beverages",
		TaxAmount: taxAmount,
		Amount: float32(float32(t.Price) - taxAmount),
	}

	return ft
}

func newTobaccoTax(t Object) tax {
	taxAmount := float32(10 + (2/100 * t.Price))

	tt := tax{
		Object: t,
		TypeName: "Food & Beverages",
		TaxAmount: taxAmount,
		Amount: float32(float32(t.Price) - taxAmount),
	}

	return tt
}

func newEntTax(t Object) tax {

	taxAmount := float32(1/100 * (t.Price - 100))

	et := tax{
		Object: t,
		TypeName: "Food & Beverages",
		TaxAmount: taxAmount,
		Amount: float32(float32(t.Price) - taxAmount),
	}

	return et
}

func NewTax(o Object) tax {

	var t tax
	if o.TaxCode == TaxCodeFood {
		t = newFoodTax(o)
	} else if o.TaxCode == TaxCodeTobacco {
		t = newTobaccoTax(o)
	} else if o.TaxCode == TaxCodeEnt {
		t = newEntTax(o)
	}

	t.Amount = float32(t.Price) + t.TaxAmount

	return t
}