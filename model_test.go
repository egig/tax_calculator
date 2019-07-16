package main

import (
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/egig/tax_calculator/tax"
	//"fmt"
	"fmt"
)

func TestModel_CreateTax(t *testing.T) {

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	//INSERT INTO tax(name, tax_code, price) VALUES(?,?,?)
	mock.ExpectExec("INSERT INTO tax").
		WithArgs("Big Mac", 1, float64(1000)).
			WillReturnResult(sqlmock.NewResult(1, 1))

	m := Model{
		DB: db,
	}

	bigMac := tax.TaxObject{
		Name: "Big Mac",
		TaxCode: tax.TaxCodeFood,
		Price: 1000,
	}

	_, err = m.CreateTaxObject(bigMac)

	if err != nil {
		t.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestModel_GetTaxObjects(t *testing.T) {

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	bigMac := tax.TaxObject{
		Name: "Big Mac",
		TaxCode: tax.TaxCodeFood,
		Price: 1000,
	}

	m := Model{
		DB: db,
	}

	_, err = m.CreateTaxObject(bigMac)

	rows := sqlmock.NewRows([]string{"id", "name", "tax_code", "price"}).AddRow(1, "Big Mac", tax.TaxCodeFood, 1000)
	mock.ExpectQuery("SELECT id, name, tax_code, price FROM tax").WillReturnRows(rows)

	taxes, err := m.GetTaxObjects()

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(taxes)

}