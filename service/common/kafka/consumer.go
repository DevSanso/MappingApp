package kafka

import (
	"common/kafka/test"


	"log"
	"sync"
	"unsafe"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	module "github.com/confluentinc/confluent-kafka-go/kafka"
	"google.golang.org/protobuf/proto"

)

type RecvMessage[T any] struct {
	Header map[string][]byte
	Key    []byte
	Data   *T

	Err error
}


func newRecvMessage[T any](origin *module.Message)  *RecvMessage[T] {
	var res = new(RecvMessage[T])
	for _, header := range origin.Headers {
		res.Header[header.Key] = header.Value
	}
	res.Key = origin.Key

	originData := new(T)
	var data = unsafe.Pointer(originData)
	dataBoxing,isCast := (*(*interface{})(data)).(proto.Message)
	if !isCast {
		panic("not cast consumer  generic type to proto.Message")
	}

	res.Err = proto.Unmarshal(origin.Value,dataBoxing)
	res.Data = originData
	return res
}

type consumersInterrupt struct {
	destoryChannel chan struct{}
	closeOrigin    chan struct{}
}

func newConsumersInterrupt() consumersInterrupt {
	return consumersInterrupt{
		make(chan struct{}),
		make(chan struct{}),
	}
}

type consumer[T any] struct {
	origin *module.Consumer
	channel chan *RecvMessage[T]

	interrupt   consumersInterrupt
	destoryOnce sync.Once
}

var TestNewConsumers = newConsumer[test.Request]

func newConsumer[T any](config *ConsumersConfig[T]) (*consumer[T], error) {
	cons := new(consumer[T])
	c, err := module.NewConsumer(config.toConfigMap())
	if err != nil {
		return nil, err
	}
	c.SubscribeTopics([]string{config.Topic}, nil)
	cons.origin = c
	cons.channel = make(chan *RecvMessage[T])
	cons.interrupt = newConsumersInterrupt()
	go cons.getLoop()
	return cons, nil
}

func (c *consumer[T]) readAndServeMessage()  {
	msg, err := c.origin.ReadMessage(-1)
	if err != nil {
		if err.(kafka.Error).Code() != kafka.ErrTimedOut {
			log.Panic("kafka consumer read error : ", err)
		}
	}
	encode := newRecvMessage[T](msg)
	c.channel <- encode
}

func (c *consumer[T]) getLoop() (T, error) {
	for {
		select {
		case <-c.interrupt.destoryChannel:
			c.destoryOnce.Do(c.destoryInterrupt)
			c.interrupt.closeOrigin <- struct{}{}
			break
		default:
			c.readAndServeMessage()
		}
	}
}
func (c *consumer[T]) Channel() <-chan *RecvMessage[T] { return c.channel }

func (c *consumer[T]) destoryInterrupt() {

}

func (c *consumer[T]) closeOrigin() error {
	var resultErr error

	<-c.interrupt.closeOrigin
	resultErr = c.origin.Close()
	
	return resultErr
}

func (c *consumer[T]) Close() error {
	c.interrupt.destoryChannel <- struct{}{}

	return c.closeOrigin()
}
