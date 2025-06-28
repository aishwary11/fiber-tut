package kafka

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

var Writer *kafka.Writer

func InitProducer(brokerAddress, topic string) {
	Writer = &kafka.Writer{
		Addr:         kafka.TCP(brokerAddress),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireAll,
		Async:        false,
		BatchSize:    100,
		BatchTimeout: 10 * time.Millisecond,
		Compression:  kafka.Snappy,
	}
	log.Println("✅ Kafka producer initialized")
}

func ProduceMessage(key, value string) error {
	if Writer == nil {
		return errors.New("❌ Kafka writer not initialized")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	msg := kafka.Message{
		Key:   []byte(key),
		Value: []byte(value),
		Time:  time.Now(),
	}
	retries := 3
	var err error
	for i := range retries {
		err = Writer.WriteMessages(ctx, msg)
		if err == nil {
			log.Printf("📤 Message sent: key=%s, value=%s\n", key, value)
			return nil
		}
		log.Printf("🔁 Retry %d - Failed to send message: %v", i+1, err)
		time.Sleep(100 * time.Millisecond)
	}
	log.Printf("❌ All retries failed: %v", err)
	return err
}

func CloseProducer() {
	if Writer != nil {
		if err := Writer.Close(); err != nil {
			log.Printf("⚠️ Error closing Kafka writer: %v", err)
		} else {
			log.Println("✅ Kafka writer closed")
		}
	}
}
