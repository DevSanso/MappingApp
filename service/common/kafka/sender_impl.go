package kafka

import (
	k "github.com/confluentinc/confluent-kafka-go/kafka"
	"google.golang.org/protobuf/proto"
)

type senderImpl struct {
	producer *k.Producer
}

func NewSender(config *SenderConfig) (Sender,error) {
	p,err := k.NewProducer(config.toConfigMap())
	if err != nil {return nil,err}

	return &senderImpl {
		producer: p,
	},nil
}

func (si *senderImpl)makeMessage(key []byte, topic string, data proto.Message) (*k.Message, error) {
	d,err := proto.Marshal(data)
	if err != nil {
		return nil,err
	}
	if key == nil {
		return &k.Message{
			Value:          d,
			TopicPartition: k.TopicPartition{Topic: &topic, Partition:k.PartitionAny},
		}, nil
	}
	return &k.Message{
		Value:          d,
		TopicPartition: k.TopicPartition{Topic: &topic, Partition: k.PartitionAny},
		Key:            key,
	}, nil
}

func (si *senderImpl)sendVars(key,topic string,idl proto.Message) (chan k.Event,*k.Message,error) {
	e,err := si.makeMessage([]byte(key),topic,idl)
	return make(chan k.Event),e,err
}

func (si *senderImpl)Send(key,topic string,idl proto.Message) error {
	channel,msg,err := si.sendVars(key,topic,idl)
	defer close(channel)
	if err != nil {return err}

	if err = si.producer.Produce(msg,channel);err != nil {
		return err
	}
	si.producer.Flush(1000)
	
	switch ev := (<- channel).(type) {
	case *k.Message:
		if ev.TopicPartition.Error != nil {
			return ev.TopicPartition.Error
		}
	}
	return nil
}
func (si *senderImpl)Close() error {
	return si.Close()
}