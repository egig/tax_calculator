package main

import (
	"bytes"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"gotest.tools/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestController_BillHandler(t *testing.T) {

	req := httptest.NewRequest("GET", "/bill", nil)
	w := httptest.NewRecorder()

	db, mock, _ := sqlmock.New()

	m := Model{
		DB: db,
	}

	c := Controller{
		m,
	}

	mock.ExpectQuery("SELECT id, name, tax_code, price FROM tax").
		WillReturnError(sql.ErrNoRows)

	c.BillHandler(w, req)

	resp := w.Result()
	assert.Equal(t, resp.StatusCode, http.StatusOK)
}

func TestModel_CreateTaxObject(t *testing.T) {

	payload := []byte(`{"name":"Big Mac","tax_code": 1,"price": 1000}`)
	req := httptest.NewRequest("POST", "/tax_objects", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	defer req.Body.Close()

	w := httptest.NewRecorder()

	db, mock, _ := sqlmock.New()

	m := Model{
		DB: db,
	}

	c := Controller{
		m,
	}

	mock.ExpectExec("INSERT INTO tax").
		WithArgs("Big Mac", 1, float64(1000)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	c.TaxObjectHandler(w, req)

	resp := w.Result()

	assert.Equal(t, resp.StatusCode, http.StatusCreated)
}
