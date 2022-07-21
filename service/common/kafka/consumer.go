package kafka

import (
	"sync"
	"log"
	module "github.com/confluentinc/confluent-kafka-go/kafka"
)


type RecvMessage[T ByteEncodeAndWriter] struct {
	Header map[string][]byte
	Key []byte
	Data T
}

func newRecvMessage[T ByteEncodeAndWriter](origin *module.Message,gen func() T) *RecvMessage[T] {
	var res = new(RecvMessage[T])
	for _,header := range origin.Headers {
		res.Header[header.Key] = header.Value
	}
	res.Key = origin.Key
	res.Data = gen()
	err := res.Data.EncodeAndWrite(origin.Value)
	if err != nil {panic(err)}

	return res
}

type Consumer[T ByteEncodeAndWriter] struct {
	origins []*module.Consumer
	channel chan *RecvMessage[T]
}


func NewConsumer[T ByteEncodeAndWriter](config *ConsumerConfig[T]) (*Consumer[T],error) {
	count := config.Count
	cons := new(Consumer[T])
	for range make([]struct{},count) {
		c,err := module.NewConsumer(config.toConfigMap())
		if err != nil {return nil,err}
		cons.origins = append(cons.origins,c)

	}
	cons.channel = make(chan *RecvMessage[T])
	go cons.getLoop(config.EncodeNewFunc)
	return cons,nil
}

func (c *Consumer[T])getLoop(gen func() T) {
	for {
		for _,origin := range c.origins {
			msg,err := origin.ReadMessage(-1)
			if err != nil {log.Println(err);continue}
			encode := newRecvMessage(msg,gen)
			c.channel <- encode
		}
	}
}
func (c *Consumer[T])Channel() <- chan *RecvMessage[T] {return c.channel}

func (c *Consumer[T])Close() error {
	var once sync.Once
	var resultErr error
	close(c.channel)
	for _,origin := range c.origins {
		err := origin.Close()
		if err != nil {
			once.Do(func() {resultErr = err})
		}
	}
	return resultErr
}