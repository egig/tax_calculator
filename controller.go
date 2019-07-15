package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"strings"
	"log"
	"./tax"
)

type Controller struct {
	model Model
}

func (c Controller) ListTaxHandler(w http.ResponseWriter, r *http.Request) {
	taxes, err := c.model.GetTaxes()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		log.Println(err)
		return
	}

	respondWithJSON(w, http.StatusOK, taxes)
}


func (c Controller) CreateTaxHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var taxObject tax.Object

	err := decoder.Decode(&taxObject)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		log.Println(err)
		return
	}

	// Simple Validation
	var fieldErr []string
	if taxObject.Name == "" {
		fieldErr = append(fieldErr, "name")
	}

	if taxObject.TaxCode == 0 {
		fieldErr = append(fieldErr, "tax_code")
	}

	if taxObject.Price == 0 {
		fieldErr = append(fieldErr, "price")
	}

	if len(fieldErr) > 0 {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%s required", strings.Join(fieldErr, ", ")))
		return
	}

	if !InSlice(int(taxObject.TaxCode), []int{tax.TaxCodeFood, tax.TaxCodeTobacco, tax.TaxCodeEnt}) {
		respondWithError(w, http.StatusBadRequest, "tax_code must be either 1,2 or 3")
		return
	}

	res, err := c.model.CreateTax(taxObject)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		log.Println(err)
		return
	}

	taxObject.ID, err = res.LastInsertId()

	respondWithJSON(w, http.StatusCreated, taxObject)
}

