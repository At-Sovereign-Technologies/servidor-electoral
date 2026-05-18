package services

import (
	"fmt"
	"time"
)

// ElectionWindowService valida las restricciones temporales de una elección
// RN-02: Duración exacta de 48 horas
// RN-03: Ventana de voto por correo de 5 días antes del inicio
type ElectionWindowService struct{}

// NewElectionWindowService crea una nueva instancia del servicio de ventanas electorales
func NewElectionWindowService() *ElectionWindowService {
	return &ElectionWindowService{}
}

// ValidateElectionDuration verifica que la elección dure exactamente 48 horas
// RN-02: La duración debe ser exactamente 48 horas (172800 segundos)
func (s *ElectionWindowService) ValidateElectionDuration(startTimeUnix, endTimeUnix int64) error {
	const expectedDurationSeconds int64 = 48 * 60 * 60 // 172800 segundos

	if startTimeUnix <= 0 || endTimeUnix <= 0 {
		return fmt.Errorf("timestamps inválidos: start=%d, end=%d", startTimeUnix, endTimeUnix)
	}

	if endTimeUnix <= startTimeUnix {
		return fmt.Errorf("fecha fin debe ser posterior a fecha inicio")
	}

	actualDuration := endTimeUnix - startTimeUnix

	if actualDuration != expectedDurationSeconds {
		return fmt.Errorf("duración debe ser exactamente 48 horas (%d segundos), obtuvo %d segundos",
			expectedDurationSeconds, actualDuration)
	}

	return nil
}

// IsWithinElectionWindow verifica si un timestamp está dentro de la ventana electoral
// Incluye los límites (inicio y fin exactos son válidos)
func (s *ElectionWindowService) IsWithinElectionWindow(voteTimestamp, startTimeUnix, endTimeUnix int64) bool {
	return voteTimestamp >= startTimeUnix && voteTimestamp <= endTimeUnix
}

// IsWithinMailVoteWindow verifica si un timestamp está dentro de la ventana de voto por correo
// RN-03: Ventana abierta exactamente 5 días antes del inicio, cierra al inicio de la jornada
func (s *ElectionWindowService) IsWithinMailVoteWindow(voteTimestamp, startTimeUnix, endTimeUnix int64) bool {
	const mailVoteDaysBeforeStart int64 = 5 * 24 * 60 * 60 // 5 días en segundos

	mailVoteStartUnix := startTimeUnix - mailVoteDaysBeforeStart

	// Ventana: [startTime - 5 días, startTime) — incluye inicio de ventana, EXCLUYE inicio de jornada
	return voteTimestamp >= mailVoteStartUnix && voteTimestamp < startTimeUnix
}

// GetElectionDurationHours retorna la duración en horas (para verificación manual)
func (s *ElectionWindowService) GetElectionDurationHours(startTimeUnix, endTimeUnix int64) float64 {
	return float64(endTimeUnix-startTimeUnix) / 3600.0
}

// GetMailVoteWindowStart retorna el timestamp de inicio de ventana de correo (5 días antes)
func (s *ElectionWindowService) GetMailVoteWindowStart(startTimeUnix int64) int64 {
	const mailVoteDaysBeforeStart int64 = 5 * 24 * 60 * 60
	return startTimeUnix - mailVoteDaysBeforeStart
}

// ValidateElectionDates es un wrapper que valida tanto duración como timestamps válidos
func (s *ElectionWindowService) ValidateElectionDates(startTimeUnix, endTimeUnix int64) error {
	if startTimeUnix <= 0 || endTimeUnix <= 0 {
		return fmt.Errorf("timestamps inválidos")
	}

	if endTimeUnix <= startTimeUnix {
		return fmt.Errorf("fecha fin debe ser posterior a fecha inicio")
	}

	return s.ValidateElectionDuration(startTimeUnix, endTimeUnix)
}

// NowUnix retorna el timestamp actual (para testing se puede mockear)
// Esta función es útil para pruebas de tiempo real
func (s *ElectionWindowService) NowUnix() int64 {
	return time.Now().Unix()
}
