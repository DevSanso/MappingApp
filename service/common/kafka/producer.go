package kafka

import (
	module "github.com/confluentinc/confluent-kafka-go/kafka"
)




type producer[T ByteDecoder] struct {
	origin *module.Producer
}


func newProducer[T ByteDecoder](config *ProducerConfig)producer[T] {
	p,err := module.NewProducer(config.toConfigMap())
	if err != nil {panic(err)}
	return producer[T]{origin : p}
}


func makeMessage[T ByteDecoder](key []byte,topic string,data T) (*module.Message,error) {
	d,err := data.Decode()

	if err != nil {return nil,err}

	if key == nil {
		return &module.Message{
			Value : d,
			TopicPartition: module.TopicPartition{Topic: &topic,Partition: module.PartitionAny},
		},nil
	}
	return &module.Message{
		Value : d,
		TopicPartition: module.TopicPartition{Topic: &topic,Partition: module.PartitionAny},
		Key: key,
	},nil
}

func sendMessage[T ByteDecoder](origin *module.Producer,key []byte,topic string,data T)error {
	message,err := makeMessage(key,topic,data)
	if err != nil {return err}
	return origin.Produce(message, nil)
}

func (p *producer[T])Send(topic string,data []T,key []byte) (err error) {
	for _,message := range data {
		err = sendMessage(p.origin,key,topic,message)
		if err != nil {return}
	}
	return
}

func (p *producer[T])Close() {p.origin.Close()}

