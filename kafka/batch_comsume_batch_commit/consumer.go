package batch_comsume_batch_commit

import (
	"context"
	kafkago "github.com/segmentio/kafka-go"
	"log"
	"time"
)

const DefaultBatchSize = 10

type Consumer struct {
	reader *kafkago.Reader
	count  int64
}

func NewConsumer(topic, groupId string) *Consumer {
	reader := kafkago.NewReader(kafkago.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   topic,
		GroupID: groupId,
	})
	return &Consumer{reader: reader}
}

func (c *Consumer) Consume() {
	for {
		msgs := make([]kafkago.Message, 0, DefaultBatchSize)
		for i := 0; i < DefaultBatchSize; i++ {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			msg, err := c.reader.ReadMessage(ctx)
			if err != nil {
				log.Println(err)
				continue
			}

		}
	}
}
