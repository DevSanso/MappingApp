package kafka

import (
	"log"
	"sync"

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

type consumers[T any] struct {
	origins []*module.Consumer
	channel chan *RecvMessage[T]
}

func newConsumers[T any](config *ConsumersConfig[T]) (*consumers[T], error) {
	count := config.Count
	cons := new(consumers[T])
	for range make([]struct{}, count) {
		c, err := module.NewConsumer(config.toConfigMap())
		if err != nil {
			return nil, err
		}
		cons.origins = append(cons.origins, c)

	}
	cons.channel = make(chan *RecvMessage[T])
	go cons.getLoop(config.EncodeMessageFunc)
	return cons, nil
}

func (c *consumers[T]) getLoop(gen func([]byte) (T,error)) {
	for {
		for _, origin := range c.origins {
			msg, err := origin.ReadMessage(-1)
			if err != nil {
				log.Println(err)
				continue
			}
			encode := newRecvMessage(msg, gen)
			c.channel <- encode
		}
	}
}
func (c *consumers[T]) Channel() <-chan *RecvMessage[T] { return c.channel }

func (c *consumers[T]) Close() error {
	var once sync.Once
	var resultErr error
	close(c.channel)
	for _, origin := range c.origins {
		err := origin.Close()
		if err != nil {
			once.Do(func() { resultErr = err })
		}
	}
	return resultErr
}
