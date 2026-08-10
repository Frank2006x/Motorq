package service

import (
	"context"

	"Frank2006xmotorq/db/sqlc"
)

type VehicleService struct {
	q *sqlc.Queries
}

func NewVehicleService(q *sqlc.Queries) *VehicleService {
	return &VehicleService{q: q}
}

type CreateVehicleRequest struct {
	FleetID int64  `json:"fleet_id"`
	Model   string `json:"model"`
	Status  string `json:"status"`
}

type UpdateVehicleRequest struct {
	ID      int64  `json:"id"`
	FleetID int64  `json:"fleet_id"`
	Model   string `json:"model"`
	Status  string `json:"status"`
}

func (s *VehicleService) Create(ctx context.Context, req CreateVehicleRequest) (sqlc.Vehicle, error) {
	return s.q.CreateVehicle(ctx, sqlc.CreateVehicleParams{
		FleetID: req.FleetID,
		Model:   req.Model,
		Status:  req.Status,
	})
}

func (s *VehicleService) Get(ctx context.Context, id int64) (sqlc.Vehicle, error) {
	return s.q.GetVehicle(ctx, id)
}

func (s *VehicleService) List(ctx context.Context) ([]sqlc.Vehicle, error) {
	return s.q.ListVehicles(ctx)
}

func (s *VehicleService) ListByFleet(ctx context.Context, fleetID int64) ([]sqlc.Vehicle, error) {
	return s.q.ListVehiclesByFleet(ctx, fleetID)
}

func (s *VehicleService) Update(ctx context.Context, req UpdateVehicleRequest) (sqlc.Vehicle, error) {
	return s.q.UpdateVehicle(ctx, sqlc.UpdateVehicleParams{
		ID:      req.ID,
		FleetID: req.FleetID,
		Model:   req.Model,
		Status:  req.Status,
	})
}

func (s *VehicleService) Delete(ctx context.Context, id int64) (sqlc.Vehicle, error) {
	return s.q.DeleteVehicle(ctx, id)
}