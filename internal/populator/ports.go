package populator

import (
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
)

type PopulatorProvider interface {
	GetElecciones() ([]*models.Eleccion, error)
}
