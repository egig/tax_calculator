package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"database/sql"
	"log"
	"net/http"
	"encoding/json"
	"fmt"
	"strings"
)

const (
	TaxCodeFood = 1
	TaxCodeTobacco = 2
	TaxCodeEnt = 3
)

func main() {

	dbDriver := "mysql"
	dbDSN := "root@/tax_calculator"
	db, err := sql.Open(dbDriver, dbDSN)
	if err != nil {
		log.Fatal(err)
	}

	m := Model{
		DB: db,
	}

	router := mux.NewRouter()
	router.HandleFunc("/tax", createTaxHandler(m))

	const port = ":8080";

	log.Printf("App run in port %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}


func createTaxHandler(m Model) http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {

		decoder := json.NewDecoder(r.Body)

		var tax Tax

		err := decoder.Decode(&tax)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			log.Println(err)
			return
		}

		// Simple Validation
		var fieldErr []string
		if tax.Name == "" {
			fieldErr = append(fieldErr, "name")
		}

		if tax.TaxCode == 0 {
			fieldErr = append(fieldErr, "tax_code")
		}

		if tax.Price == 0 {
			fieldErr = append(fieldErr, "price")
		}

		if len(fieldErr) > 0 {
			respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%s required", strings.Join(fieldErr, ", ")))
			return
		}

		res, err := m.CreateTax(tax)

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			log.Println(err)
			return
		}

		tax.ID, err = res.LastInsertId()

		respondWithJSON(w, http.StatusCreated, tax)
	}
}