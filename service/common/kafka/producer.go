package kafka

import (
	"errors"


	module "github.com/confluentinc/confluent-kafka-go/kafka"
)

type producerMessage interface {
	Key() []byte
	Topic() string
	Message() []byte
}

type producer struct {
	origin *module.Producer
}

var TestNewProducer = newProducer

func newProducer(config *ProducerConfig) producer {
	p, err := module.NewProducer(config.toConfigMap())
	if err != nil {
		panic(err)
	}
	return producer{origin: p}
}

func makeMessage(key []byte, topic string, data producerMessage) (*module.Message, error) {
	d := data.Message()

	if key == nil {
		return &module.Message{
			Value:          d,
			TopicPartition: module.TopicPartition{Topic: &topic, Partition: module.PartitionAny},
		}, nil
	}
	return &module.Message{
		Value:          d,
		TopicPartition: module.TopicPartition{Topic: &topic, Partition: module.PartitionAny},
		Key:            key,
	}, nil
}

func sendMessage(origin *module.Producer, key []byte, topic string, data producerMessage) error {
	message, err := makeMessage(key, topic, data)
	if err != nil {
		return err
	}
	if err = origin.Produce(message, nil);err != nil {
		return errors.New("produce send error : " +  err.Error())
	}
	origin.Flush(5)
	return nil
}

func (p *producer) Send(data []producerMessage) (err error) {
	for _, message := range data {
		err = sendMessage(p.origin, message.Key(), message.Topic(), message)
		if err != nil {
			return
		}
	}
	return
}

func (p *producer) SendOne(data producerMessage) (err error) {
	err = sendMessage(p.origin, data.Key(), data.Topic(), data)
	return
}

func (p *producer) Close() { p.origin.Close() }
