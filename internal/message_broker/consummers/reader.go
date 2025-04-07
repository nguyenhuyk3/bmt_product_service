package consummers

import (
	"bmt_product_service/global"
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaReader struct {
}

var topic = []string{}

func (kr *KafkaReader) InitReaders() {
	log.Println("=============== Product Service is listening for messages... ===============")
}

func (kr *KafkaReader) startReader(topic string) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{global.Config.ServiceSetting.KafkaSetting.KafkaBroker_1},
		GroupID: global.PRODUCT_SERVICE_GROUP,
		Topic:   topic,
	})
	defer reader.Close()

	for {
		_, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v\n", err)
			continue
		}

	}
}
