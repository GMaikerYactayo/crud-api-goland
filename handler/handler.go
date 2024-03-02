package handler

import (
	"github.com/GMaikerYactayo/crud-api-goland/model"
)

// ProductStorage interface that must implement a db storage
type ProductStorage interface {
	Migrate() error
	Create(*model.Product) error
	GetAll() (model.Products, error)
	GetByID(int) (*model.Product, error)
	Update(int, *model.Product) error
	Delete(int) error
}

type UserStorage interface {
	ValidateCredentials(email, password string) (*model.User, error)
}
