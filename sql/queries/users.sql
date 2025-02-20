-- name: SearchUsers :many
SELECT * FROM users
WHERE username LIKE $1 || '%';

-- name: SearchPosts :many
SELECT * FROM posts
WHERE body LIKE $1 || '%';