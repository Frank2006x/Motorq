-- name: CreateFleet :one
INSERT INTO fleet (name, email)
VALUES ($1, $2)
RETURNING id, name, email;

-- name: GetFleet :one
SELECT id, name, email
FROM fleet
WHERE id = $1;

-- name: ListFleets :many
SELECT id, name, email
FROM fleet
ORDER BY id;

-- name: UpdateFleet :one
UPDATE fleet
SET name = $2, email = $3
WHERE id = $1
RETURNING id, name, email;

-- name: DeleteFleet :one
DELETE FROM fleet
WHERE id = $1
RETURNING id, name, email;