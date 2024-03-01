package handler

import (
	"encoding/json"
	"errors"
	"github.com/GMaikerYactayo/crud-api-goland/authorization"
	"github.com/GMaikerYactayo/crud-api-goland/model"
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

func (l *loginService) login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		res := newResponse(Error, "Disallowed method", nil)
		responseJSON(w, http.StatusBadRequest, *res)
		return
	}

	data := model.Login{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		res := newResponse(Error, "Invalid structure", nil)
		responseJSON(w, http.StatusBadRequest, *res)
		return
	}

	_, err = l.storage.ValidateCredentials(data.Email, data.Password)
	if err != nil {
		if errors.Is(err, model.ErrUserNotExists) {
			res := newResponse(Error, "User not found or invalid credentials", nil)
			responseJSON(w, http.StatusUnauthorized, *res)
			return
		}
		// other error
		res := newResponse(Error, "Error validating credentials", nil)
		responseJSON(w, http.StatusInternalServerError, *res)
		return
	}

	token, err := authorization.GenerateToken(&data)
	if err != nil {
		res := newResponse(Error, "Cannot generate token", nil)
		responseJSON(w, http.StatusInternalServerError, *res)
		return
	}

	dataToken := map[string]string{"token": token}
	res := newResponse(Message, "Ok", dataToken)
	responseJSON(w, http.StatusOK, *res)
}
