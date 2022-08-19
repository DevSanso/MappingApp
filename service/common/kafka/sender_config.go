package kafka

import (
	module "github.com/confluentinc/confluent-kafka-go/kafka"
)

type SenderConfig struct {
	Address string
}

func (p *SenderConfig) toConfigMap() *module.ConfigMap {
	return &module.ConfigMap{
		"bootstrap.servers": p.Address,
	}
}
