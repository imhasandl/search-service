-- name: SearchUsers :many
SELECT * FROM users
WHERE username LIKE $1 || '%';

-- name: SearchUsersByDate :many
SELECT * FROM users
WHERE username LIKE $1 || '%'
ORDER BY created_at;