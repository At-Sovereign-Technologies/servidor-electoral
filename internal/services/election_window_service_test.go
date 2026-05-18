package services

import (
	"testing"
)

// ============================================================
//  TIPO: Unitaria — ElectionWindowService (Go)
//  Valida: RN-02 (duración exacta 48h), RN-03 (ventana correo 5 días)
//  Criterios: Límites inclusivos, exactitud de timing
//  Stack: Go standard testing library
// ============================================================

// TestElectionWindowService_ValidateDuration_Exactly48Hours
// TC-SE-001: Acepta elección con duración exactamente 48 horas
func TestElectionWindowService_ValidateDuration_Exactly48Hours(t *testing.T) {
	service := NewElectionWindowService()

	// 2026-06-01 08:00:00 UTC
	startTime := int64(1777008000)
	// 2026-06-03 08:00:00 UTC (exactamente 48h después)
	endTime := startTime + (48 * 60 * 60)

	err := service.ValidateElectionDuration(startTime, endTime)
	if err != nil {
		t.Fatalf("TC-SE-001 FALLÓ: se esperaba error nil para duración exacta de 48h, obtuvo: %v", err)
	}
}

// TestElectionWindowService_ValidateDuration_47Hours
// TC-SE-002: Rechaza elección con duración 47 horas
func TestElectionWindowService_ValidateDuration_47Hours(t *testing.T) {
	service := NewElectionWindowService()

	startTime := int64(1777008000)
	// 47 horas después
	endTime := startTime + (47 * 60 * 60)

	err := service.ValidateElectionDuration(startTime, endTime)
	if err == nil {
		t.Fatal("TC-SE-002 FALLÓ: se esperaba error para duración 47h, obtuvo nil")
	}

	if err.Error() != "duración debe ser exactamente 48 horas (172800 segundos), obtuvo 169200 segundos" {
		t.Fatalf("TC-SE-002: mensaje de error inesperado: %v", err)
	}
}

// TestElectionWindowService_ValidateDuration_49Hours
// TC-SE-003: Rechaza elección con duración 49 horas
func TestElectionWindowService_ValidateDuration_49Hours(t *testing.T) {
	service := NewElectionWindowService()

	startTime := int64(1777008000)
	// 49 horas después
	endTime := startTime + (49 * 60 * 60)

	err := service.ValidateElectionDuration(startTime, endTime)
	if err == nil {
		t.Fatal("TC-SE-003 FALLÓ: se esperaba error para duración 49h, obtuvo nil")
	}
}

// TestElectionWindowService_ValidateDuration_24Hours
// TC-SE-004: Rechaza elección con duración 24 horas (1 día)
func TestElectionWindowService_ValidateDuration_24Hours(t *testing.T) {
	service := NewElectionWindowService()

	startTime := int64(1777008000)
	// 24 horas después
	endTime := startTime + (24 * 60 * 60)

	err := service.ValidateElectionDuration(startTime, endTime)
	if err == nil {
		t.Fatal("TC-SE-004 FALLÓ: se esperaba error para duración 24h, obtuvo nil")
	}
}

// TestElectionWindowService_ValidateDuration_72Hours
// TC-SE-005: Rechaza elección con duración 72 horas (3 días)
func TestElectionWindowService_ValidateDuration_72Hours(t *testing.T) {
	service := NewElectionWindowService()

	startTime := int64(1777008000)
	// 72 horas después
	endTime := startTime + (72 * 60 * 60)

	err := service.ValidateElectionDuration(startTime, endTime)
	if err == nil {
		t.Fatal("TC-SE-005 FALLÓ: se esperaba error para duración 72h, obtuvo nil")
	}
}

