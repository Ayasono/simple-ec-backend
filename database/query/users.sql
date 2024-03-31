-- name: CreateUser :many
INSERT INTO users (username,
                   email,
                   password_hash,
                   address,
                   phone)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, username, email, address, phone;

-- name: ListUsers :many
SELECT id,
       username,
       email
FROM users
ORDER BY id;

-- name: GetUserByEmail :one
SELECT username,
       email,
       address,
       phone
FROM users
WHERE email = $1;

-- name: GetUserByID :one
SELECT id,
       username,
       email,
       password_hash
FROM users
WHERE id = $1;

-- name: CheckUserPassword :one
SELECT password_hash
FROM users
WHERE email = $1;
