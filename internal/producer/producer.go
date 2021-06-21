package producer

import (
	"context"
	"errors"

	"github.com/Shopify/sarama"

	"github.com/ozoncp/ocp-prize-api/internal/configuration"
	"github.com/rs/zerolog/log"
)

var brokers = []string{"kafka:9093", "kafka:9094"}

// IProducer Implement interface for kafka producer
type IProducer interface {
	SendMessage(msg string) bool
	Close() error
}

// ProducerState enum for producer states
type ProducerState int

const (
	// ProducerNotInitialized is setted when producer isn't initialized or close
	ProducerNotInitialized ProducerState = iota
	// ProducerInitialized is setted whenproducer initialized
	ProducerInitialized
)

// Producer struct contains sarama producer
type Producer struct {
	AsyncProducer sarama.AsyncProducer
	Topic         string
	State         ProducerState
}

// NewProducer creates new producer with setted topic
func NewProducer(ctx context.Context, topic string) IProducer {
	val := ctx.Value("configuration")
	var conf *configuration.Configuration
	if val != nil {
		conf = val.(*configuration.Configuration)
	}
	if conf != nil {
		brokers = conf.KafkaBrokers
	}
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 5
	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		return nil
	}
	return &Producer{
		AsyncProducer: producer,
		Topic:         topic,
		State:         ProducerInitialized,
	}
}

// SendMessage from producer to brocker
func (prod *Producer) SendMessage(msg string) bool {

	if prod.State == ProducerNotInitialized {
		log.Printf("Producer is not initialized")
		return false
	}
	go func() {
		saramaMsg := prepareMessage(prod.Topic, msg)
		prod.AsyncProducer.Input() <- saramaMsg

		select {
		case msg := <-prod.AsyncProducer.Successes():
			log.Printf("Produced message successes: [%s]\n", msg.Value)
		case err := <-prod.AsyncProducer.Errors():
			log.Printf("Produced message failure: %s", err)
		}
	}()
	return true
}

// Close producer
func (prod *Producer) Close() error {
	if prod.State == ProducerNotInitialized {
		log.Print("Can't close producer: Producer is not initialized")
		return errors.New("Producer is not initialized")
	}
	prod.State = ProducerNotInitialized
	err := prod.AsyncProducer.Close()
	if err != nil {
		log.Printf("Can't close producer: %s", err.Error())
	}
	return err
}

func prepareMessage(topic, message string) *sarama.ProducerMessage {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(message),
	}
	return msg
}
