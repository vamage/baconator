-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY name;

-- name: CreateUser :one
INSERT INTO users (
    name, title,picture, display_name, email, uuid, description, labels, annotations, tags
) VALUES (
             $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
         )
    RETURNING *;

-- name: UpdateUser :one
UPDATE users
set name = $2,
    title = $3,
    picture = $4,
    display_name = $5,
    email = $6,
    description = $4,
    labels = $5,
    annotations = $6

WHERE id = $1
    RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;