package service

import (
	"context"

	"Frank2006xmotorq/db/sqlc"
)

type SimpleRuleVehicleService struct {
	q *sqlc.Queries
}

func NewSimpleRuleVehicleService(q *sqlc.Queries) *SimpleRuleVehicleService {
	return &SimpleRuleVehicleService{q: q}
}

type CreateSimpleRuleVehicleRequest struct {
	VehicleID      int64             `json:"vehicle_id"`
	TargetField    string            `json:"target_field"`
	Operator       sqlc.RuleOperator `json:"operator"`
	ThresholdValue float64           `json:"threshold_value"`
	Priority       sqlc.Priority     `json:"priority"`
}

func (s *SimpleRuleVehicleService) Create(ctx context.Context, req CreateSimpleRuleVehicleRequest) (sqlc.SimpleRuleVehicle, error) {
	return s.q.CreateSimpleRuleVehicle(ctx, sqlc.CreateSimpleRuleVehicleParams{
		VehicleID:      req.VehicleID,
		TargetField:    req.TargetField,
		Operator:       req.Operator,
		ThresholdValue: floatToNumeric(req.ThresholdValue),
		Priority:       req.Priority,
	})
}

func (s *SimpleRuleVehicleService) Get(ctx context.Context, id int64) (sqlc.SimpleRuleVehicle, error) {
	return s.q.GetSimpleRuleVehicle(ctx, id)
}

func (s *SimpleRuleVehicleService) ListByVehicle(ctx context.Context, vehicleID int64) ([]sqlc.SimpleRuleVehicle, error) {
	return s.q.ListSimpleRuleVehicleByVehicle(ctx, vehicleID)
}

func (s *SimpleRuleVehicleService) Delete(ctx context.Context, id int64) (sqlc.SimpleRuleVehicle, error) {
	return s.q.DeleteSimpleRuleVehicle(ctx, id)
}