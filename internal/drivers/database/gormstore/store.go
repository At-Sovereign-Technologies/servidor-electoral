package gormstore

import (
	"errors"
	"strings"

	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/store"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type Store struct {
	DB *gorm.DB
}

func translateError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return store.ErrNotFound
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return store.ErrConflict
	}

	if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
		return store.ErrConflict
	}

	return err
}

func translateResult(result *gorm.DB) error {
	if err := translateError(result.Error); err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return store.ErrNotFound
	}

	return nil
}
