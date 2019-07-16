package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"database/sql"
	"log"
	"net/http"
)



func main() {

	dbDriver := "mysql"
	dbDSN := "root:password@tcp(db)/tax_calculator"
	db, err := sql.Open(dbDriver, dbDSN)
	if err != nil {
		log.Fatal(err)
	}

	m := Model{
		DB: db,
	}

	c := Controller{
		m,
	}

	router := mux.NewRouter()
	router.HandleFunc("/tax", c.ListTaxHandler).Methods(http.MethodGet)
	router.HandleFunc("/tax", c.CreateTaxHandler).Methods(http.MethodPost)

	const port = ":8080"

	log.Printf("App run in port %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}