package service

import (
	"context"

	"Frank2006xmotorq/db/sqlc"
)

type SimpleRuleFleetService struct {
	q *sqlc.Queries
}

func NewSimpleRuleFleetService(q *sqlc.Queries) *SimpleRuleFleetService {
	return &SimpleRuleFleetService{q: q}
}

type CreateSimpleRuleFleetRequest struct {
	FleetID        int64         `json:"fleet_id"`
	TargetField    string        `json:"target_field"`
	ThresholdValue float64       `json:"threshold_value"`
	Priority       sqlc.Priority `json:"priority"`
}

func (s *SimpleRuleFleetService) Create(ctx context.Context, req CreateSimpleRuleFleetRequest) (sqlc.SimpleRuleFleet, error) {
	return s.q.CreateSimpleRuleFleet(ctx, sqlc.CreateSimpleRuleFleetParams{
		FleetID:        req.FleetID,
		TargetField:    req.TargetField,
		ThresholdValue: floatToNumeric(req.ThresholdValue),
		Priority:       req.Priority,
	})
}

func (s *SimpleRuleFleetService) Get(ctx context.Context, id int64) (sqlc.SimpleRuleFleet, error) {
	return s.q.GetSimpleRuleFleet(ctx, id)
}

func (s *SimpleRuleFleetService) ListByFleet(ctx context.Context, fleetID int64) ([]sqlc.SimpleRuleFleet, error) {
	return s.q.ListSimpleRuleFleetByFleet(ctx, fleetID)
}

func (s *SimpleRuleFleetService) Delete(ctx context.Context, id int64) (sqlc.SimpleRuleFleet, error) {
	return s.q.DeleteSimpleRuleFleet(ctx, id)
}