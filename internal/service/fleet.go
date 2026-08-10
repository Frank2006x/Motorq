package service

import (
	"context"

	"Frank2006xmotorq/db/sqlc"
)

type FleetService struct {
	q *sqlc.Queries
}

func NewFleetService(q *sqlc.Queries) *FleetService {
	return &FleetService{q: q}
}

type CreateFleetRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateFleetRequest struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (s *FleetService) Create(ctx context.Context, req CreateFleetRequest) (sqlc.Fleet, error) {
	return s.q.CreateFleet(ctx, sqlc.CreateFleetParams{
		Name:  req.Name,
		Email: req.Email,
	})
}

func (s *FleetService) Get(ctx context.Context, id int64) (sqlc.Fleet, error) {
	return s.q.GetFleet(ctx, id)
}

func (s *FleetService) List(ctx context.Context) ([]sqlc.Fleet, error) {
	return s.q.ListFleets(ctx)
}

func (s *FleetService) Update(ctx context.Context, req UpdateFleetRequest) (sqlc.Fleet, error) {
	return s.q.UpdateFleet(ctx, sqlc.UpdateFleetParams{
		ID:    req.ID,
		Name:  req.Name,
		Email: req.Email,
	})
}

func (s *FleetService) Delete(ctx context.Context, id int64) (sqlc.Fleet, error) {
	return s.q.DeleteFleet(ctx, id)
}