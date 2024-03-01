package handler

import "C"
import (
	"errors"
	"github.com/GMaikerYactayo/crud-api-goland/model"
	"github.com/labstack/echo"
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
func (p *productService) create(c echo.Context) error {
	data := model.Product{}
	data.CreateAt = time.Now()
	err := c.Bind(&data)
	if err != nil {
		res := newResponse(Error, "The product does not have a correct structure.", nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	err = p.storage.Create(&data)
	if err != nil {
		res := newResponse(Error, "Problem creating product", nil)
		return c.JSON(http.StatusInternalServerError, res)
	}

	res := newResponse(Message, "Product created successfully", nil)
	return c.JSON(http.StatusCreated, res)
}

// getAll is used for get all products
func (p *productService) getAll(c echo.Context) error {
	data, err := p.storage.GetAll()
	if err != nil {
		res := newResponse(Error, "Error getting products", nil)
		return c.JSON(http.StatusInternalServerError, *res)
	}

	res := newResponse(Message, "Ok", data)
	return c.JSON(http.StatusOK, *res)
}

// getByID is used for get by ID a product
func (p *productService) getByID(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res := newResponse(Error, "The ID must be numerical", nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	data, err := p.storage.GetByID(ID)
	if err != nil {
		res := newResponse(Error, "Problem getting product", nil)
		return c.JSON(http.StatusInternalServerError, res)
	}

	res := newResponse(Message, "Ok", data)
	return c.JSON(http.StatusOK, res)
}

// update is used for update a product
func (p *productService) update(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res := newResponse(Error, "The ID must be numerical", nil)

		return c.JSON(http.StatusBadRequest, *res)
	}

	data := model.Product{}
	data.UpdateAt = time.Now()
	err = c.Bind(&data)
	if err != nil {
		res := newResponse(Error, "The product does not have a correct structure", nil)
		return c.JSON(http.StatusBadRequest, *res)
	}

	err = p.storage.Update(ID, &data)
	if err != nil {
		if errors.Is(err, model.ErrIDProductDoesNotExists) {
			res := newResponse(Error, "The ID product does not exist", nil)
			return c.JSON(http.StatusBadRequest, *res)
		}
		res := newResponse(Error, "Problem updating product", nil)
		return c.JSON(http.StatusInternalServerError, *res)
	}

	res := newResponse(Message, "Product updated successfully", nil)
	return c.JSON(http.StatusOK, *res)
}

// delete is used for delete a product
func (p *productService) delete(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res := newResponse(Error, "The ID must be numerical", nil)
		return c.JSON(http.StatusBadRequest, *res)
	}

	err = p.storage.Delete(ID)
	if err != nil {
		if errors.Is(err, model.ErrIDProductDoesNotExists) {
			res := newResponse(Error, "The ID product does not exist", nil)
			return c.JSON(http.StatusBadRequest, *res)
		}
		res := newResponse(Error, "Problem deleting product", nil)
		return c.JSON(http.StatusInternalServerError, *res)
	}

	res := newResponse(Message, "Product successfully removed", nil)
	return c.JSON(http.StatusOK, *res)
}
