package handler

import (
	"errors"
	"github.com/GMaikerYactayo/crud-api-goland/authorization"
	"github.com/GMaikerYactayo/crud-api-goland/model"
	"github.com/labstack/echo"
	"net/http"
)

// loginService of product
type loginService struct {
	storage UserStorage
}

// newLoginService return a pinter of productService
func newLoginService(s UserStorage) *loginService {
	return &loginService{s}
}

func (l *loginService) login(c echo.Context) error {
	data := model.Login{}
	err := c.Bind(&data)
	if err != nil {
		res := newResponse(Error, "Invalid structure", nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	_, err = l.storage.ValidateCredentials(data.Email, data.Password)
	if err != nil {
		if errors.Is(err, model.ErrUserNotExists) {
			res := newResponse(Error, "User not found or invalid credentials", nil)
			return c.JSON(http.StatusUnauthorized, res)
		}
		res := newResponse(Error, "Error validating credentials", nil)
		return c.JSON(http.StatusInternalServerError, res)
	}

	token, err := authorization.GenerateToken(&data)
	if err != nil {
		res := newResponse(Error, "Cannot generate token", nil)
		return c.JSON(http.StatusInternalServerError, res)
	}

	dataToken := map[string]string{"token": token}
	res := newResponse(Message, "Ok", dataToken)
	return c.JSON(http.StatusOK, *res)
}
