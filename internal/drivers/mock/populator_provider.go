package mock

import (
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/populator"
)

var _ populator.PopulatorProvider = (*MockPopulatorProvider)(nil)

type MockPopulatorProvider struct{}

func (m *MockPopulatorProvider) GetElecciones() ([]*models.Eleccion, error) {
	return []*models.Eleccion{
		{
			ID:           1,
			Nombre:       "Elección Presidencial 2026",
			TipoEleccion: models.EleccionPresidencial,
			FechaInicio:  1700000000,
			FechaFin:     1702592000,
			Candidatos: []models.Candidato{
				{
					ID:         1,
					EleccionID: 1,
					Nombre:     "Iván Cepeda",
					Documento:  "123456789",
					Partido:    "Pacto Histórico",
					FotoURL:    "https://upload.wikimedia.org/wikipedia/commons/5/58/Perfil_Iv%C3%A1n_Cepeda_%28cropped%29.jpg",
				},
				{
					ID:         2,
					EleccionID: 1,
					Nombre:     "Paloma Valencia",
					Documento:  "987654321",
					Partido:    "Centro Democrático",
					FotoURL:    "https://upload.wikimedia.org/wikipedia/commons/f/f9/Discurso_Paloma_Valencia.jpg",
				},
				{
					ID:         3,
					EleccionID: 1,
					Nombre:     "Abelardo de la Espriella",
					Documento:  "555555555",
					Partido:    "Defensores de la patria",
					FotoURL:    "https://upload.wikimedia.org/wikipedia/commons/7/7b/A_de_la_Espriella.jpg",
				},
				{
					ID:         4,
					EleccionID: 1,
					Nombre:     "Sergio Fajardo",
					Documento:  "111111111",
					Partido:    "Dignidad & Compromiso",
					FotoURL:    "https://upload.wikimedia.org/wikipedia/commons/a/aa/Sergio_Fajardo_2025.jpg",
				},
			},
			Puntos: []models.Punto{
				{
					ID:       1,
					Nombre:   "Corferias",
					Latitud:  4.6306,
					Longitud: -74.0927,
					Jurados: []models.Jurado{
						{
							ID:        1,
							PuntoID:   1,
							Nombre:    "Miguel Ángel Pérez",
							Documento: "1234567890",
							Usuario:   "jurado1",
							Hash:      "$argon2i$v=19$m=16,t=2,p=1$UWtiaVJteEVOcFBralU4TQ$+zVMq/f0RBy5JtgFy+WGUw",
						},
						{
							ID:        2,
							PuntoID:   1,
							Nombre:    "Laura Gómez",
							Documento: "0987654321",
							Usuario:   "jurado2",
							Hash:      "$argon2i$v=19$m=16,t=2,p=1$V3l5c1JtZk9XQjVhT3l6TQ$+zVMq/f0RBy5JtgFy+WGUw",
						},
						{
							ID:        3,
							PuntoID:   1,
							Nombre:    "Carlos Rodríguez",
							Documento: "5555555555",
							Usuario:   "jurado3",
							Hash:      "$argon2i$v=19$m=16,t=2,p=1$V3l5c1JtZk9XQjVhT3l6TQ$+zVMq/f0RBy5JtgFy+WGUw",
						},
					},
					Terminales: []models.Terminal{
						{
							ID:      1,
							PuntoID: 1,
							Votantes: []models.Votante{
								{
									ID:         1,
									TerminalID: 1,
									Nombre:     "Juan Pérez",
									Documento:  "1234567890",
								},
								{
									ID:         2,
									TerminalID: 1,
									Nombre:     "María Gómez",
									Documento:  "0987654321",
								},
								{
									ID:         3,
									TerminalID: 1,
									Nombre:     "Luis Rodríguez",
									Documento:  "5555555555",
								},
								{
									ID:         4,
									TerminalID: 1,
									Nombre:     "Ana Martínez",
									Documento:  "1111111111",
								},
							},
						},
						{
							ID:      2,
							PuntoID: 1,
							Votantes: []models.Votante{
								{
									ID:         5,
									TerminalID: 2,
									Nombre:     "Pedro Sánchez",
									Documento:  "2222222222",
								},
								{
									ID:         6,
									TerminalID: 2,
									Nombre:     "Sofía López",
									Documento:  "3333333333",
								},
								{
									ID:         7,
									TerminalID: 2,
									Nombre:     "Diego Fernández",
									Documento:  "4444444444",
								},
							},
						},
					},
				},
				{
					ID:       2,
					Nombre:   "Plaza de Bolívar",
					Latitud:  4.5981,
					Longitud: -74.0758,
					Jurados: []models.Jurado{
						{
							ID:        4,
							PuntoID:   2,
							Nombre:    "Sofía Martínez",
							Documento: "2222222222",
							Usuario:   "jurado4",
							Hash:      "$argon2i$v=19$m=16,t=2,p=1$V3l5c1JtZk9XQjVhT3l6TQ$+zVMq/f0RBy5JtgFy+WGUw",
						},
						{
							ID:        5,
							PuntoID:   2,
							Nombre:    "Andrés Gómez",
							Documento: "3333333333",
							Usuario:   "jurado5",
							Hash:      "$argon2i$v=19$m=16,t=2,p=1$V3l5c1JtZk9XQjVhT3l6TQ$+zVMq/f0RBy5JtgFy+WGUw",
						},
					},
					Terminales: []models.Terminal{
						{
							ID:      3,
							PuntoID: 2,
							Votantes: []models.Votante{
								{
									ID:         8,
									TerminalID: 3,
									Nombre:     "Laura Sánchez",
									Documento:  "5555555555",
								},
								{
									ID:         9,
									TerminalID: 3,
									Nombre:     "Carlos Martínez",
									Documento:  "6666666666",
								},
							},
						},
						{
							ID:      4,
							PuntoID: 2,
							Votantes: []models.Votante{
								{
									ID:         10,
									TerminalID: 4,
									Nombre:     "Ana Rodríguez",
									Documento:  "7777777777",
								},
								{
									ID:         11,
									TerminalID: 4,
									Nombre:     "Luis Gómez",
									Documento:  "8888888888",
								},
							},
						},
					},
				},
			},
		},
	}, nil
}
