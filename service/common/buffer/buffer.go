package buffer

import (
	"sync"
	"sync/atomic"
)


const (
	_SwapOn int32 = 0
	_SwapOff int32 = 1
)

type Buffer[T any] struct {
	writeBuffer []T
	readBuffer []T



	cond *sync.Cond
	swapFlag int32
}

func NewBuffer[T any]() *Buffer[T] {
	buf := new(Buffer[T])
	buf.cond = sync.NewCond(&sync.Mutex{})
	buf.swapFlag = _SwapOff
	return buf
}

func (b *Buffer[T])SwapBuffer() {
	b.cond.L.Lock()
	atomic.StoreInt32(&b.swapFlag,_SwapOn)
	defer b.cond.L.Unlock()
	
	temp := b.writeBuffer
	b.writeBuffer = b.readBuffer
	b.readBuffer = temp
	b.cond.Broadcast()
	atomic.StoreInt32(&b.swapFlag,_SwapOff)
}



func (b *Buffer[T])Push(data T) {
	b.cond.L.Lock()
	if atomic.LoadInt32(&b.swapFlag) == _SwapOn {
		b.cond.Wait()
	}
	b.writeBuffer = append(b.writeBuffer,data)
	b.cond.L.Unlock()
}


func (b *Buffer[T])Read() []T {
	b.cond.L.Lock()
	if atomic.LoadInt32(&b.swapFlag) == _SwapOn {
		b.cond.Wait()
	}
	var res []T = make([]T,len(b.readBuffer))
	copy(res,b.readBuffer)
	b.cond.L.Unlock()
	return res
	
}

