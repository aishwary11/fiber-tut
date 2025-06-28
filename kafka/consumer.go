package kafka

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/segmentio/kafka-go"
)

func StartConsumer(broker, topic, groupID string) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{broker},
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 1,
		MaxBytes: 10e6,
	})
	log.Println("✅ Kafka consumer initialized")
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer reader.Close()
		for {
			select {
			case <-ctx.Done():
				log.Println("🛑 Kafka consumer context cancelled")
				return
			default:
				m, err := reader.ReadMessage(ctx)
				if err != nil {
					log.Printf("❌ Kafka read error: %v", err)
					continue
				}
				log.Printf("📨 Message received: key=%s, value=%s", string(m.Key), string(m.Value))
			}
		}
	}()
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM)
	<-sigchan
	log.Println("📴 Shutdown signal received")
	cancel()
	wg.Wait()
	log.Println("✅ Kafka consumer stopped")
}
