package graph

// THIS CODE WILL BE UPDATED WITH SCHEMA CHANGES. PREVIOUS IMPLEMENTATION FOR SCHEMA CHANGES WILL BE KEPT IN THE COMMENT SECTION. IMPLEMENTATION FOR UNCHANGED SCHEMA WILL BE KEPT.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/models"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/services"
)

type Resolver struct {
	EleccionService  *services.EleccionService
	CandidatoService *services.CandidatoService
	PuntoService     *services.PuntoService
	JuradoService    *services.JuradoService
	TerminalService  *services.TerminalService
	VotanteService   *services.VotanteService
	NodoService      *services.NodoService
}

// Elecciones is the resolver for the elecciones field.
func (r *queryResolver) Elecciones(ctx context.Context) ([]*Eleccion, error) {
	elecciones, err := r.Resolver.EleccionService.List()
	if err != nil {
		return nil, fmt.Errorf("failed to get elecciones: %w", err)
	}

	result := make([]*Eleccion, len(elecciones))
	for i, e := range elecciones {
		result[i] = modelToGraphQL(&e)
	}
	return result, nil
}

// Eleccion is the resolver for the eleccion field.
func (r *queryResolver) Eleccion(ctx context.Context, id string) (*Eleccion, error) {
	eleccionID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid election id: %w", err)
	}

	eleccion, err := r.Resolver.EleccionService.GetByID(uint(eleccionID))
	if err != nil {
		return nil, fmt.Errorf("failed to get eleccion: %w", err)
	}

	return modelToGraphQL(eleccion), nil
}

// CandidatosByEleccion is the resolver for the candidatosByEleccion field.
func (r *queryResolver) CandidatosByEleccion(ctx context.Context, eleccionID string) ([]*Candidato, error) {
	id, err := strconv.ParseUint(eleccionID, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid election id: %w", err)
	}

	candidatos, err := r.Resolver.CandidatoService.GetByEleccionID(uint(id))
	if err != nil {
		return nil, fmt.Errorf("failed to get candidatos: %w", err)
	}

	result := make([]*Candidato, len(candidatos))
	for i, c := range candidatos {
		result[i] = candidatoToGraphQL(&c)
	}
	return result, nil
}

// Candidato is the resolver for the candidato field.
func (r *queryResolver) Candidato(ctx context.Context, id string) (*Candidato, error) {
	candidatoID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid candidato id: %w", err)
	}

	candidato, err := r.Resolver.CandidatoService.GetByID(uint(candidatoID))
	if err != nil {
		return nil, fmt.Errorf("failed to get candidato: %w", err)
	}

	return candidatoToGraphQL(candidato), nil
}

// PuntosByEleccion is the resolver for the puntosByEleccion field.
func (r *queryResolver) PuntosByEleccion(ctx context.Context, eleccionID string) ([]*Punto, error) {
	id, err := strconv.ParseUint(eleccionID, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid election id: %w", err)
	}

	puntos, err := r.Resolver.PuntoService.GetByEleccionID(uint(id))
	if err != nil {
		return nil, fmt.Errorf("failed to get puntos: %w", err)
	}

	result := make([]*Punto, len(puntos))
	for i, p := range puntos {
		result[i] = puntoToGraphQL(&p)
	}
	return result, nil
}

// Punto is the resolver for the punto field.
func (r *queryResolver) Punto(ctx context.Context, id string) (*Punto, error) {
	puntoID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid punto id: %w", err)
	}

	punto, err := r.Resolver.PuntoService.GetByID(uint(puntoID))
	if err != nil {
		return nil, fmt.Errorf("failed to get punto: %w", err)
	}

	return puntoToGraphQL(punto), nil
}

// TerminalesByPunto is the resolver for the terminalesByPunto field.
func (r *queryResolver) TerminalesByPunto(ctx context.Context, puntoID string) ([]*Terminal, error) {
	id, err := strconv.ParseUint(puntoID, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid punto id: %w", err)
	}

	terminales, err := r.Resolver.TerminalService.GetByPuntoID(uint(id))
	if err != nil {
		return nil, fmt.Errorf("failed to get terminales: %w", err)
	}

	result := make([]*Terminal, len(terminales))
	for i, t := range terminales {
		result[i] = terminalToGraphQL(&t)
	}
	return result, nil
}

