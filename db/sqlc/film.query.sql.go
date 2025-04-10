// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: film.query.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getAllFilms = `-- name: GetAllFilms :many
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
ORDER BY f.release_date DESC
`

type GetAllFilmsRow struct {
	ID          int32           `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	ReleaseDate pgtype.Date     `json:"release_date"`
	Duration    pgtype.Interval `json:"duration"`
	Genres      interface{}     `json:"genres"`
	Status      NullStatuses    `json:"status"`
	PosterUrl   pgtype.Text     `json:"poster_url"`
	TrailerUrl  pgtype.Text     `json:"trailer_url"`
}

func (q *Queries) GetAllFilms(ctx context.Context) ([]GetAllFilmsRow, error) {
	rows, err := q.db.Query(ctx, getAllFilms)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAllFilmsRow{}
	for rows.Next() {
		var i GetAllFilmsRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.ReleaseDate,
			&i.Duration,
			&i.Genres,
			&i.Status,
			&i.PosterUrl,
			&i.TrailerUrl,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFilmByTitle = `-- name: GetFilmByTitle :one
SELECT id, title, description, release_date, duration
FROM films
WHERE title = $1
`

func (q *Queries) GetFilmByTitle(ctx context.Context, title string) (Films, error) {
	row := q.db.QueryRow(ctx, getFilmByTitle, title)
	var i Films
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.ReleaseDate,
		&i.Duration,
	)
	return i, err
}

const updatePosterUrlAndCheckStatus = `-- name: UpdatePosterUrlAndCheckStatus :exec
UPDATE "other_film_informations"
SET poster_url = $2, 
    status = CASE 
        WHEN trailer_url IS NOT NULL
            AND LENGTH(trailer_url) > 0 
            AND LENGTH($2::text) > 0 THEN 'success' 
        ELSE status
    END
WHERE "film_id" = $1
`

type UpdatePosterUrlAndCheckStatusParams struct {
	FilmID    int32       `json:"film_id"`
	PosterUrl pgtype.Text `json:"poster_url"`
}

func (q *Queries) UpdatePosterUrlAndCheckStatus(ctx context.Context, arg UpdatePosterUrlAndCheckStatusParams) error {
	_, err := q.db.Exec(ctx, updatePosterUrlAndCheckStatus, arg.FilmID, arg.PosterUrl)
	return err
}

const updateVideoUrlAndCheckStatus = `-- name: UpdateVideoUrlAndCheckStatus :exec
UPDATE "other_film_informations"
SET trailer_url = $2, 
    status = CASE 
        WHEN poster_url IS NOT NULL 
        AND LENGTH(poster_url) > 0
        AND LENGTH($2::text) > 0 THEN 'success' 
        ELSE status
    END
WHERE "film_id" = $1
`

type UpdateVideoUrlAndCheckStatusParams struct {
	FilmID     int32       `json:"film_id"`
	TrailerUrl pgtype.Text `json:"trailer_url"`
}

func (q *Queries) UpdateVideoUrlAndCheckStatus(ctx context.Context, arg UpdateVideoUrlAndCheckStatusParams) error {
	_, err := q.db.Exec(ctx, updateVideoUrlAndCheckStatus, arg.FilmID, arg.TrailerUrl)
	return err
}

const insertFilm = `-- name: insertFilm :one
INSERT INTO "films" ("title", "description", "release_date", "duration")
VALUES ($1, $2, $3, $4)
RETURNING id
`

type insertFilmParams struct {
	Title       string          `json:"title"`
	Description string          `json:"description"`
	ReleaseDate pgtype.Date     `json:"release_date"`
	Duration    pgtype.Interval `json:"duration"`
}

func (q *Queries) insertFilm(ctx context.Context, arg insertFilmParams) (int32, error) {
	row := q.db.QueryRow(ctx, insertFilm,
		arg.Title,
		arg.Description,
		arg.ReleaseDate,
		arg.Duration,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const insertFilmChange = `-- name: insertFilmChange :exec
INSERT INTO "fillm_changes" ("film_id", "changed_by", "created_at", "updated_at")
VALUES ($1, $2, $3, $4)
`

type insertFilmChangeParams struct {
	FilmID    int32            `json:"film_id"`
	ChangedBy string           `json:"changed_by"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) insertFilmChange(ctx context.Context, arg insertFilmChangeParams) error {
	_, err := q.db.Exec(ctx, insertFilmChange,
		arg.FilmID,
		arg.ChangedBy,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const insertFilmGenre = `-- name: insertFilmGenre :exec
INSERT INTO "film_genres" (film_id, genre)
VALUES ($1, $2)
ON CONFLICT (film_id, genre) DO NOTHING
`

type insertFilmGenreParams struct {
	FilmID pgtype.Int4 `json:"film_id"`
	Genre  NullGenres  `json:"genre"`
}

func (q *Queries) insertFilmGenre(ctx context.Context, arg insertFilmGenreParams) error {
	_, err := q.db.Exec(ctx, insertFilmGenre, arg.FilmID, arg.Genre)
	return err
}

const insertOtherFilmInformation = `-- name: insertOtherFilmInformation :exec
INSERT INTO "other_film_informations" ("film_id","status", "poster_url", "trailer_url")
VALUES ($1, $2, $3, $4)
`

type insertOtherFilmInformationParams struct {
	FilmID     int32        `json:"film_id"`
	Status     NullStatuses `json:"status"`
	PosterUrl  pgtype.Text  `json:"poster_url"`
	TrailerUrl pgtype.Text  `json:"trailer_url"`
}

func (q *Queries) insertOtherFilmInformation(ctx context.Context, arg insertOtherFilmInformationParams) error {
	_, err := q.db.Exec(ctx, insertOtherFilmInformation,
		arg.FilmID,
		arg.Status,
		arg.PosterUrl,
		arg.TrailerUrl,
	)
	return err
}
