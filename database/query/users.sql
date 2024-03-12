-- name: CreateUser :many
INSERT INTO users (username,
                   email,
                   password_hash,
                   created_at,
                   updated_at)
VALUES ($1, $2, $3, NOW(), NOW())
RETURNING id, username, email, created_at, updated_at;

-- name: ListUsers :many
SELECT id,
       username,
       email,
       created_at,
       updated_at
FROM users
ORDER BY id;

-- name: GetUserByEmail :one
SELECT id,
       username,
       email,
       password_hash,
       created_at,
       updated_at
FROM users
WHERE email = $1;

-- name: GetUserByID :one
SELECT id,
       username,
       email,
       password_hash,
       created_at,
       updated_at
FROM users
WHERE id = $1;

-- name: CheckUserPassword :one
SELECT password_hash
FROM users
WHERE email = $1;
