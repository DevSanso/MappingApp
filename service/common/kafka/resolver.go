package kafka

import (
	"common/buffer"
	"context"
	"errors"
	"sync"
	"unsafe"
	"time"

	"google.golang.org/protobuf/proto"
)

type producerMessageImpl struct {
	MessageKey   []byte
	MessageTopic string
	MessageData  proto.Message
}

func (c *producerMessageImpl) Key() []byte     { return c.MessageKey }
func (c *producerMessageImpl) Topic() string   { return c.MessageTopic }
func (c *producerMessageImpl) Message() proto.Message { return c.MessageData }

type KafkaResolverHandle[C any] func(send SendMessage, key []byte, header map[string][]byte, data C)

type kafkaResolver[C any] struct {
	buf          *buffer.Buffer[producerMessage]
	producerPool sync.Pool
	consumers    *consumer[C]
	handler      map[string]KafkaResolverHandle[C]

	handleMut sync.Mutex

	ticker        time.Ticker
	producerUseWg sync.WaitGroup

	errorHandle func(error)
}


type ResolverConfig[C any] struct {
	BufferTick time.Duration

	ProducerConfig  []*ProducerConfig
	ConsumersConfig *ConsumersConfig[C]

	ErrorHandle func(error)
}

func makeProducerPool(configs []*ProducerConfig) sync.Pool {
	var p = sync.Pool{}
	p.New = func() interface{} { return nil }
	for _, cfg := range configs {
		p.Put(newProducer(cfg))
	}
	return p
}

type Resolver[REQ any] interface {
	AddHandle(name string, handle KafkaResolverHandle[REQ])
	DeleteHandle(name string)
	Start() context.CancelFunc
}
func chkGenericConsumerType[C any]() {
	originData := new(C)
	var data = unsafe.Pointer(originData)
	_,isCast := (*(*interface{})(data)).(proto.Message)
	if !isCast {
		panic("not cast consumer  generic type to proto.Message")
	}
}

func NewKafkaResolver[P producerMessage, C any](
	config *ResolverConfig[C],
) (Resolver[C], error) {

	chkGenericConsumerType[C]()

	buf := buffer.NewBuffer[producerMessage]()
	producer := makeProducerPool(config.ProducerConfig)
	cons, err := newConsumer(config.ConsumersConfig)

	if err != nil {
		return nil, err
	}

	var resolver = &kafkaResolver[C]{
		buf:          buf,
		producerPool: producer,
		consumers:    cons,
		errorHandle:  config.ErrorHandle,

		ticker:        *time.NewTicker(config.BufferTick),
		producerUseWg: sync.WaitGroup{},
		handler:       make(map[string]KafkaResolverHandle[C]),
	}

	return resolver, nil
}

func (resovler *kafkaResolver[C]) AddHandle(name string, handle KafkaResolverHandle[C]) {
	resovler.handleMut.Lock()
	resovler.handler[name] = handle
	resovler.handleMut.Unlock()
}
func (resovler *kafkaResolver[C]) DeleteHandle(name string) {
	resovler.handleMut.Lock()
	delete(resovler.handler, name)
	resovler.handleMut.Unlock()
}
func (resovler *kafkaResolver[C]) Start() context.CancelFunc {
	ctx, cancalFn := context.WithCancel(context.Background())
	go resovler.loop(ctx)
	return cancalFn
}

func (resovler *kafkaResolver[C]) deferFn() {
	err := resovler.consumers.Close()
	if err != nil {
		resovler.errorHandle(errors.New("DeferFn Error : " + err.Error()))
	}
	resovler.producerUseWg.Wait()
	for p := resovler.producerPool.Get(); p != nil; {
		cast := p.(producer)
		cast.Close()
	}
}

func (resovler *kafkaResolver[C]) recvMessages(messages []producerMessage) {
	if len(messages) == 0 {
		return
	}

	for {
		p, ok := resovler.producerPool.Get().(producer)
		if !ok {
			continue
		}

		resovler.producerUseWg.Add(1)
		defer resovler.producerUseWg.Done()
		defer resovler.producerPool.Put(p)

		err := p.Send(messages)
		if err != nil {
			resovler.errorHandle(errors.New("recvMessage Error : " + err.Error()))
		}

		break
	}
}

func (resovler *kafkaResolver[C]) call(msg *RecvMessage[C]) {
	if msg.Err != nil {
		resovler.errorHandle(msg.Err)
		return
	}
	go func() {
		resovler.handleMut.Lock()
		shaowCopy := resovler.handler
		resovler.handleMut.Unlock()

		resovler.callHandles(msg,shaowCopy)
		
	}()
}
func (resovler *kafkaResolver[C]) callHandles(msg *RecvMessage[C],copyd map[string]KafkaResolverHandle[C]) {
	for _, handle := range copyd {
		var send = new(sendMessageImpl)
		handle(send, msg.Key, msg.Header, *msg.Data)

		if send.err != nil {
			resovler.errorHandle(errors.New("call Error : " + send.err.Error()))
			return
		}
		if send.topic == "" {
			continue
		}
		sendMessage := &producerMessageImpl{
			MessageKey:   msg.Key,
			MessageTopic: send.topic,
			MessageData:  send.message,
		}
		resovler.buf.Push(sendMessage)
	}
}

func (resovler *kafkaResolver[C]) loop(ctx context.Context) {
	var ch = resovler.consumers.Channel()
Loop:
	for {
		select {
		case <-ctx.Done():
			resovler.deferFn()
			break Loop
		case <-resovler.ticker.C:
			resovler.buf.SwapBuffer()
			go resovler.recvMessages(resovler.buf.Read())
			continue
		case msg := <-ch:
			resovler.call(msg)
		default:
			continue
		}
	}
}
