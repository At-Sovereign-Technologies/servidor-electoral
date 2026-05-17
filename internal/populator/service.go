package populator

import (
	"fmt"
	"log"
	"time"

	"github.com/At-Sovereign-Technologies/servidor-electoral/internal/services"
)

type Options struct {
	Populator        PopulatorProvider
	EleccionService  *services.EleccionService
	CandidatoService *services.CandidatoService
	JuradoService    *services.JuradoService
	PuntoService     *services.PuntoService
	TerminalService  *services.TerminalService
	VotanteService   *services.VotanteService
	RefreshInterval  time.Duration
}

type ElectionPopulator struct {
	populator        PopulatorProvider
	eleccionService  *services.EleccionService
	candidatoService *services.CandidatoService
	juradoService    *services.JuradoService
	puntoService     *services.PuntoService
	terminalService  *services.TerminalService
	votanteService   *services.VotanteService
	refreshInterval  time.Duration

	timer *time.Timer
}

func NewElectionPopulator(options Options) *ElectionPopulator {
	return &ElectionPopulator{
		populator:        options.Populator,
		eleccionService:  options.EleccionService,
		candidatoService: options.CandidatoService,
		juradoService:    options.JuradoService,
		puntoService:     options.PuntoService,
		terminalService:  options.TerminalService,
		votanteService:   options.VotanteService,
		refreshInterval:  options.RefreshInterval,
	}
}

func (p *ElectionPopulator) Refresh() error {
	elections, err := p.populator.GetElecciones()
	if err != nil {
		return fmt.Errorf("failed to fetch elections: %w", err)
	}

	for _, election := range elections {
		if err := p.eleccionService.Upsert(election); err != nil {
			log.Printf("failed to upsert election with ID: %d: %v", election.ID, err)
		}

		for _, candidate := range election.Candidatos {
			if err := p.candidatoService.Upsert(&candidate); err != nil {
				log.Printf("failed to upsert candidate with ID: %d: %v", candidate.ID, err)
			}
		}

		for _, punto := range election.Puntos {
			if err := p.puntoService.Upsert(&punto); err != nil {
				log.Printf("failed to upsert punto with ID: %d: %v", punto.ID, err)
			}

			for _, terminal := range punto.Terminales {
				if err := p.terminalService.Upsert(&terminal); err != nil {
					log.Printf("failed to upsert terminal with ID: %d: %v", terminal.ID, err)
				}

				for _, votante := range terminal.Votantes {
					if err := p.votanteService.Upsert(&votante); err != nil {
						log.Printf("failed to upsert votante with ID: %d for terminal ID: %d: %v", votante.ID, terminal.ID, err)
					}
				}
			}

			for _, jurado := range punto.Jurados {
				if err := p.juradoService.Upsert(&jurado); err != nil {
					log.Printf("failed to upsert jurado with ID: %d: %v", jurado.ID, err)
				}
			}
		}
	}

	return nil
}

func (p *ElectionPopulator) Start() error {
	if p.timer != nil {
		p.timer.Stop()
	}

	p.timer = time.AfterFunc(p.refreshInterval, func() {
		if err := p.Refresh(); err != nil {
			log.Printf("Error refreshing elections: %v", err)
		}

		p.Start()
	})

	return nil
}

func (p *ElectionPopulator) Stop() {
	if p.timer != nil {
		p.timer.Stop()
	}
}
