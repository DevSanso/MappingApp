package kafka

import (
	"context"
	"log"
	"sync"
	"time"
)

type resolverImpl[T any] struct {
	h Handle[T]
	recv Recv[T]
	sender Sender

	handleOnce sync.Once
	cancel context.CancelFunc
}



func NewResolver[T any](config *ResovlerConfig) (Resovler[T], error) {
	var r Recv[T]
	var s Sender
	var err error

	if r,err = NewRecv[T](config.recvConfig); err != nil {
		return nil,err
	}
	if s,err = NewSender(config.senderConfig);err != nil {
		return nil,err
	}

	instance := &resolverImpl[T]{}
	instance.recv = r
	instance.sender = s
	return instance,nil
}



func (ri *resolverImpl[T])setCancelFunc(cancel context.CancelFunc) {ri.cancel = cancel}
func (ri *resolverImpl[T])freeCancelFunc() {ri.setCancelFunc(nil)}

func (ri *resolverImpl[T])newCancelCtx() context.Context {
	ctx,cancel := context.WithCancel(context.Background())
	ri.setCancelFunc(cancel)
	return ctx
}
func (ri *resolverImpl[T])close() error {
	if err := ri.recv.Close();err != nil {return err}

	return ri.sender.Close()
}

func (ri *resolverImpl[T])loop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			err := ri.close()
			if err != nil {log.Panicln(err)}
		default:
			data,err := ri.recv.Recv(time.Millisecond)
			if err != ErrTimeoutGetMessage {log.Println(err);continue}
			go ri.h(ri.sender.Send,data)
		}
	}
}



func (ri *resolverImpl[T])SetHandle(handle Handle[T]) error {
	ri.handleOnce.Do(func() {
		ri.h = handle
	})
	return nil
}


func (ri *resolverImpl[T])Start() error {
	
	go ri.loop(ri.newCancelCtx())
	return nil
}
func (ri *resolverImpl[T])Stop() error {
	ri.cancel()
	return nil
}
