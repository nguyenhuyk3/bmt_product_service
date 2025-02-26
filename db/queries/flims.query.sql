-- name: GetFilmByTitle :one
SELECT *
FROM films
WHERE title = $1;
