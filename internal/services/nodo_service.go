package services

import (
	"errors"
	"fmt"

	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/store"
)

type NodoService struct {
	Store store.NodoStore
}

func (s *NodoService) Upsert(nodo *models.Nodo) error {
	if err := nodo.Validate(); err != nil {
		return fmt.Errorf("invalid nodo data: %w", err)
	}

	if nodo.ID == 0 {
		return s.Store.Create(nodo)
	}

	existing, err := s.Store.GetByID(nodo.ID)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			return s.Store.Create(nodo)
		}
		return fmt.Errorf("failed to check existing nodo: %w", err)
	}

	existing.Activo = nodo.Activo
	existing.Secreto = nodo.Secreto

	if err := existing.Validate(); err != nil {
		return fmt.Errorf("invalid nodo data: %w", err)
	}

	return s.Store.Update(existing)
}

func (s *NodoService) GetByID(id uint) (*models.Nodo, error) {
	return s.Store.GetByID(id)
}

func (s *NodoService) List() ([]models.Nodo, error) {
	return s.Store.List()
}

func (s *NodoService) Delete(id uint) error {
	return s.Store.Delete(id)
}