// TestElectionWindowService_IsWithinElectionWindow_Inside
// TC-SE-010: Voto dentro de la ventana electoral
func TestElectionWindowService_IsWithinElectionWindow_Inside(t *testing.T) {
	service := NewElectionWindowService()

	startTime := int64(1777008000)
	endTime := startTime + (48 * 60 * 60)
	voteTime := startTime + (24 * 60 * 60) // Mitad de la ventana

	if !service.IsWithinElectionWindow(voteTime, startTime, endTime) {
		t.Fatal("TC-SE-010 FALLÓ: se esperaba que el voto esté dentro de la ventana")
	}
}

// TestElectionWindowService_IsWithinElectionWindow_Before
// TC-SE-011: Voto 1 minuto antes del inicio
func TestElectionWindowService_IsWithinElectionWindow_Before(t *testing.T) {
	service := NewElectionWindowService()

	startTime := int64(1777008000)
	endTime := startTime + (48 * 60 * 60)
	voteTime := startTime - 60 // 1 minuto antes

	if service.IsWithinElectionWindow(voteTime, startTime, endTime) {
		t.Fatal("TC-SE-011 FALLÓ: se esperaba que el voto esté fuera de la ventana (antes)")
	}
}

// TestElectionWindowService_IsWithinElectionWindow_After
// TC-SE-012: Voto 1 minuto después del fin
func TestElectionWindowService_IsWithinElectionWindow_After(t *testing.T) {
	service := NewElectionWindowService()

	startTime := int64(1777008000)
	endTime := startTime + (48 * 60 * 60)
	voteTime := endTime + 60 // 1 minuto después

	if service.IsWithinElectionWindow(voteTime, startTime, endTime) {
		t.Fatal("TC-SE-012 FALLÓ: se esperaba que el voto esté fuera de la ventana (después)")
	}
}

// TestElectionWindowService_IsWithinElectionWindow_ExactStart
// TC-SE-013: Voto exactamente en el timestamp inicial (inclusivo)
func TestElectionWindowService_IsWithinElectionWindow_ExactStart(t *testing.T) {
	service := NewElectionWindowService()

	startTime := int64(1777008000)
	endTime := startTime + (48 * 60 * 60)

	if !service.IsWithinElectionWindow(startTime, startTime, endTime) {
		t.Fatal("TC-SE-013 FALLÓ: el timestamp inicial debe ser válido (inclusivo)")
	}
}

// TestElectionWindowService_IsWithinElectionWindow_ExactEnd
// TC-SE-014: Voto exactamente en el timestamp final (inclusivo)
func TestElectionWindowService_IsWithinElectionWindow_ExactEnd(t *testing.T) {
	service := NewElectionWindowService()

	startTime := int64(1777008000)
	endTime := startTime + (48 * 60 * 60)

	if !service.IsWithinElectionWindow(endTime, startTime, endTime) {
		t.Fatal("TC-SE-014 FALLÓ: el timestamp final debe ser válido (inclusivo)")
	}
}

// ============================================================
// RN-03: Ventana de Voto por Correo (5 días antes)
// ============================================================

// TestElectionWindowService_IsWithinMailVoteWindow_Exactly5Days
// TC-SE-301: Acepta voto correo exactamente 5 días antes del inicio (primer segundo de ventana)
func TestElectionWindowService_IsWithinMailVoteWindow_Exactly5Days(t *testing.T) {
	service := NewElectionWindowService()

	startTime := int64(1777008000) // 2026-06-01 08:00:00
	endTime := startTime + (48 * 60 * 60)
	// Exactamente 5 días antes (primer segundo de ventana de correo)
	voteTime := startTime - (5 * 24 * 60 * 60)

	if !service.IsWithinMailVoteWindow(voteTime, startTime, endTime) {
		t.Fatal("TC-SE-301 FALLÓ: voto correo a 5 días exactos debe ser válido")
	}
}

