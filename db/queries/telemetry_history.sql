-- name: CreateTelemetryHistory :one
INSERT INTO telemetry_history (vehicle_id, timestamp, speed_mph, fuel_level_percent, engine_state, odometer_miles)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, vehicle_id, timestamp, speed_mph, fuel_level_percent, engine_state, odometer_miles;

-- name: GetTelemetryHistory :one
SELECT id, vehicle_id, timestamp, speed_mph, fuel_level_percent, engine_state, odometer_miles
FROM telemetry_history
WHERE id = $1;

-- name: ListTelemetryHistoryByVehicle :many
SELECT id, vehicle_id, timestamp, speed_mph, fuel_level_percent, engine_state, odometer_miles
FROM telemetry_history
WHERE vehicle_id = $1
ORDER BY timestamp DESC;

-- name: ListLatestTelemetryByVehicle :many
SELECT id, vehicle_id, timestamp, speed_mph, fuel_level_percent, engine_state, odometer_miles
FROM telemetry_history
WHERE vehicle_id = $1
ORDER BY timestamp DESC
LIMIT $2;