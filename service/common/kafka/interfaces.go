package kafka

import (
	"time"

	"google.golang.org/protobuf/proto"
)

type Sender interface {
	Send(key, topic string, idl proto.Message) error
	Close() error
}


type Recv[T any] interface {
	Recv(time.Duration) (*T,error)
	Close() error
}


type HandleContext[T any] interface {
	Next(string,*T,...interface{}) error
	Error(error,...interface{})	error
	Send(key, topic string, idl proto.Message) error
}
type Handle[T any] func(HandleContext[T],T,...interface{}) error



type Resovler[T any] interface {
	AddHandle(string,Handle[T])
	DelHandle(string)

	Start()
	Stop()
}