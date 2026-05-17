package main

import (
	"log"
	"time"

	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/drivers/database"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/drivers/database/gormstore"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/drivers/mock"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/populator"
	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/services"
)

func main() {
	if err := database.Init(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	log.Println("Migrating database...")

	if err := database.Migrate(); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database migrated successfully")

	// Store creation
	store := gormstore.Store{DB: database.GetDB()}
	eleccionStore := &gormstore.GormEleccionStore{Store: store}
	candidatoStore := &gormstore.GormCandidatoStore{Store: store}
	puntoStore := &gormstore.GormPuntoStore{Store: store}
	juradoStore := &gormstore.GormJuradoStore{Store: store}
	terminalStore := &gormstore.GormTerminalStore{Store: store}
	votanteStore := &gormstore.GormVotanteStore{Store: store}
	// nodoStore := &gormstore.GormNodoStore{Store: store}

	// Service creation
	eleccionService := &services.EleccionService{Store: eleccionStore}
	candidatoService := &services.CandidatoService{Store: candidatoStore}
	puntoService := &services.PuntoService{Store: puntoStore}
	juradoService := &services.JuradoService{Store: juradoStore}
	terminalService := &services.TerminalService{Store: terminalStore}
	votanteService := &services.VotanteService{Store: votanteStore}
	// nodoService := &services.NodoService{Store: nodoStore}

	// Populator creation
	populator := populator.NewElectionPopulator(populator.Options{
		Populator:        &mock.MockPopulatorProvider{},
		EleccionService:  eleccionService,
		CandidatoService: candidatoService,
		JuradoService:    juradoService,
		PuntoService:     puntoService,
		TerminalService:  terminalService,
		VotanteService:   votanteService,
		RefreshInterval:  10 * time.Second,
	})

	log.Println("Starting election populator...")

	populator.Start()
	defer populator.Stop()

	// Router creation
	// router, err := web.NewRouter(eleccionService, candidatoService, puntoService, juradoService)
	// if err != nil {
	// 	log.Fatalf("Failed to create router: %v", err)
	// }

	// Start HTTP server
	// log.Println("Starting server on :8080")
	// if err := router.Start(":8080"); err != nil {
	// 	log.Fatalf("Failed to start server: %v", err)
	// }

	select {}
}
