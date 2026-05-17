package services

// ============================================================
//  TIPO: Unitaria — VotanteService (Go)
//  Verifica: Lógica de negocio de votantes, validaciones de
//            esquema (documento, nombre) y operaciones sobre Store.
//  Stack: Go standard testing library
// ============================================================

import (
	"errors"
	"testing"

	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/store"
)

// MockVotanteStore es una implementación mock de la interfaz store.VotanteStore
type MockVotanteStore struct {
	onCreate          func(v *models.Votante) error
	onGetByID         func(id uint) (*models.Votante, error)
	onGetByTerminalID func(terminalID uint) ([]models.Votante, error)
	onList            func() ([]models.Votante, error)
	onUpdate          func(v *models.Votante) error
	onDelete          func(id uint) error
}

func (m *MockVotanteStore) Create(v *models.Votante) error {
	if m.onCreate != nil {
		return m.onCreate(v)
	}
	return nil
}

func (m *MockVotanteStore) GetByID(id uint) (*models.Votante, error) {
	if m.onGetByID != nil {
		return m.onGetByID(id)
	}
	return nil, nil
}

func (m *MockVotanteStore) GetByTerminalID(terminalID uint) ([]models.Votante, error) {
	if m.onGetByTerminalID != nil {
		return m.onGetByTerminalID(terminalID)
	}
	return nil, nil
}

func (m *MockVotanteStore) List() ([]models.Votante, error) {
	if m.onList != nil {
		return m.onList()
	}
	return nil, nil
}

func (m *MockVotanteStore) Update(v *models.Votante) error {
	if m.onUpdate != nil {
		return m.onUpdate(v)
	}
	return nil
}

func (m *MockVotanteStore) Delete(id uint) error {
	if m.onDelete != nil {
		return m.onDelete(id)
	}
	return nil
}

// TC-SE-001 | Upsert - Falla validación de nombre corto
func TestVotanteService_Upsert_InvalidNombre(t *testing.T) {
	mockStore := &MockVotanteStore{}
	service := &VotanteService{Store: mockStore}

	votante := &models.Votante{
		TerminalID: 1,
		Nombre:     "Lu", // Mínimo 3 caracteres requeridos
		Documento:  "10203040",
	}

	err := service.Upsert(votante)
	if err == nil {
		t.Fatal("se esperaba error de validación debido a nombre muy corto, pero se obtuvo nil")
	}
}

// TC-SE-002 | Upsert - Falla validación de documento corto
func TestVotanteService_Upsert_InvalidDocumento(t *testing.T) {
	mockStore := &MockVotanteStore{}
	service := &VotanteService{Store: mockStore}

	votante := &models.Votante{
		TerminalID: 1,
		Nombre:     "Luis Perez",
		Documento:  "123", // Mínimo 5 caracteres requeridos
	}

	err := service.Upsert(votante)
	if err == nil {
		t.Fatal("se esperaba error de validación debido a documento corto, pero se obtuvo nil")
	}
}

// TC-SE-003 | Upsert - ID = 0 (Llamada exitosa a Create)
func TestVotanteService_Upsert_CreateNew(t *testing.T) {
	createdCalled := false
	mockStore := &MockVotanteStore{
		onCreate: func(v *models.Votante) error {
			createdCalled = true
			if v.Nombre != "Santiago Mesa" {
				t.Errorf("se esperaba nombre 'Santiago Mesa', se obtuvo '%s'", v.Nombre)
			}
			return nil
		},
	}
	service := &VotanteService{Store: mockStore}

	votante := &models.Votante{
		ID:         0,
		TerminalID: 2,
		Nombre:     "Santiago Mesa",
		Documento:  "100200300",
	}

	err := service.Upsert(votante)
	if err != nil {
		t.Fatalf("error inesperado en Upsert: %v", err)
	}

	if !createdCalled {
		t.Error("se esperaba que se invocara el método Create del store")
	}
}

