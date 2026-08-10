-- name: CreateSimpleRuleFleet :one
INSERT INTO simple_rule_fleet (fleet_id, target_field, threshold_value, priority)
VALUES ($1, $2, $3, $4)
RETURNING id, fleet_id, target_field, threshold_value, priority;

-- name: GetSimpleRuleFleet :one
SELECT id, fleet_id, target_field, threshold_value, priority
FROM simple_rule_fleet
WHERE id = $1;

-- name: ListSimpleRuleFleetByFleet :many
SELECT id, fleet_id, target_field, threshold_value, priority
FROM simple_rule_fleet
WHERE fleet_id = $1;

-- name: DeleteSimpleRuleFleet :one
DELETE FROM simple_rule_fleet
WHERE id = $1
RETURNING id, fleet_id, target_field, threshold_value, priority;