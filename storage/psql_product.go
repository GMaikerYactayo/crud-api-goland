package storage

import (
	"database/sql"
	"fmt"
	"github.com/GMaikerYactayo/crud-api-goland/model"
)

type scanner interface {
	Scan(dest ...interface{}) error
}

const (
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
		id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price INT NOT NULL,
		create_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
    	CONSTRAINT products_id_pk PRIMARY KEY (id)
	)`

	psqlCreateProduct = `INSERT INTO products(name, observations, price, create_at)
		VALUES($1, $2 ,$3 ,$4) RETURNING id`

	psqlGetAllProduct = `SELECT id, name, observations, price, create_at, updated_at
		FROM products`

	psqlGetByIdProduct = psqlGetAllProduct + ` WHERE id = $1`

	psqlUpdateProduct = `UPDATE products 
				SET name = $1, observations = $2, price = $3, updated_at = $4 
                WHERE id = $5`

	psqlDeleteProduct = `DELETE FROM products 
				WHERE id = $1`
)

// psqlProduct used for work with postgres - product
type psqlProduct struct {
	db *sql.DB
}

// newPsqlProduct return a new pinter of psqlProduct
func newPsqlProduct(db *sql.DB) *psqlProduct {
	return &psqlProduct{db}
}

// Migrate implement the interface handler.Storage
func (p *psqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	fmt.Println("Product migration successfully executed")
	return nil
}

// Create implement the interface handler.Storage
func (p *psqlProduct) Create(m *model.Product) error {
	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		m.CreateAt,
	).Scan(&m.ID)
	if err != nil {
		return err
	}

	fmt.Println("The product was created successfully")
	return nil
}

// GetAll implement the interface handler.Storage
func (p *psqlProduct) GetAll() (model.Products, error) {
	stmt, err := p.db.Prepare(psqlGetAllProduct)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(model.Products, 0)
	for rows.Next() {
		m, err := scanRowProduct(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ms, nil
}

// GetByID implement the interface handler.Storage
func (p *psqlProduct) GetByID(id int) (*model.Product, error) {
	stmt, err := p.db.Prepare(psqlGetByIdProduct)
	if err != nil {
		return &model.Product{}, err
	}
	defer stmt.Close()

	return scanRowProduct(stmt.QueryRow(id))
}

// Update implement the interface handler.Storage
func (p *psqlProduct) Update(id int, m *model.Product) error {
	_, err := p.GetByID(id)
	if err != nil {
		return fmt.Errorf("ID: %d: %w", id, model.ErrIDProductDoesNotExists)
	}

	stmt, err := p.db.Prepare(psqlUpdateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		timeToNull(m.UpdateAt),
		id,
	)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return fmt.Errorf("don't exist the product id: %d", m.ID)
	}
	fmt.Println("Product was updated successfully")
	return nil
}

// Delete implement the interface handler.Storage
func (p *psqlProduct) Delete(id int) error {
	_, err := p.GetByID(id)
	if err != nil {
		return fmt.Errorf("ID: %d: %w", id, model.ErrIDProductDoesNotExists)
	}

	stmt, err := p.db.Prepare(psqlDeleteProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	fmt.Println("Product was deleted successfully")
	return nil
}

func scanRowProduct(s scanner) (*model.Product, error) {
	m := &model.Product{}
	observationsNull := sql.NullString{}
	updateAtNull := sql.NullTime{}
	err := s.Scan(
		&m.ID,
		&m.Name,
		&observationsNull,
		&m.Price,
		&m.CreateAt,
		&updateAtNull,
	)
	if err != nil {
		return nil, err
	}
	m.Observations = observationsNull.String
	m.UpdateAt = updateAtNull.Time
	return m, nil
}
