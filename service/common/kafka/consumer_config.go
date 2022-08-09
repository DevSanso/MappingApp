package kafka


import (
	module "github.com/confluentinc/confluent-kafka-go/kafka"
)

type ConsumersConfig[T any] struct{
	Address string
	GroupID string
	Offset string
	Topic string
}

func (c *ConsumersConfig[T])toConfigMap() *module.ConfigMap {
	return &module.ConfigMap{
		"bootstrap.servers": c.Address,
		"group.id":          c.GroupID,
		"auto.offset.reset": c.Offset,
	}
}


