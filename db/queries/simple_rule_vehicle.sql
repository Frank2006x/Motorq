-- name: CreateSimpleRuleVehicle :one
INSERT INTO simple_rule_vehicle (vehicle_id, target_field, threshold_value, priority)
VALUES ($1, $2, $3, $4)
RETURNING id, vehicle_id, target_field, threshold_value, priority;

-- name: GetSimpleRuleVehicle :one
SELECT id, vehicle_id, target_field, threshold_value, priority
FROM simple_rule_vehicle
WHERE id = $1;

-- name: ListSimpleRuleVehicleByVehicle :many
SELECT id, vehicle_id, target_field, threshold_value, priority
FROM simple_rule_vehicle
WHERE vehicle_id = $1;

-- name: DeleteSimpleRuleVehicle :one
DELETE FROM simple_rule_vehicle
WHERE id = $1
RETURNING id, vehicle_id, target_field, threshold_value, priority;