package service

import (
	"context"
	"time"

	"Frank2006xmotorq/db/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

type TelemetryService struct {
	q *sqlc.Queries
}

func NewTelemetryService(q *sqlc.Queries) *TelemetryService {
	return &TelemetryService{q: q}
}

type CreateTelemetryRequest struct {
	VehicleID         int64     `json:"vehicle_id"`
	Timestamp         time.Time `json:"timestamp"`
	SpeedMph          float64   `json:"speed_mph"`
	FuelLevelPercent  float64   `json:"fuel_level_percent"`
	EngineState       string    `json:"engine_state"`
	OdometerMiles     float64   `json:"odometer_miles"`
	Status            string    `json:"status"`
}

func (s *TelemetryService) Create(ctx context.Context, req CreateTelemetryRequest) (sqlc.TelemetryHistory, error) {
	status := req.Status
	if status == "" {
		status = "pending"
	}

	return s.q.CreateTelemetryHistory(ctx, sqlc.CreateTelemetryHistoryParams{
		VehicleID:         req.VehicleID,
		Timestamp:         pgtype.Timestamp{Time: req.Timestamp, Valid: true},
		SpeedMph:          floatToNumeric(req.SpeedMph),
		FuelLevelPercent:  floatToNumeric(req.FuelLevelPercent),
		EngineState:       req.EngineState,
		OdometerMiles:     floatToNumeric(req.OdometerMiles),
		Status:            status,
	})
}

func (s *TelemetryService) Get(ctx context.Context, id int64) (sqlc.TelemetryHistory, error) {
	return s.q.GetTelemetryHistory(ctx, id)
}

func (s *TelemetryService) ListByVehicle(ctx context.Context, vehicleID int64) ([]sqlc.TelemetryHistory, error) {
	return s.q.ListTelemetryHistoryByVehicle(ctx, vehicleID)
}

func (s *TelemetryService) ListLatestByVehicle(ctx context.Context, vehicleID int64, limit int32) ([]sqlc.TelemetryHistory, error) {
	return s.q.ListLatestTelemetryByVehicle(ctx, sqlc.ListLatestTelemetryByVehicleParams{
		VehicleID: vehicleID,
		Limit:     limit,
	})
}

func (s *TelemetryService) UpdateStatus(ctx context.Context, id int64, status string) (sqlc.TelemetryHistory, error) {
	return s.q.UpdateTelemetryHistoryStatus(ctx, sqlc.UpdateTelemetryHistoryStatusParams{
		ID:     id,
		Status: status,
	})
}