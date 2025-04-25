package writer

import (
	"bmt_product_service/global"
	"log"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

var (
	writer  *kafka.Writer
	once    sync.Once
	closeCh = make(chan struct{})
	// Store broker list for easy access
	brokerAddresses []string
)

func initKafkaWriter() {
	once.Do(func() {
		brokerAddresses = []string{
			global.Config.ServiceSetting.KafkaSetting.KafkaBroker_1,
			global.Config.ServiceSetting.KafkaSetting.KafkaBroker_2,
			global.Config.ServiceSetting.KafkaSetting.KafkaBroker_3,
		}

		validBrokers := []string{}
		for _, addr := range brokerAddresses {
			if addr != "" {
				validBrokers = append(validBrokers, addr)
			}
		}

		brokerAddresses = validBrokers
		if len(brokerAddresses) == 0 {
			// Exit if no broker is valid
			log.Fatal("KAFKA FATAL: No valid broker addresses configured. Check ServiceSetting.KafkaSetting in config")
		}

		writer = &kafka.Writer{
			Addr:     kafka.TCP(brokerAddresses...),
			Balancer: &kafka.LeastBytes{},
			// Reduce wait times for faster batch submissions
			BatchTimeout: 1000 * time.Millisecond,
			MaxAttempts:  3,
			BatchSize:    100,
			WriteTimeout: 5 * time.Second,
		}
		// global.Logger.Info("kafka producer initialized")
		// log.Println("kafka producer initialized")
	})
}

func getAvailableBroker() string {
	for _, broker := range brokerAddresses {
		conn, err := kafka.Dial("tcp", broker)
		if err == nil {
			conn.Close()
			return broker
		}
	}

	// log.Fatal("no available Kafka broker")
	// global.Logger.Info("no available Kafka broker")

	return ""
}

// Check if a topic already exists on the Kafka broker, and if not, automatically create it
func ensureTopicExists(topic string) error {
	// Connect to Kafka broker
	conn, err := kafka.Dial("tcp", getAvailableBroker())
	if err != nil {
		return err
	}
	defer conn.Close()
	// Check for topic existence
	partitions, err := conn.ReadPartitions(topic)
	if err == nil && len(partitions) > 0 {
		// Topic is exists
		return nil
	}
	// Create topic with 1 partition and replication-factor = 1
	return conn.CreateTopics(kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     3,
		ReplicationFactor: 1,
		ConfigEntries:     []kafka.ConfigEntry{},
	})
}
