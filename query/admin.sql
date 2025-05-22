-- name: CountAdmins :one
SELECT COUNT(*) FROM users
WHERE type = 'admin';

-- name: CreateAdmin :exec
INSERT INTO users (name, email, password, type)
VALUES (?, ?, ?, 'admin');