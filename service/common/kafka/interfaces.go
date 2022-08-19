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

type SendFunc func(key, topic string, idl proto.Message) error
type Handle[T any] func(SendFunc,*T) error



type Resovler[T any] interface {
	SetHandle(Handle[T]) error
	

	Start() error
	Stop() error
}