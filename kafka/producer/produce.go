package producer

import (
	"context"
	"fmt"
	kafkago "github.com/segmentio/kafka-go"
	"log"

	"time"
)

const topic = "sequence_comsume"

type TestProducer struct {
	writer *kafkago.Writer
}

func (t *TestProducer) Produce() {
	val := []byte(fmt.Sprintf("%d", time.Now().UnixMilli()))
	msg := kafkago.Message{
		Key: []byte( "test"),
		Value: val,
	}
	err := t.writer.WriteMessages(context.Background(), msg)
	if err != nil {
		log.Printf("发送消息失败 %v",err)
		return
	}
}

func NewTestProducer() *TestProducer {
	// 配置Kafka写入器
	writer := &kafkago.Writer{
		Addr:     kafkago.TCP("localhost:9092"),
		Topic:    topic,
		Balancer: &kafkago.Hash{},
	}
	return &TestProducer{
		writer: writer,
	}
}
