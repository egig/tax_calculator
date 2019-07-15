package main

import "database/sql"

type Model struct {
	DB *sql.DB
}

type Tax struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	TaxCode int8 `json:"tax_code"`
	Price int64 `json:"price"`
}

func (m Model) CreateTax(tax Tax) (sql.Result, error) {
	query := `INSERT INTO tax(name, tax_code, price) VALUES(?,?,?)`
	return m.DB.Exec(query, tax.Name, tax.TaxCode, tax.Price)
}