// Terminal is the resolver for the terminal field.
func (r *queryResolver) Terminal(ctx context.Context, id string) (*Terminal, error) {
	terminalID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid terminal id: %w", err)
	}

	terminal, err := r.Resolver.TerminalService.GetByID(uint(terminalID))
	if err != nil {
		return nil, fmt.Errorf("failed to get terminal: %w", err)
	}

	return terminalToGraphQL(terminal), nil
}

// VotantesByTerminal is the resolver for the votantesByTerminal field.
func (r *queryResolver) VotantesByTerminal(ctx context.Context, terminalID string) ([]*Votante, error) {
	id, err := strconv.ParseUint(terminalID, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid terminal id: %w", err)
	}

	votantes, err := r.Resolver.VotanteService.GetByTerminalID(uint(id))
	if err != nil {
		return nil, fmt.Errorf("failed to get votantes: %w", err)
	}

	result := make([]*Votante, len(votantes))
	for i, v := range votantes {
		result[i] = votanteToGraphQL(&v)
	}
	return result, nil
}

// Votante is the resolver for the votante field.
func (r *queryResolver) Votante(ctx context.Context, id string) (*Votante, error) {
	votanteID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid votante id: %w", err)
	}

	votante, err := r.Resolver.VotanteService.GetByID(uint(votanteID))
	if err != nil {
		return nil, fmt.Errorf("failed to get votante: %w", err)
	}

	return votanteToGraphQL(votante), nil
}

// JuradosByPunto is the resolver for the juradosByPunto field.
func (r *queryResolver) JuradosByPunto(ctx context.Context, puntoID string) ([]*Jurado, error) {
	id, err := strconv.ParseUint(puntoID, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid punto id: %w", err)
	}

	jurados, err := r.Resolver.JuradoService.GetByPuntoID(uint(id))
	if err != nil {
		return nil, fmt.Errorf("failed to get jurados: %w", err)
	}

	result := make([]*Jurado, len(jurados))
	for i, j := range jurados {
		result[i] = juradoToGraphQL(&j)
	}
	return result, nil
}

// Jurado is the resolver for the jurado field.
func (r *queryResolver) Jurado(ctx context.Context, id string) (*Jurado, error) {
	juradoID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid jurado id: %w", err)
	}

	jurado, err := r.Resolver.JuradoService.GetByID(uint(juradoID))
	if err != nil {
		return nil, fmt.Errorf("failed to get jurado: %w", err)
	}

	return juradoToGraphQL(jurado), nil
}

// NodosByEleccion is the resolver for the nodosByEleccion field.
func (r *queryResolver) NodosByEleccion(ctx context.Context, eleccionID string) ([]*Nodo, error) {
	id, err := strconv.ParseUint(eleccionID, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid election id: %w", err)
	}

	nodos, err := r.Resolver.NodoService.GetByEleccionID(uint(id))
	if err != nil {
		return nil, fmt.Errorf("failed to get nodos: %w", err)
	}

	result := make([]*Nodo, len(nodos))
	for i, n := range nodos {
		result[i] = nodoToGraphQL(&n)
	}
	return result, nil
}

// Nodo is the resolver for the nodo field.
func (r *queryResolver) Nodo(ctx context.Context, id string) (*Nodo, error) {
	nodoID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid nodo id: %w", err)
	}

	nodo, err := r.Resolver.NodoService.GetByID(uint(nodoID))
	if err != nil {
		return nil, fmt.Errorf("failed to get nodo: %w", err)
	}

	return nodoToGraphQL(nodo), nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// ======== MODEL CONVERSION HELPERS ========

func modelToGraphQL(e *models.Eleccion) *Eleccion {
	if e == nil {
		return nil
	}
	return &Eleccion{
		ID:           fmt.Sprintf("%d", e.ID),
		Nombre:       e.Nombre,
		TipoEleccion: TipoEleccion(e.TipoEleccion),
		FechaInicio:  int(e.FechaInicio),
		FechaFin:     int(e.FechaFin),
		Candidatos:   candidatosToGraphQL(e.Candidatos),
		Puntos:       puntosToGraphQL(e.Puntos),
		Nodos:        nodosToGraphQL(e.Nodos),
		CreatedAt:    e.CreatedAt.String(),
		UpdatedAt:    e.UpdatedAt.String(),
	}
}