// TestElectionWindowService_IsWithinMailVoteWindow_5DaysAnd1Minute
// TC-SE-302: Rechaza voto correo 5 días y 1 minuto antes (fuera de ventana, muy antiguo)
func TestElectionWindowService_IsWithinMailVoteWindow_5DaysAnd1Minute(t *testing.T) {
	service := NewElectionWindowService()

	startTime := int64(1777008000)
	endTime := startTime + (48 * 60 * 60)
	// 5 días + 1 minuto antes (fuera de ventana, demasiado temprano)
	voteTime := startTime - (5*24*60*60 + 60)

	if service.IsWithinMailVoteWindow(voteTime, startTime, endTime) {
		t.Fatal("TC-SE-302 FALLÓ: voto correo más de 5 días antes debe ser rechazado")
	}
}

// TestElectionWindowService_IsWithinMailVoteWindow_1DayBefore
// TC-SE-303: Acepta voto correo el día antes de la elección (aún dentro de ventana)
func TestElectionWindowService_IsWithinMailVoteWindow_1DayBefore(t *testing.T) {
	service := NewElectionWindowService()

	startTime := int64(1777008000)
	endTime := startTime + (48 * 60 * 60)
	// 1 día antes (aún dentro de la ventana correo de 5 días)
	voteTime := startTime - (24 * 60 * 60)

	if !service.IsWithinMailVoteWindow(voteTime, startTime, endTime) {
		t.Fatal("TC-SE-303 FALLÓ: voto correo 1 día antes debe ser válido")
	}
}

// TestElectionWindowService_IsWithinMailVoteWindow_DuringElection
// TC-SE-304: Rechaza voto correo exactamente al inicio de la jornada (ventana cerrada)
func TestElectionWindowService_IsWithinMailVoteWindow_DuringElection(t *testing.T) {
	service := NewElectionWindowService()

	startTime := int64(1777008000)
	endTime := startTime + (48 * 60 * 60)
	// Exactamente al inicio (ventana correo cierra aquí, comienza electoral)
	voteTime := startTime

	if service.IsWithinMailVoteWindow(voteTime, startTime, endTime) {
		t.Fatal("TC-SE-304 FALLÓ: voto correo al inicio de jornada debe ser rechazado (ventana cierra)")
	}
}

// TestElectionWindowService_IsWithinMailVoteWindow_AfterElection
// TC-SE-305: Rechaza voto correo después del fin de la elección
func TestElectionWindowService_IsWithinMailVoteWindow_AfterElection(t *testing.T) {
	service := NewElectionWindowService()

	startTime := int64(1777008000)
	endTime := startTime + (48 * 60 * 60)
	// 1 hora después del fin
	voteTime := endTime + (60 * 60)

	if service.IsWithinMailVoteWindow(voteTime, startTime, endTime) {
		t.Fatal("TC-SE-305 FALLÓ: voto correo después de la elección debe ser rechazado")
	}
}

// TestElectionWindowService_GetMailVoteWindowStart
// TC-SE-310: Calcula correctamente el inicio de ventana correo (5 días antes)
func TestElectionWindowService_GetMailVoteWindowStart(t *testing.T) {
	service := NewElectionWindowService()

	startTime := int64(1777008000)
	expected := startTime - (5 * 24 * 60 * 60)

	result := service.GetMailVoteWindowStart(startTime)
	if result != expected {
		t.Fatalf("TC-SE-310 FALLÓ: esperado %d, obtuvo %d", expected, result)
	}
}

// TestElectionWindowService_GetElectionDurationHours
// TC-SE-311: Calcula duración correcta en horas
func TestElectionWindowService_GetElectionDurationHours(t *testing.T) {
	service := NewElectionWindowService()

	startTime := int64(1777008000)
	endTime := startTime + (48 * 60 * 60)

	result := service.GetElectionDurationHours(startTime, endTime)
	if result != 48.0 {
		t.Fatalf("TC-SE-311 FALLÓ: esperado 48.0 horas, obtuvo %f", result)
	}
}

