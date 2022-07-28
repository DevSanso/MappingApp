package kafka_test

import (
	"common/kafka"
	"os"
	"testing"
	"sync"
)

func f(b []byte) (string, error) {
	return string(b), nil
}

var consumerConfig = kafka.ConsumersConfig[string]{
	Address:           "192.168.0.3:9092",
	GroupID:           "test1",
	Offset:            "earliest",
	Count:             1,
	Topic:             "test",
	EncodeMessageFunc: f,
}

func TestConsumer(t *testing.T) {
	var p = kafka.TestNewProducer(&producerConfig)
	defer p.Close()


	var c, err = kafka.TestNewConsumers(&consumerConfig)
	if err != nil {
		t.Fatal(err)
	}
	channel := c.Channel()
	t.Log("start")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		t.Log("read message")
		p.SendOne(&TestProducerMessage{})
		message := <-channel

		if message.Data != "message" {
			t.Error("not matching message :" + message.Data)
			os.Exit(1)
		}
		t.Log("Get message : " + message.Data)
		p.SendOne(&TestProducerMessage{})
		message = <-channel
		if message.Data != "message" {
			t.Error("not matching message :" + message.Data)
			os.Exit(1)
		}
		t.Log("Get message : " + message.Data)
		p.SendOne(&TestProducerMessage{})
		message = <-channel
		if message.Data != "message" {
			t.Error("not matching message :" + message.Data)
			os.Exit(1)
		}
		t.Log("Get message : " + message.Data)
		wg.Done()
	}()
	wg.Wait()
}
