package store

import (
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
)

type JuradoStore interface {
	Create(j *models.Jurado) error
	GetByID(id uint) (*models.Jurado, error)
	GetByPuntoID(puntoID uint) ([]models.Jurado, error)
	List() ([]models.Jurado, error)
	Update(j *models.Jurado) error
	Delete(id uint) error
}
