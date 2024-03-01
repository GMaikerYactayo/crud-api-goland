package handler

import (
	"github.com/GMaikerYactayo/crud-api-goland/middleware"
	"net/http"
)

// RouteProduct configures product-related routes in a given ServeMux,
// using the provided storage to handle product operations.
func RouteProduct(mux *http.ServeMux, storage Storage) {
	h := newProductService(storage)
	mux.HandleFunc("/v1/products/create", h.create)
	mux.HandleFunc("/v1/products/get-all", middleware.Log(middleware.Authentication(h.getAll)))
	mux.HandleFunc("/v1/products/get-by-id", middleware.Log(h.getByID))
	mux.HandleFunc("/v1/products/update", h.update)
	mux.HandleFunc("/v1/products/delete", h.delete)
}

func RuteUser(mux *http.ServeMux, userStorage UserStorage) {
	h := newLoginService(userStorage)
	mux.HandleFunc("/v1/login", h.login)
}
