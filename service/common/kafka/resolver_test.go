package kafka_test

import (
	"testing"
	"time"
	"common/kafka"
)



func TestResolver(t *testing.T) {
	var p = kafka.TestNewProducer(&producerConfig)
	defer p.Close()

	var resolver,err = kafka.NewKafkaResolver[*TestProducerMessage,string](&kafka.ResolverConfig[string]{
		BufferTick: time.Second,
		ConsumersConfig: &consumerConfig,
		ErrorHandle: func(err error) {t.Error(err)},
		ProducerConfig: []*kafka.ProducerConfig{&producerConfig},
	})


	if err != nil {
		t.Fatal(err)
	}
	count := 0
	var testHandle = func(send kafka.SendMessage,key []byte,header map[string][]byte, data string) {
		if data == "message" {
			count += 1
			send.Ok("test",[]byte("send to"))
		}
		
	}
	resolver.AddHandle("testing",testHandle)
	cancalFn := resolver.Start()
	
	send := 0
	go func() {
		tick := time.NewTicker(time.Second  * 10)
		for  {
			select {
			case <-tick.C:
			default:
				p.SendOne(&TestProducerMessage{})
				send += 1
			}
		}
	}()
	

	time.Sleep(time.Second * 10)
	if count == 0 {
		t.Error("not recived  all message : ",count," send message : ",send)
	}
	cancalFn()
}