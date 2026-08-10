package service

import (
	"context"
	"fmt"

	"Frank2006xmotorq/db/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

type RuleEvaluator struct {
	q *sqlc.Queries
}

func NewRuleEvaluator(q *sqlc.Queries) *RuleEvaluator {
	return &RuleEvaluator{q: q}
}

// RuleResult holds the result of evaluating a rule
type RuleResult struct {
	RuleID   int64
	Passed   bool
	Priority sqlc.Priority
	Message  string
}

// EvaluateTelemetry evaluates all rules for a given telemetry record
func (e *RuleEvaluator) EvaluateTelemetry(ctx context.Context, telemetry sqlc.TelemetryHistory) ([]RuleResult, error) {
	var results []RuleResult

	// Get vehicle to find fleet_id
	vehicle, err := e.q.GetVehicle(ctx, telemetry.VehicleID)
	if err != nil {
		return nil, fmt.Errorf("failed to get vehicle: %w", err)
	}

	// Fetch vehicle-specific rules
	vehicleRules, err := e.q.ListSimpleRuleVehicleByVehicle(ctx, telemetry.VehicleID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch vehicle rules: %w", err)
	}

	// Fetch fleet-wide rules
	fleetRules, err := e.q.ListSimpleRuleFleetByFleet(ctx, vehicle.FleetID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch fleet rules: %w", err)
	}

	// Evaluate vehicle rules
	for _, rule := range vehicleRules {
		result := e.evaluateVehicleRule(rule, telemetry)
		result.RuleID = rule.ID
		results = append(results, result)
	}

	// Evaluate fleet rules
	for _, rule := range fleetRules {
		result := e.evaluateFleetRule(rule, telemetry)
		result.RuleID = rule.ID
		results = append(results, result)
	}

	return results, nil
}

// evaluateVehicleRule compares a vehicle rule against telemetry data
func (e *RuleEvaluator) evaluateVehicleRule(rule sqlc.SimpleRuleVehicle, telemetry sqlc.TelemetryHistory) RuleResult {
	fieldValue := e.getFieldValue(rule.TargetField, telemetry)
	threshold := e.numericToFloat64(rule.ThresholdValue)

	passed := e.checkRuleViolation(fieldValue, threshold, rule.Operator)
	status := "passed"
	if !passed {
		status = "failed"
	}

	return RuleResult{
		Passed:   passed,
		Priority: rule.Priority,
		Message:  fmt.Sprintf("%s: got %.2f, threshold %.2f (%s) - %s", rule.TargetField, fieldValue, threshold, rule.Operator, status),
	}
}

// evaluateFleetRule compares a fleet rule against telemetry data
func (e *RuleEvaluator) evaluateFleetRule(rule sqlc.SimpleRuleFleet, telemetry sqlc.TelemetryHistory) RuleResult {
	fieldValue := e.getFieldValue(rule.TargetField, telemetry)
	threshold := e.numericToFloat64(rule.ThresholdValue)

	passed := e.checkRuleViolation(fieldValue, threshold, rule.Operator)
	status := "passed"
	if !passed {
		status = "failed"
	}

	return RuleResult{
		Passed:   passed,
		Priority: rule.Priority,
		Message:  fmt.Sprintf("%s: got %.2f, threshold %.2f (%s) - %s", rule.TargetField, fieldValue, threshold, rule.Operator, status),
	}
}

// checkRuleViolation checks if the value violates the rule
// Returns true if rule is PASSED (no violation)
// Returns false if rule is FAILED (violation detected)
func (e *RuleEvaluator) checkRuleViolation(value, threshold float64, operator sqlc.RuleOperator) bool {
	switch operator {
	case ">":
		// Value > threshold = violation (failed)
		return value <= threshold
	case "<":
		// Value < threshold = violation (failed)
		return value >= threshold
	default:
		return true
	}
}

// getFieldValue extracts the numeric value for a given target field
func (e *RuleEvaluator) getFieldValue(targetField string, telemetry sqlc.TelemetryHistory) float64 {
	switch targetField {
	case "speed_mph":
		return e.numericToFloat64(telemetry.SpeedMph)
	case "fuel_level_percent":
		return e.numericToFloat64(telemetry.FuelLevelPercent)
	case "odometer_miles":
		return e.numericToFloat64(telemetry.OdometerMiles)
	default:
		return 0
	}
}

// numericToFloat64 converts pgtype.Numeric to float64
func (e *RuleEvaluator) numericToFloat64(n pgtype.Numeric) float64 {
	if !n.Valid {
		return 0
	}

	if n.Int == nil {
		return 0
	}

	// Convert big.Int to float64
	f, _ := n.Int.Float64()

	// Apply exponent (scale by powers of 10)
	exp := int(n.Exp)
	if exp > 0 {
		for i := 0; i < exp; i++ {
			f *= 10
		}
	} else if exp < 0 {
		for i := 0; i > exp; i-- {
			f /= 10
		}
	}

	return f
}

// CheckRulesAndUpdateStatus checks all rules and updates telemetry status
func (e *RuleEvaluator) CheckRulesAndUpdateStatus(ctx context.Context, telemetryID int64) error {
	telemetry, err := e.q.GetTelemetryHistory(ctx, telemetryID)
	if err != nil {
		return err
	}

	results, err := e.EvaluateTelemetry(ctx, telemetry)
	if err != nil {
		return err
	}

	// Determine overall status
	allPassed := true
	for _, result := range results {
		if !result.Passed {
			allPassed = false
			break
		}
	}

	newStatus := "completed"
	if !allPassed {
		newStatus = "failed"
	}

	// Update telemetry status
	_, err = e.q.UpdateTelemetryHistoryStatus(ctx, sqlc.UpdateTelemetryHistoryStatusParams{
		ID:     telemetryID,
		Status: newStatus,
	})

	return err
}