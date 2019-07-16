package main

import (
	"database/sql"
	"github.com/egig/tax_calculator/tax"
)

type Model struct {
	DB *sql.DB
}

func (m Model) CreateTaxObject(t tax.TaxObject) (sql.Result, error) {
	query := `INSERT INTO tax(name, tax_code, price) VALUES(?,?,?)`
	return m.DB.Exec(query, t.Name, t.TaxCode, t.Price)
}

func (m Model) GetTaxObjects() ([]tax.Tax, error) {

	rows, err := m.DB.Query("SELECT id, name, tax_code, price FROM tax")

	if err == sql.ErrNoRows {
		return make([]tax.Tax, 0), nil
	}

	if err != nil {
		return nil, err
	}

	var taxes []tax.Tax

	defer rows.Close()

	for rows.Next() {
		var to tax.TaxObject
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


