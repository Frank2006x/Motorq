-- name: CreateTelemetryHistory :one
INSERT INTO telemetry_history (vehicle_id, timestamp, speed_mph, fuel_level_percent, engine_state, odometer_miles, status)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, vehicle_id, timestamp, speed_mph, fuel_level_percent, engine_state, odometer_miles, status;

-- name: GetTelemetryHistory :one
SELECT id, vehicle_id, timestamp, speed_mph, fuel_level_percent, engine_state, odometer_miles, status
FROM telemetry_history
WHERE id = $1;

-- name: ListTelemetryHistoryByVehicle :many
SELECT id, vehicle_id, timestamp, speed_mph, fuel_level_percent, engine_state, odometer_miles, status
FROM telemetry_history
WHERE vehicle_id = $1
ORDER BY timestamp DESC;

-- name: ListLatestTelemetryByVehicle :many
SELECT id, vehicle_id, timestamp, speed_mph, fuel_level_percent, engine_state, odometer_miles, status
FROM telemetry_history
WHERE vehicle_id = $1
ORDER BY timestamp DESC
LIMIT $2;

-- name: UpdateTelemetryHistoryStatus :one
UPDATE telemetry_history
SET status = $2
WHERE id = $1
RETURNING id, vehicle_id, timestamp, speed_mph, fuel_level_percent, engine_state, odometer_miles, status;