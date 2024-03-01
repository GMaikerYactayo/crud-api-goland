package handler

import (
	"encoding/json"
	"errors"
	"github.com/GMaikerYactayo/crud-api-goland/model"
	"net/http"
	"strconv"
	"time"
)

// productService of product
type productService struct {
	storage Storage
}

// newProductService return a pinter of productService
func newProductService(s Storage) *productService {
	return &productService{s}
}

// create is used for create a product
func (p *productService) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		res := newResponse(Error, "Disallowed method", nil)
		responseJSON(w, http.StatusBadRequest, *res)
		return
	}

	data := model.Product{}
	data.CreateAt = time.Now()
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		res := newResponse(Error, "The product does not have a correct structure.", nil)
		responseJSON(w, http.StatusBadRequest, *res)
		return
	}

	err = p.storage.Create(&data)
	if err != nil {
		res := newResponse(Error, "Problem creating product", nil)
		responseJSON(w, http.StatusInternalServerError, *res)
		return
	}

	res := newResponse(Message, "Product created successfully", nil)
	responseJSON(w, http.StatusCreated, *res)
}

// getAll is used for get all products
func (p *productService) getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		res := newResponse(Error, "Disallowed method", nil)
		responseJSON(w, http.StatusBadRequest, *res)
		return
	}

	data, err := p.storage.GetAll()
	if err != nil {
		res := newResponse(Error, "Error getting products", nil)
		responseJSON(w, http.StatusInternalServerError, *res)
		return
	}

	res := newResponse(Message, "Ok", data)
	responseJSON(w, http.StatusOK, *res)
}

// getByID is used for get by ID a product
func (p *productService) getByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		res := newResponse(Error, "Disallowed method", nil)
		responseJSON(w, http.StatusBadRequest, *res)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		res := newResponse(Error, "The ID must be numerical", nil)
		responseJSON(w, http.StatusBadRequest, *res)
		return
	}

	data, err := p.storage.GetByID(ID)
	if err != nil {
		res := newResponse(Error, "Problem getting product", nil)
		responseJSON(w, http.StatusInternalServerError, *res)
		return
	}

	res := newResponse(Message, "Ok", data)
	responseJSON(w, http.StatusOK, *res)
}

// update is used for update a product
func (p *productService) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		res := newResponse(Error, "Disallowed method", nil)
		responseJSON(w, http.StatusBadRequest, *res)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		res := newResponse(Error, "The ID must be numerical", nil)
		responseJSON(w, http.StatusBadRequest, *res)
		return
	}

	data := model.Product{}
	data.UpdateAt = time.Now()
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		res := newResponse(Error, "The product does not have a correct structure", nil)
		responseJSON(w, http.StatusBadRequest, *res)
		return
	}

	err = p.storage.Update(ID, &data)
	if errors.Is(err, model.ErrIDProductDoesNotExists) {
		res := newResponse(Error, "The ID product does not exist", nil)
		responseJSON(w, http.StatusBadRequest, *res)
	}
	if err != nil {
		res := newResponse(Error, "Problem updating product", nil)
		responseJSON(w, http.StatusInternalServerError, *res)
		return
	}

	res := newResponse(Message, "Product updated successfully", nil)
	responseJSON(w, http.StatusOK, *res)
}

// delete is used for delete a product
func (p *productService) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		res := newResponse(Error, "Disallowed method", nil)
		responseJSON(w, http.StatusBadRequest, *res)
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		res := newResponse(Error, "The ID must be numerical", nil)
		responseJSON(w, http.StatusBadRequest, *res)
		return
	}

	err = p.storage.Delete(ID)
	if errors.Is(err, model.ErrIDProductDoesNotExists) {
		res := newResponse(Error, "The ID product does not exist", nil)
		responseJSON(w, http.StatusBadRequest, *res)
	}
	if err != nil {
		res := newResponse(Error, "Problem deleting product", nil)
		responseJSON(w, http.StatusInternalServerError, *res)
		return
	}

	res := newResponse(Message, "Product successfully removed", nil)
	responseJSON(w, http.StatusOK, *res)
}
