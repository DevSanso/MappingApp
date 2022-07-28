package kafka

import (
	"log"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	module "github.com/confluentinc/confluent-kafka-go/kafka"
)

type RecvMessage[T any] struct {
	Header map[string][]byte
	Key    []byte
	Data   T

	Err error
}


func newRecvMessage[T any](origin *module.Message, gen func([]byte) (T,error)) *RecvMessage[T] {
	var res = new(RecvMessage[T])
	for _, header := range origin.Headers {
		res.Header[header.Key] = header.Value
	}
	res.Key = origin.Key
	res.Data,res.Err = gen(origin.Value)

	return res
}


type consumersInterrupt struct {
	destoryChannel chan struct{}
	closeOrigin chan struct{}
}

func newConsumersInterrupt() consumersInterrupt {
	return consumersInterrupt{
		make(chan struct{}),
		make(chan struct{}),
	}
}

type consumers[T any] struct {
	origins []*module.Consumer
	channel chan *RecvMessage[T]

	interrupt consumersInterrupt
	destoryOnce sync.Once
}

var TestNewConsumers = newConsumers[string]

func newConsumers[T any](config *ConsumersConfig[T]) (*consumers[T], error) {
	count := config.Count
	cons := new(consumers[T])
	for range make([]struct{}, count) {
		c, err := module.NewConsumer(config.toConfigMap())
		if err != nil {
			return nil, err
		}
		c.SubscribeTopics([]string{config.Topic},nil)
		cons.origins = append(cons.origins, c)

	}
	cons.channel = make(chan *RecvMessage[T])
	cons.interrupt = newConsumersInterrupt()
	go cons.getLoop(config.EncodeMessageFunc)
	return cons, nil
}



func (c *consumers[T]) readAndServeMessage(gen func([]byte)(T,error)) {
	for _, origin := range c.origins {
		msg, err := origin.ReadMessage(-1)
		
		if err != nil {
			if err.(kafka.Error).Code() != kafka.ErrTimedOut {log.Panic("kafka consumer read error : ",err)}
			continue
		}
		
		
		encode := newRecvMessage(msg, gen)
		c.channel <- encode
		
	}
}

func (c *consumers[T]) getLoop(gen func([]byte) (T,error)) {
	for {
		select {
		case <- c.interrupt.destoryChannel:
			c.destoryOnce.Do(c.destoryInterrupt)
			c.interrupt.closeOrigin <- struct{}{}
			break
		default:
			c.readAndServeMessage(gen)
		}
	}
}
func (c *consumers[T]) Channel() <-chan *RecvMessage[T] { return c.channel }


func (c *consumers[T])destoryInterrupt() {
	
}

func (c *consumers[T])closeOrigin() error {
	var once sync.Once
	var resultErr error

	<-c.interrupt.closeOrigin

	for _, origin := range c.origins {
		err := origin.Close()
		if err != nil {
			once.Do(func() { resultErr = err })
		}
	}
	return resultErr
}

func (c *consumers[T]) Close() error {
	c.interrupt.destoryChannel <- struct{}{}

	return c.closeOrigin()
}