// TestElectionWindowService_ValidateDates_InvalidTimestamps
// TC-SE-320: Rechaza timestamps inválidos (≤ 0)
func TestElectionWindowService_ValidateDates_InvalidTimestamps(t *testing.T) {
	service := NewElectionWindowService()

	err := service.ValidateElectionDates(0, 100)
	if err == nil {
		t.Fatal("TC-SE-320 FALLÓ: debería rechazar timestamp inicio = 0")
	}

	err = service.ValidateElectionDates(100, 0)
	if err == nil {
		t.Fatal("TC-SE-320 FALLÓ: debería rechazar timestamp fin = 0")
	}
}

// TestElectionWindowService_ValidateDates_EndBeforeStart
// TC-SE-321: Rechaza fin anterior al inicio
func TestElectionWindowService_ValidateDates_EndBeforeStart(t *testing.T) {
	service := NewElectionWindowService()

	startTime := int64(2000)
	endTime := int64(1000)

	err := service.ValidateElectionDates(startTime, endTime)
	if err == nil {
		t.Fatal("TC-SE-321 FALLÓ: debería rechazar fin < inicio")
	}
}

// BenchmarkIsWithinElectionWindow
// Benchmark para validar rendimiento (sub-milisegundo esperado)
func BenchmarkIsWithinElectionWindow(b *testing.B) {
	service := NewElectionWindowService()
	startTime := int64(1777008000)
	endTime := startTime + (48 * 60 * 60)
	voteTime := startTime + (24 * 60 * 60)

	for i := 0; i < b.N; i++ {
		_ = service.IsWithinElectionWindow(voteTime, startTime, endTime)
	}
}

// BenchmarkValidateElectionDuration
// Benchmark para duración
func BenchmarkValidateElectionDuration(b *testing.B) {
	service := NewElectionWindowService()
	startTime := int64(1777008000)
	endTime := startTime + (48 * 60 * 60)

	for i := 0; i < b.N; i++ {
		_ = service.ValidateElectionDuration(startTime, endTime)
	}
}

// TableDrivenTest para validar múltiples valores de duración
// TC-SE-330 a TC-SE-339: Paramétrico
func TestElectionWindowService_MultiDurations(t *testing.T) {
	tests := []struct {
		name          string
		durationHours int64
		shouldPass    bool
	}{
		{"47 horas", 47, false},
		{"48 horas", 48, true},
		{"49 horas", 49, false},
		{"24 horas", 24, false},
		{"72 horas", 72, false},
		{"96 horas", 96, false},
	}

	service := NewElectionWindowService()
	startTime := int64(1777008000)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			endTime := startTime + (tt.durationHours * 60 * 60)
			err := service.ValidateElectionDuration(startTime, endTime)

			if tt.shouldPass && err != nil {
				t.Errorf("esperado éxito para %s, obtuvo error: %v", tt.name, err)
			}
			if !tt.shouldPass && err == nil {
				t.Errorf("esperado error para %s, obtuvo nil", tt.name)
			}
		})
	}
}

// TableDrivenTest para ventanas de correo
// TC-SE-340 a TC-SE-349: Paramétrico
func TestElectionWindowService_MailVoteWindows(t *testing.T) {
	tests := []struct {
		name       string
		daysBefore float64
		shouldPass bool
	}{
		{"5 días exacto", 5.0, true},
		{"4.99 días", 4.99, true},
		{"4 días", 4.0, true},
		{"1 día", 1.0, true},
		{"0 días (inicio jornada)", 0.0, false}, // Ventana cierra al inicio
		{"5.01 días", 5.01, false},
		{"6 días", 6.0, false},
	}

	service := NewElectionWindowService()
	startTime := int64(1777008000)
	endTime := startTime + (48 * 60 * 60)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			voteTime := startTime - int64(tt.daysBefore*24*60*60)
			result := service.IsWithinMailVoteWindow(voteTime, startTime, endTime)

			if tt.shouldPass && !result {
				t.Errorf("esperado válido para %s, obtuvo inválido", tt.name)
			}
			if !tt.shouldPass && result {
				t.Errorf("esperado inválido para %s, obtuvo válido", tt.name)
			}
		})
	}
}
