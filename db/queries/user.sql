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
