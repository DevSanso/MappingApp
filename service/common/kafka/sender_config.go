package kafka

import (
	module "github.com/confluentinc/confluent-kafka-go/kafka"
)

type SenderConfig struct {
	Address string `json:"address"`
}

func (p *SenderConfig) toConfigMap() *module.ConfigMap {
	return &module.ConfigMap{
		"bootstrap.servers": p.Address,
	}
}
