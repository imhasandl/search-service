-- name: SearchPosts :many
SELECT * FROM posts
WHERE body LIKE $1 || '%';

-- name: SearchPostsByDate :many
SELECT * FROM posts
WHERE body LIKE $1 || '%'
ORDER BY created_at;