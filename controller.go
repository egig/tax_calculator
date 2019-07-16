package main

import (
	"encoding/json"
	"fmt"
	"github.com/egig/tax_calculator/tax"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Controller is object holding the http handler
type Controller struct {
	model Model
}

// BillHandler http handler to get bills
func (c Controller) BillHandler(w http.ResponseWriter, r *http.Request) {

	type Response struct {
		PriceSubTotal float64   `json:"price_sub_total"`
		TaxSubTotal   float64   `json:"tax_sub_total"`
		GrandTotal    float64   `json:"grand_total"`
		TaxList       []tax.Tax `json:"tax_list"`
	}

	res := Response{
		TaxList: make([]tax.Tax, 0),
	}

	taxList, err := c.model.GetTaxObjects()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		log.Println(err)
		return
	}

	if taxList != nil {
		res.TaxList = taxList

		for _, t := range taxList {
			res.PriceSubTotal += t.Price
			res.TaxSubTotal += t.TaxAmount
			res.GrandTotal += t.Amount
		}
	}

	respondWithJSON(w, http.StatusOK, res)
}

// TaxObjectHandler http handler to create Tax Object
func (c Controller) TaxObjectHandler(w http.ResponseWriter, r *http.Request) {

	b, _ := ioutil.ReadAll(r.Body)

	var taxObject tax.TaxObject

	err := json.Unmarshal(b, &taxObject)

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

	res, err := c.model.CreateTaxObject(taxObject)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		log.Println(err)
		return
	}

	taxObject.ID, err = res.LastInsertId()

	respondWithJSON(w, http.StatusCreated, taxObject)
}
