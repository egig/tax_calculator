package main

import (
	"database/sql"
	"./tax"
)

type Model struct {
	DB *sql.DB
}

func (m Model) CreateTax(t tax.Object) (sql.Result, error) {
	query := `INSERT INTO tax(name, tax_code, price) VALUES(?,?,?)`
	return m.DB.Exec(query, t.Name, t.TaxCode, t.Price)
}

func (m Model) GetTaxes() ([]interface{}, error) {

	rows, err := m.DB.Query("SELECT id, name, tax_code, price FROM tax")
	if err != nil {
		return nil, err
	}

	var taxes []interface{}

	defer rows.Close()

	for rows.Next() {
		var to tax.Object
		err := rows.Scan(
			&to.ID,
			&to.Name,
			&to.TaxCode,
			&to.Price)

		if err != nil {
			return nil, err
		}

		t := tax.NewTax(to)

		taxes = append(taxes, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return taxes,nil
}


