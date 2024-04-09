-- name: GetUser :one
SELECT * FROM users
WHERE id = $1;

-- name: InsertUser :one
INSERT INTO users(
  "username",
  "role"
)  VALUES(
    $1, $2
)
RETURNING *;

-- name: UpdateRole :exec
UPDATE users SET role = $2
WHERE id = $1;

-- name: DeleteItem :exec
DELETE FROM users
WHERE ID = $1;