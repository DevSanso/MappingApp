package kafka_test

import (
	"testing"
	"common/kafka"
)


type TestProducerMessage struct {}

func (t *TestProducerMessage)Key() []byte {return nil}
func (t *TestProducerMessage)Topic() string {return "test"}
func (t *TestProducerMessage)Message() []byte {return []byte("message")}
var producerConfig = kafka.ProducerConfig {
	Address: "192.168.0.3:9092",
}
func TestProducer(t *testing.T) {
	var p = kafka.TestNewProducer(&producerConfig)
	defer p.Close()
	err := p.SendOne(&TestProducerMessage{})


	if err != nil {
		t.Fatal(err)
	}
}