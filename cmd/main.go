package main

import (
	"github.com/GMaikerYactayo/crud-api-goland/authorization"
	"github.com/GMaikerYactayo/crud-api-goland/handler"
	"github.com/GMaikerYactayo/crud-api-goland/storage"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
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

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	handler.RouteProduct(e, myStorage)

	log.Println("Server initialized on port 8080")
	err = e.Start(":8080")
	if err != nil {
		log.Printf("error en el servidor: %v\n", err)
	}

}
