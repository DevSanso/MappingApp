package kafka

import (
	module "github.com/confluentinc/confluent-kafka-go/kafka"
)




type Producer[T ByteDecoder] struct {
	origin *module.Producer
}


func NewProducer[T ByteDecoder](config *ProducerConfig) Producer[T] {
	p,err := module.NewProducer(config.toConfigMap())
	if err != nil {panic(err)}
	return Producer[T]{origin : p}
}


func makeMessage[T ByteDecoder](key []byte,topic string,data T) module.Message {
	if(key == nil) {
		return module.Message{
			Value : data.Decode(),
			TopicPartition: module.TopicPartition{Topic: &topic,Partition: module.PartitionAny},
		}
	}
	return module.Message{
		Value : data.Decode(),
		TopicPartition: module.TopicPartition{Topic: &topic,Partition: module.PartitionAny},
		Key: key,
	}
}

func sendMessage[T ByteDecoder](origin *module.Producer,key []byte,topic string,data T)error {
	message := makeMessage(key,topic,data)
	return origin.Produce(&message, nil)
}

func (p *Producer[T])Send(topic string,data []T,key []byte) (err error) {
	for _,message := range data {
		err = sendMessage(p.origin,key,topic,message)
		if err != nil {return}
	}
	return
}

func (p *Producer[T])Close() {p.origin.Close()}

