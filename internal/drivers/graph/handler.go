package graph

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/services"
	"github.com/labstack/echo/v4"
)

type GraphQLHandler struct {
	schema *graphql.ExecutableSchema
}

func NewGraphQLHandler(
	eleccionService *services.EleccionService,
	candidatoService *services.CandidatoService,
	puntoService *services.PuntoService,
	juradoService *services.JuradoService,
	terminalService *services.TerminalService,
	votanteService *services.VotanteService,
	nodoService *services.NodoService,
) *GraphQLHandler {
	resolvers := &Resolver{
		EleccionService:  eleccionService,
		CandidatoService: candidatoService,
		PuntoService:     puntoService,
		JuradoService:    juradoService,
		TerminalService:  terminalService,
		VotanteService:   votanteService,
		NodoService:      nodoService,
	}

	schema := NewExecutableSchema(Config{
		Resolvers: resolvers,
	})

	return &GraphQLHandler{
		schema: &schema,
	}
}

// Register mounts the GraphQL handlers to the Echo router
func (h *GraphQLHandler) Register(e *echo.Echo) {
	srv := handler.NewDefaultServer(*h.schema)

	// GraphQL endpoint
	e.POST("/query", h.GraphQL(srv))

	// GraphQL Playground (for development)
	e.GET("/", h.Playground())
}

// GraphQL returns the GraphQL query handler
func (h *GraphQLHandler) GraphQL(srv *handler.Server) echo.HandlerFunc {
	return func(c echo.Context) error {
		srv.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}

// Playground returns the GraphQL playground handler
func (h *GraphQLHandler) Playground() echo.HandlerFunc {
	return func(c echo.Context) error {
		playground.Handler("GraphQL Playground", "/query").ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
