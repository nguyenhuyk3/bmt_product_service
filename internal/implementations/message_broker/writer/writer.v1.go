package writer

import (
	"bmt_product_service/internal/services"
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaWriter struct {
}

// Close implements services.IMessageBroker.
func (k *KafkaWriter) Close() {
	if writer != nil {
		writer.Close()
		log.Println("kafka producer closed")
	}

	close(closeCh)
}

// SendMessage implements services.IMessageBroker.
func (k *KafkaWriter) SendMessage(ctx context.Context, topic string, key string, message interface{}) error {
	if writer == nil {
		initKafkaWriter()
	}

	if err := ensureTopicExists(topic); err != nil {
		// global.Logger.Error("failed to ensure topic exists", zap.Any("err", err))
		return err
	}

	msgBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = writer.WriteMessages(ctx, kafka.Message{
		Topic: topic,
		Key:   []byte(key),
		Value: msgBytes,
	})

	if err != nil {
		// global.Logger.Error("failed to send message to Kafka", zap.Any("err", err))
		return err
	}

	// global.Logger.Error(fmt.Sprintf("message sent to Kafka topic %s", topic))

	return nil
}

func NewKafkaWriter() services.IMessageBroker {
	return &KafkaWriter{}
}
