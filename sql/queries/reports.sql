-- name: SearchReports :many
SELECT * FROM reports
WHERE reported_by LIKE $1 || '%';
