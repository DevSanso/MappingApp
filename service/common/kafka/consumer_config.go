package kafka


import (
	module "github.com/confluentinc/confluent-kafka-go/kafka"
)

type ConsumerConfig[T ByteEncodeAndWriter] struct{
	Address string
	GroupID string
	Offset string
	Count int
	EncodeNewFunc func() T
}

func (c *ConsumerConfig[T])toConfigMap() *module.ConfigMap {
	return &module.ConfigMap{
		"bootstrap.servers": c.Address,
		"group.id":          c.GroupID,
		"auto.offset.reset": c.Offset,
	}
}


