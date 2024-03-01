package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/GMaikerYactayo/crud-api-goland/model"
)

const (
	psqlValidateCredentials = `SELECT id, email, password 
								FROM users 
								WHERE email = $1 AND password = $2`
)

// psqlUser used for work with postgres - user
type psqlUser struct {
	db *sql.DB
}

// newPsqlUser return a new pinter of psqlUser
func newPsqlUser(db *sql.DB) *psqlUser {
	return &psqlUser{db}
}

func (p *psqlUser) ValidateCredentials(email, password string) (*model.User, error) {
	stmt, err := p.db.Prepare(psqlValidateCredentials)
	if err != nil {
		return nil, fmt.Errorf("prepare statement: %w", err)
	}
	defer stmt.Close()

	user, err := scanRowUser(stmt.QueryRow(email, password))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// User not found or invalid credentials
			return nil, fmt.Errorf("user not found: %w", model.ErrUserNotExists)
		}
		// other error
		return nil, fmt.Errorf("scan user: %w", err)
	}

	return user, nil
}

func scanRowUser(s scanner) (*model.User, error) {
	m := &model.User{}
	err := s.Scan(
		&m.ID,
		&m.Email,
		&m.Password,
	)
	if err != nil {
		return nil, err
	}
	return m, nil
}
