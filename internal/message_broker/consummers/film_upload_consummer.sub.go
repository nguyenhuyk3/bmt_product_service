package consummers

import (
	"bmt_product_service/db/sqlc"
	"bmt_product_service/dto/messages"
	"bmt_product_service/global"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/segmentio/kafka-go"
)

func (f *FilmUploadConsummer) startReader(topic string) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{global.Config.ServiceSetting.KafkaSetting.KafkaBroker_1},
		GroupID: global.PRODUCT_SERVICE_GROUP,
		Topic:   topic,
	})
	defer reader.Close()

	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v\n", err)
			continue
		}

		f.processMessage(topic, message.Value)
	}
}

func (f *FilmUploadConsummer) processMessage(topic string, value []byte) {
	switch topic {
	case global.RETURNED_IMAGE_OBJECT_KEY_TOPIC:
		var message messages.ReturnedObjectKeyMessage
		if err := json.Unmarshal(value, &message); err != nil {
			log.Printf("failed to unmarshal image message: %v\n", err)
			return
		}

		fmt.Println(message)

		f.handleImageObjectKeyTopic(message)

	case global.RETURNED_VIDEO_OBJECT_KEY_TOPIC:
		var message messages.ReturnedObjectKeyMessage
		if err := json.Unmarshal(value, &message); err != nil {
			log.Printf("failed to unmarshal image message: %v\n", err)
			return
		}

		f.handleVideoObjectKeyTopic(message)

	default:
		log.Printf("unknown topic received: %s\n", topic)
	}
}

func (f *FilmUploadConsummer) handleImageObjectKeyTopic(message messages.ReturnedObjectKeyMessage) {
	productId, err := strconv.Atoi(message.ProductId)
	if err != nil {
		log.Printf("product_id (%s) is not in correct format: %v\n", message.ProductId, err)
		return
	}

	err = f.SqlQuery.UpdatePosterUrlAndCheckStatus(f.Context, sqlc.UpdatePosterUrlAndCheckStatusParams{
		FilmID: int32(productId),
		PosterUrl: pgtype.Text{
			String: message.ObjectKey,
			Valid:  true,
		},
	})
	if err != nil {
		log.Printf("failed to update poster url for film id %d: %v\n", productId, err)
	}
}

func (f *FilmUploadConsummer) handleVideoObjectKeyTopic(message messages.ReturnedObjectKeyMessage) {
	productId, err := strconv.Atoi(message.ProductId)
	if err != nil {
		log.Printf("product_id (%s) is not in correct format: %v\n", message.ProductId, err)
		return
	}

	err = f.SqlQuery.UpdateVideoUrlAndCheckStatus(f.Context, sqlc.UpdateVideoUrlAndCheckStatusParams{
		FilmID: int32(productId),
		TrailerUrl: pgtype.Text{
			String: message.ObjectKey,
			Valid:  true,
		},
	})
	if err != nil {
		log.Printf("failed to update trailer url for film id %d: %v\n", productId, err)
	}
}
