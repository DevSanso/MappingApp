package buffer

import (
	"sync"
)


type Buffer[T any] struct {
	writeBuffer []T
	readBuffer []T

	mtx sync.RWMutex

}

func NewBuffer[T any]() *Buffer[T] {
	buf := new(Buffer[T])
	return buf
}

func (b *Buffer[T])SwapBuffer() {
	b.mtx.Lock()
	b.mtx.RLock()
	defer b.mtx.Unlock()
	defer b.mtx.RUnlock()

	temp := b.writeBuffer
	b.writeBuffer = b.readBuffer
	b.readBuffer = temp
}



func (b *Buffer[T])Push(data T) {
	b.mtx.Lock()
	defer b.mtx.Unlock()

	b.writeBuffer = append(b.writeBuffer,data)
}


func (b *Buffer[T])Read() []T {
	b.mtx.RLock()
	defer b.mtx.RUnlock()

	var res []T
	copy(res,b.readBuffer)
	b.readBuffer = b.readBuffer[:0]	
	return res
}

