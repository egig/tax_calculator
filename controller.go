package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"strings"
	"log"
	"github.com/egig/tax_calculator/tax"
)

type Controller struct {
	model Model
}

func (c Controller) BillHandler(w http.ResponseWriter, r *http.Request) {

	type Response struct {
		PriceSubTotal float64 `json:"price_sub_total"`
		TaxSubTotal float64 `json:"tax_sub_total"`
		GrandTotal float64 `json:"grand_total"`
		TaxList []tax.Tax `json:"tax_list"`
	}

	res := Response{}

	taxList, err := c.model.GetTaxObjects()

	for _,t := range taxList {
		res.PriceSubTotal += t.Price
		res.TaxSubTotal += t.TaxAmount
		res.GrandTotal += t.Amount
	}

	res.TaxList = taxList

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		log.Println(err)
		return
	}

	respondWithJSON(w, http.StatusOK, res)
}


func (c Controller) TaxObjectHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var taxObject tax.TaxObject

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

	res, err := c.model.CreateTaxObject(taxObject)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		log.Println(err)
		return
	}

	taxObject.ID, err = res.LastInsertId()

	respondWithJSON(w, http.StatusCreated, taxObject)
}

