package store

import (
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
)

type NodoStore interface {
	Create(n *models.Nodo) error
	GetByID(id uint) (*models.Nodo, error)
	List() ([]models.Nodo, error)
	Update(n *models.Nodo) error
	Delete(id uint) error
}
