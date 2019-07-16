package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"database/sql"
	"log"
	"net/http"
	"os"
)



func main() {

	dbDriver := os.Getenv("DB_DRIVER")
	dbDSN := os.Getenv("DB_DSN")
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

	port := os.Getenv("APP_PORT")

	log.Printf("App run in port %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}