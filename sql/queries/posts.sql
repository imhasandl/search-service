-- name: SearchPosts :many
SELECT * FROM posts
WHERE body LIKE $1 || '%';