package main

import (
	"github.com/GMaikerYactayo/crud-api-goland/authorization"
	"github.com/GMaikerYactayo/crud-api-goland/handler"
	"github.com/GMaikerYactayo/crud-api-goland/storage"
	"log"
	"net/http"
)

func main() {
	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("Could not load certificates: %v", err)
	}

	driver := storage.Postgres
	storage.New(driver)
	myStorage, err := storage.DAOProduct(driver)
	if err != nil {
		log.Fatalf("DAOProduct: %v", err)
	}

	mux := http.NewServeMux()

	handler.RouteProduct(mux, myStorage)

	log.Println("Server initialized on port 8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("error en el servidor: %v\n", err)
	}

}
