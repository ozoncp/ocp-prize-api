package producer

import (
	"github.com/Shopify/sarama"

	"github.com/rs/zerolog/log"
)

var brokers = []string{"127.0.0.1:9094"}

type IProducer interface {
	SendMessage(msg string) bool
	Close() error
}

type Producer struct {
	SyncProducer sarama.SyncProducer
	Topic        string
}

func NewProducer(topic string) IProducer {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.ClientID = "123"
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil
	}
	return &Producer{
		SyncProducer: producer,
		Topic:        topic,
	}
}

func (prod *Producer) SendMessage(msg string) bool {
	saramaMsg := prepareMessage(prod.Topic, msg)
	_, _, err := prod.SyncProducer.SendMessage(saramaMsg)
	if err != nil {
		log.Printf(err.Error())
		return false
	}
	return true
}

func (prod *Producer) Close() error {
	return prod.SyncProducer.Close()
}

func prepareMessage(topic, message string) *sarama.ProducerMessage {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(message),
	}
	return msg
}
