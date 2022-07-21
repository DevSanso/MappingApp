package kafka


import (
	module "github.com/confluentinc/confluent-kafka-go/kafka"
)

type ProducerConfig struct{
	Address string
}


func (p *ProducerConfig)toConfigMap() *module.ConfigMap {
	return &module.ConfigMap{
		"bootstrap.servers": p.Address,
	}
}
