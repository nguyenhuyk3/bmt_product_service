package sqlc

import (
	"bmt_product_service/dto/messages"
	"bmt_product_service/dto/request"
	"bmt_product_service/global"
	"bmt_product_service/internal/message_broker/producers"

	"bmt_product_service/utils/convertors"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SqlStore struct {
	connPool *pgxpool.Pool
	*Queries
}

func NewStore(connPool *pgxpool.Pool) *SqlStore {
	return &SqlStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}

func (s *SqlStore) execTran(ctx context.Context, fn func(*Queries) error) error {
	// Start transaction
	tran, err := s.connPool.Begin(ctx)
	if err != nil {
		return err
	}

	q := New(tran)
	// fn performs a series of operations down the db
	err = fn(q)
	if err != nil {
		// If an error occurs, rollback the transaction
		if rbErr := tran.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tran err: %v, rollback err: %v", err, rbErr)
		}

		return err
	}

	return tran.Commit(ctx)
}

func parseDurationToPGInterval(durationStr string) (pgtype.Interval, error) {
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return pgtype.Interval{}, fmt.Errorf("invalid duration format: %v", err)
	}

	return pgtype.Interval{
		Microseconds: duration.Microseconds(),
		Valid:        true,
	}, nil
}

func sendMessage(filmId int32, imageUrl, videoUrl string) error {
	uploadFilmImageMessage := messages.UploadFilmImageMessage{
		ProductId: strconv.Itoa(int(filmId)),
		ImageUrl:  imageUrl,
	}

	err := producers.SendMessage(
		global.UPLOAD_IMAGE_TOPIC,
		strconv.Itoa(int(filmId)),
		uploadFilmImageMessage)
	if err != nil {
		return fmt.Errorf("failed to send upload film image message to kafka: %v", err)
	}

	uploadFilmVideoMessage := messages.UploadFilmVideoMessage{
		ProductId: strconv.Itoa(int(filmId)),
		VideoUrl:  videoUrl,
	}

	err = producers.SendMessage(
		global.UPLOAD_VIDEO_TOPIC,
		strconv.Itoa(int(filmId)),
		uploadFilmVideoMessage)
	if err != nil {
		return fmt.Errorf("failed to send upload film video message to kafka: %v", err)
	}

	return nil
}

func (s *SqlStore) InsertFilmTran(ctx context.Context, arg request.AddProductReq) error {
	err := s.execTran(ctx, func(q *Queries) error {
		interval, err := parseDurationToPGInterval(arg.FilmInformation.Duration)
		if err != nil {
			return err
		}

		releaseDate, err := convertors.GetReleaseDateAsTime(arg.FilmInformation.ReleaseDate)
		if err != nil {
			return err
		}

		filmId, err := q.insertFilm(ctx, insertFilmParams{
			Title:       arg.FilmInformation.Title,
			Description: arg.FilmInformation.Description,
			ReleaseDate: pgtype.Date{
				Time:  releaseDate,
				Valid: true,
			},
			Duration: interval,
		})
		if err != nil {
			return fmt.Errorf("failed to insert film: %v", err)
		}

		err = q.insertFilmChange(ctx, insertFilmChangeParams{
			FilmID:    filmId,
			ChangedBy: arg.FilmChanges.ChangedBy,
			CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
			UpdatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
		})
		if err != nil {
			return fmt.Errorf("failed to insert film change: %v", err)
		}

		for _, genre := range arg.FilmInformation.Genres {
			var tmpGenre NullGenres
			err := tmpGenre.Scan(genre)
			if err != nil {
				return fmt.Errorf("failed to scan role: %v", err)
			}

			err = q.insertFilmGenre(ctx, insertFilmGenreParams{
				FilmID: pgtype.Int4{Int32: filmId, Valid: true},
				Genre: NullGenres{
					Genres: tmpGenre.Genres,
					Valid:  true,
				},
			})
			if err != nil {
				return fmt.Errorf("failed to insert genre %s: %v", genre, err)
			}
		}

		var filmStatus NullStatuses
		if err = filmStatus.Scan(arg.OtherFilmInformation.Status); err != nil {
			return fmt.Errorf("failed to scan status: %v", err)
		}

		err = sendMessage(filmId, arg.OtherFilmInformation.PosterUrl, arg.OtherFilmInformation.TrailerUrl)
		if err != nil {
			return err
		}

		err = q.insertOtherFilmInformation(ctx, insertOtherFilmInformationParams{
			FilmID: filmId,
			Status: NullStatuses{
				Statuses: filmStatus.Statuses,
				Valid:    true,
			},
			PosterUrl: pgtype.Text{
				String: "",
				Valid:  true,
			},
			TrailerUrl: pgtype.Text{
				String: "",
				Valid:  true,
			},
		})
		if err != nil {
			return fmt.Errorf("failed to insert other film information: %v", err)
		}

		return nil
	})

	if err != nil {
		// If the transaction failed, return the error
		return fmt.Errorf("transaction failed: %v", err)
	}

	return nil
}
