package batch_comsume_batch_commit

import (
	"cases/kafka/testwebserver"
	"context"
	kafkago "github.com/segmentio/kafka-go"
	"log"
)

const DefaultBatchSize = 10

type Consumer struct {
	reader *kafkago.Reader
}

func NewConsumer(topic, groupId string) *Consumer {
	reader := kafkago.NewReader(kafkago.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   topic,
		GroupID: groupId,
	})
	return &Consumer{reader: reader}
}

func (c *Consumer) Consume(ctx context.Context) {
	for {
		if ctx.Done() != nil {
			return
		}
		msgs := make([]kafkago.Message, 0, DefaultBatchSize)
		for i := 0; i < DefaultBatchSize; i++ {
			msg, err := c.reader.ReadMessage(ctx)
			if err != nil {
				log.Println(err)
				return
			}
			msgs = append(msgs, msg)
		}
		testwebserver.BatchTest()
		err := c.reader.CommitMessages(ctx, msgs...)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
