-- name: CreateVehicle :one
INSERT INTO vehicle (fleet_id, model, status)
VALUES ($1, $2, $3)
RETURNING id, fleet_id, model, status;

-- name: GetVehicle :one
SELECT id, fleet_id, model, status
FROM vehicle
WHERE id = $1;

-- name: ListVehicles :many
SELECT id, fleet_id, model, status
FROM vehicle
ORDER BY id;

-- name: ListVehiclesByFleet :many
SELECT id, fleet_id, model, status
FROM vehicle
WHERE fleet_id = $1
ORDER BY id;

-- name: UpdateVehicle :one
UPDATE vehicle
SET fleet_id = $2, model = $3, status = $4
WHERE id = $1
RETURNING id, fleet_id, model, status;

-- name: DeleteVehicle :one
DELETE FROM vehicle
WHERE id = $1
RETURNING id, fleet_id, model, status;