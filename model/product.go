package model

import (
	"time"
)

// Product model
type Product struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	Observations string    `json:"observations"`
	Price        int       `json:"price"`
	CreateAt     time.Time `json:"create_at"`
	UpdateAt     time.Time `json:"update_at"`
}

// Products slice of Product
type Products []*Product
