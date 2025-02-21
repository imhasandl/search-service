-- name: SearchUsers :many
SELECT * FROM users
WHERE username LIKE $1 || '%';