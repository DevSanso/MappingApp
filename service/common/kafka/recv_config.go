package kafka

import (
	module "github.com/confluentinc/confluent-kafka-go/kafka"
)

type RecvConfig struct {
	Address string `json:"address"`
	GroupID string `json:"groupID"`
	Offset  string `json:"offset"`
	Topic   string `json:"topic"`
}

func (c *RecvConfig) toConfigMap() *module.ConfigMap {
	return &module.ConfigMap{
		"bootstrap.servers": c.Address,
		"group.id":          c.GroupID,
		"auto.offset.reset": c.Offset,
	}
}
