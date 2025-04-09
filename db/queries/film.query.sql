-- name: GetFilmByTitle :one
SELECT *
FROM films
WHERE title = $1;

-- name: insertFilm :one
INSERT INTO "films" ("title", "description", "release_date", "duration")
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: insertFilmChange :exec
INSERT INTO "fillm_changes" ("film_id", "changed_by", "created_at", "updated_at")
VALUES ($1, $2, $3, $4);

-- name: insertFilmGenre :exec
INSERT INTO "film_genres" (film_id, genre)
VALUES ($1, $2)
ON CONFLICT (film_id, genre) DO NOTHING;

-- name: insertOtherFilmInformation :exec 
INSERT INTO "other_film_informations" ("film_id","status", "poster_url", "trailer_url")
VALUES ($1, $2, $3, $4);

-- name: UpdatePosterUrlAndCheckStatus :exec
UPDATE "other_film_informations"
SET poster_url = $2, 
    status = CASE 
        WHEN trailer_url IS NOT NULL
            AND LENGTH(trailer_url) > 0 
            AND LENGTH($2::text) > 0 THEN 'success' 
        ELSE status
    END
WHERE "film_id" = $1;

-- name: UpdateVideoUrlAndCheckStatus :exec
UPDATE "other_film_informations"
SET trailer_url = $2, 
    status = CASE 
        WHEN poster_url IS NOT NULL 
        AND LENGTH(poster_url) > 0
        AND LENGTH($2::text) > 0 THEN 'success' 
        ELSE status
    END
WHERE "film_id" = $1;

-- name: GetAllFilms :many
SELECT 
    f.id, f.title, f.description, f.release_date, f.duration,
    ARRAY_AGG(DISTINCT fg.genre::text) AS genres,
    ofi.status, ofi.poster_url, ofi.trailer_url
FROM films AS f
LEFT JOIN other_film_informations AS ofi ON f.id = ofi.film_id
LEFT JOIN film_genres AS fg ON fg.film_id = f.id
GROUP BY 
    f.id, f.title, f.description, f.release_date, f.duration,
    ofi.status, ofi.poster_url, ofi.trailer_url
ORDER BY f.release_date DESC;




