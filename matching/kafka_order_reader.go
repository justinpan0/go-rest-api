package matching

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
	logger "github.com/siddontang/go-log/log"
)

const (
	TopicOrderPrefix = "matching_order_"
)

type KafkaOrderReader struct {
	orderReader *kafka.Reader
}

func NewKafkaOrderReader(productId string, brokers []string) *KafkaOrderReader {
	s := &KafkaOrderReader{}
	logger.Println("NewKafkaOrderReader: consume kafka order stream", "matching_order_"+productId)
	s.orderReader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:   brokers,
		Topic:     "matching_order_" + productId,
		Partition: 0,
		MinBytes:  1,
		MaxBytes:  10e6,
	})
	return s
}

func (s *KafkaOrderReader) SetOffset(offset int64) error {
	return s.orderReader.SetOffset(offset)
}

func (s *KafkaOrderReader) FetchOrder() (offset int64, order *Order, err error) {
	logger.Info("fetch order")

	message, err := s.orderReader.FetchMessage(context.Background())
	logger.Info("after fetch order")
	if err != nil {
		return 0, nil, err
	}
	logger.Info("commit message")
	s.orderReader.CommitMessages(context.Background(), message)
	logger.Info("after commit messsage")
	err = json.Unmarshal(message.Value, &order)
	if err != nil {
		return 0, nil, err
	}
	logger.Info("fetch order end")
	return message.Offset, order, nil
}
