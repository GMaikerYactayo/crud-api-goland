package handler

import "net/http"

// RouteProduct configures product-related routes in a given ServeMux,
// using the provided storage to handle product operations.
func RouteProduct(mux *http.ServeMux, storage Storage) {
	h := newService(storage)
	mux.HandleFunc("/v1/products/create", h.create)
	mux.HandleFunc("/v1/products/get-all", h.getAll)
	mux.HandleFunc("/v1/products/get-by-id", h.getByID)
	mux.HandleFunc("/v1/products/update", h.update)
	mux.HandleFunc("/v1/products/delete", h.delete)
}
