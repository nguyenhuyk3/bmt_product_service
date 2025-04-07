package consummers

import (
	"bmt_product_service/db/sqlc"
	"bmt_product_service/global"
	"context"
	"log"
)

type FilmUploadConsummer struct {
	SqlQuery sqlc.Querier
	Context  context.Context
}

var topics = []string{
	global.RETURNED_IMAGE_OBJECT_KEY_TOPIC,
	global.RETURNED_VIDEO_OBJECT_KEY_TOPIC,
}

func NewFilmUploadConsummer(sqlQuery *sqlc.Queries) *FilmUploadConsummer {
	return &FilmUploadConsummer{
		SqlQuery: sqlQuery,
		Context:  context.Background(),
	}
}

func (f *FilmUploadConsummer) InitReaders() {
	log.Println("=============== Product Service is listening for film uploading messages... ===============")

	for _, topic := range topics {
		go f.startReader(topic)
	}
}
