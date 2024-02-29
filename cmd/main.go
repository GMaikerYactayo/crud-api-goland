package main

import (
	"github.com/GMaikerYactayo/crud-api-goland/handler"
	"github.com/GMaikerYactayo/crud-api-goland/storage"
	"log"
	"net/http"
)

func main() {
	driver := storage.Postgres
	storage.New(driver)
	myStorage, err := storage.DAOProduct(driver)
	if err != nil {
		log.Fatalf("DAOProduct: %v", err)
	}

	mux := http.NewServeMux()

	handler.RouteProduct(mux, myStorage)

	log.Println("Servidor se inicializo en el puerto 8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("error en el servidor: %v\n", err)
	}

}
