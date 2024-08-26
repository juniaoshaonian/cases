package comsume

import (
	"cases/kafka/testwebserver"
	"context"
	kafkago "github.com/segmentio/kafka-go"
	"log"
)

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
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			log.Println(err)
			return
		}
		testwebserver.BatchTest()
		err = c.reader.CommitMessages(ctx, msg)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
