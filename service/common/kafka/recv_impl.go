package kafka

import (
	"errors"
	"time"

	k "github.com/confluentinc/confluent-kafka-go/kafka"
	"google.golang.org/protobuf/proto"
)

type recvImpl[T any] struct {
	consumer *k.Consumer
}


var (
	ErrTimeoutGetMessage= errors.New("ErrTimeoutGetMessage")
	ErrCantCastGenericValue = errors.New("ErrCantCastGenericValue")
)

func NewRecv[T any](config *RecvConfig) (Recv[T], error) {
	cons,err := k.NewConsumer(config.toConfigMap())
	if err != nil {
		return nil,err
	}
	if err = cons.SubscribeTopics([]string{config.Topic}, nil);err != nil {
		return nil,err
	}


	r := &recvImpl[T]{
		consumer: cons,
	}
	return r,nil
}

func (ri *recvImpl[T])Recv(t time.Duration) (*T,error) {
	origin := new(T)
	casting := interface{}(origin)
	ref,isRef := casting.(proto.Message)
	if !isRef {
		return nil,ErrCantCastGenericValue
	}
	
	msg,err := ri.consumer.ReadMessage(t)
	
	if cast,ok := err.(k.Error); ok {
		if cast.Code() == k.ErrTimedOut {return nil,ErrTimeoutGetMessage}
	}else if err != nil {return nil,err}

	
	err = proto.Unmarshal(msg.Value,ref)
	return origin,err
}

func (ri *recvImpl[T])Close() error {
	return ri.consumer.Close()
}
