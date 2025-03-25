-- name: SearchReports :many
SELECT * FROM reports
WHERE reported_by LIKE $1 || '%';

-- name: SearchReportsByDate :many
SELECT * FROM reports
WHERE reported_by LIKE $1 || '%'
ORDER BY reported_at;