func candidatoToGraphQL(c *models.Candidato) *Candidato {
	if c == nil {
		return nil
	}
	return &Candidato{
		ID:         fmt.Sprintf("%d", c.ID),
		Nombre:     c.Nombre,
		Documento:  c.Documento,
		Partido:    c.Partido,
		FotoURL:    c.FotoURL,
		EleccionID: fmt.Sprintf("%d", c.EleccionID),
		CreatedAt:  c.CreatedAt.String(),
		UpdatedAt:  c.UpdatedAt.String(),
	}
}

func candidatosToGraphQL(cs []models.Candidato) []*Candidato {
	result := make([]*Candidato, len(cs))
	for i, c := range cs {
		result[i] = candidatoToGraphQL(&c)
	}
	return result
}

func puntoToGraphQL(p *models.Punto) *Punto {
	if p == nil {
		return nil
	}
	return &Punto{
		ID:         fmt.Sprintf("%d", p.ID),
		Nombre:     p.Nombre,
		Latitud:    p.Latitud,
		Longitud:   p.Longitud,
		Activo:     p.Activo,
		EleccionID: fmt.Sprintf("%d", p.EleccionID),
		Jurados:    juradosToGraphQL(p.Jurados),
		Terminales: terminalsToGraphQL(p.Terminales),
		CreatedAt:  p.CreatedAt.String(),
		UpdatedAt:  p.UpdatedAt.String(),
	}
}

func puntosToGraphQL(ps []models.Punto) []*Punto {
	result := make([]*Punto, len(ps))
	for i, p := range ps {
		result[i] = puntoToGraphQL(&p)
	}
	return result
}

func terminalToGraphQL(t *models.Terminal) *Terminal {
	if t == nil {
		return nil
	}
	return &Terminal{
		ID:           fmt.Sprintf("%d", t.ID),
		Activo:       t.Activo,
		PuntoID:      fmt.Sprintf("%d", t.PuntoID),
		ClavePublica: t.ClavePublica,
		Votantes:     votantesToGraphQL(t.Votantes),
		CreatedAt:    t.CreatedAt.String(),
		UpdatedAt:    t.UpdatedAt.String(),
	}
}

func terminalsToGraphQL(ts []models.Terminal) []*Terminal {
	result := make([]*Terminal, len(ts))
	for i, t := range ts {
		result[i] = terminalToGraphQL(&t)
	}
	return result
}

func votanteToGraphQL(v *models.Votante) *Votante {
	if v == nil {
		return nil
	}
	return &Votante{
		ID:         fmt.Sprintf("%d", v.ID),
		Nombre:     v.Nombre,
		Documento:  v.Documento,
		TerminalID: fmt.Sprintf("%d", v.TerminalID),
		CreatedAt:  v.CreatedAt.String(),
		UpdatedAt:  v.UpdatedAt.String(),
	}
}

func votantesToGraphQL(vs []models.Votante) []*Votante {
	result := make([]*Votante, len(vs))
	for i, v := range vs {
		result[i] = votanteToGraphQL(&v)
	}
	return result
}

func juradoToGraphQL(j *models.Jurado) *Jurado {
	if j == nil {
		return nil
	}
	return &Jurado{
		ID:        fmt.Sprintf("%d", j.ID),
		Nombre:    j.Nombre,
		Documento: j.Documento,
		Usuario:   j.Usuario,
		PuntoID:   fmt.Sprintf("%d", j.PuntoID),
		CreatedAt: j.CreatedAt.String(),
		UpdatedAt: j.UpdatedAt.String(),
	}
}

func juradosToGraphQL(js []models.Jurado) []*Jurado {
	result := make([]*Jurado, len(js))
	for i, j := range js {
		result[i] = juradoToGraphQL(&j)
	}
	return result
}

func nodoToGraphQL(n *models.Nodo) *Nodo {
	if n == nil {
		return nil
	}
	return &Nodo{
		ID:         fmt.Sprintf("%d", n.ID),
		Activo:     n.Activo,
		EleccionID: fmt.Sprintf("%d", n.EleccionID),
		CreatedAt:  "N/A",
		UpdatedAt:  "N/A",
	}
}

func nodosToGraphQL(ns []models.Nodo) []*Nodo {
	result := make([]*Nodo, len(ns))
	for i, n := range ns {
		result[i] = nodoToGraphQL(&n)
	}
	return result
}