// TC-SE-004 | Upsert - ID != 0 Existente (Llamada exitosa a Update)
func TestVotanteService_Upsert_UpdateExisting(t *testing.T) {
	getByIDCalled := false
	updateCalled := false

	mockStore := &MockVotanteStore{
		onGetByID: func(id uint) (*models.Votante, error) {
			getByIDCalled = true
			return &models.Votante{
				ID:         id,
				TerminalID: 4,
				Nombre:     "Juan Original",
				Documento:  "555666777",
			}, nil
		},
		onUpdate: func(v *models.Votante) error {
			updateCalled = true
			if v.Nombre != "Juan Modificado" {
				t.Errorf("se esperaba nombre modificado 'Juan Modificado', se obtuvo '%s'", v.Nombre)
			}
			return nil
		},
	}
	service := &VotanteService{Store: mockStore}

	votanteModificado := &models.Votante{
		ID:         99,
		TerminalID: 4,
		Nombre:     "Juan Modificado",
		Documento:  "555666777",
	}

	err := service.Upsert(votanteModificado)
	if err != nil {
		t.Fatalf("error inesperado en Upsert: %v", err)
	}

	if !getByIDCalled {
		t.Error("se esperaba llamada a GetByID")
	}
	if !updateCalled {
		t.Error("se esperaba llamada a Update")
	}
}

// TC-SE-005 | Upsert - ID != 0 Pero no existe en BD (Create Fallback)
func TestVotanteService_Upsert_UpdateNotFound_FallbackToCreate(t *testing.T) {
	getByIDCalled := false
	createCalled := false

	mockStore := &MockVotanteStore{
		onGetByID: func(id uint) (*models.Votante, error) {
			getByIDCalled = true
			return nil, store.ErrNotFound
		},
		onCreate: func(v *models.Votante) error {
			createCalled = true
			return nil
		},
	}
	service := &VotanteService{Store: mockStore}

	votante := &models.Votante{
		ID:         777,
		TerminalID: 10,
		Nombre:     "Votante Perdido",
		Documento:  "999888777",
	}

	err := service.Upsert(votante)
	if err != nil {
		t.Fatalf("error inesperado en Upsert: %v", err)
	}

	if !getByIDCalled {
		t.Error("se esperaba llamada a GetByID")
	}
	if !createCalled {
		t.Error("se esperaba llamada a Create por fallback de ErrNotFound")
	}
}

// TC-SE-006 | Upsert - Error genérico en GetByID propagado
func TestVotanteService_Upsert_GetByIDErrorPropagated(t *testing.T) {
	mockStore := &MockVotanteStore{
		onGetByID: func(id uint) (*models.Votante, error) {
			return nil, errors.New("db connection failure")
		},
	}
	service := &VotanteService{Store: mockStore}

	votante := &models.Votante{
		ID:         88,
		TerminalID: 2,
		Nombre:     "Prueba Error",
		Documento:  "12345678",
	}

	err := service.Upsert(votante)
	if err == nil {
		t.Fatal("se esperaba propagación del error de base de datos, pero se obtuvo nil")
	}
}

// TC-SE-007 | Métodos de lectura y borrado (List, Delete, GetByID)
func TestVotanteService_CRUD_Passthrough(t *testing.T) {
	getByIDCalled := false
	listCalled := false
	deleteCalled := false

	mockStore := &MockVotanteStore{
		onGetByID: func(id uint) (*models.Votante, error) {
			getByIDCalled = true
			return &models.Votante{ID: id}, nil
		},
		onList: func() ([]models.Votante, error) {
			listCalled = true
			return []models.Votante{{ID: 1}, {ID: 2}}, nil
		},
		onDelete: func(id uint) error {
			deleteCalled = true
			return nil
		},
	}
	service := &VotanteService{Store: mockStore}

	// 1. GetByID
	v, err := service.GetByID(42)
	if err != nil || v.ID != 42 || !getByIDCalled {
		t.Errorf("error o falla en GetByID passthrough: %v", err)
	}

	// 2. List
	list, err := service.List()
	if err != nil || len(list) != 2 || !listCalled {
		t.Errorf("error o falla en List passthrough: %v", err)
	}

	// 3. Delete
	err = service.Delete(10)
	if err != nil || !deleteCalled {
		t.Errorf("error o falla en Delete passthrough: %v", err)
	}
}
