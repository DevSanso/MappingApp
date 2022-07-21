package buffer

import (
	"sync"
	"time"
)




type Buffer[T any] struct {
	wChannel chan T
	rChannel chan T
	ticker time.Ticker

	mtx sync.RWMutex

}

func NewBuffer[T any](tick time.Duration) *Buffer[T] {
	buf := new(Buffer[T])
	buf.ticker = *time.NewTicker(tick)
	go buf.tickerHandle(buf.swapBuffer)

	return buf
}

func (b *Buffer[T])swapBuffer() {
	b.mtx.Lock()
	b.mtx.RLock()
	temp := b.wChannel
	b.wChannel = b.rChannel
	b.rChannel = temp
	b.mtx.Unlock()
	b.mtx.RUnlock()
}

func (b *Buffer[T])tickerHandle(fn func()) {
	for {
		select {
		case <-b.ticker.C:
			fn()
		default:
			continue
		}
	}
}

func (b *Buffer[T])Push(data T) {
	b.mtx.Lock()
	b.wChannel <- data
	b.mtx.Unlock()
}


func (b *Buffer[T])copyRChannelThread(other chan T) {
	for {
		b.mtx.RLock()
		data,ok := <- b.rChannel
		if !ok {panic("buffer copyRChannelThread panic")}
		other <- data
		b.mtx.RUnlock()
	}
}

func (b *Buffer[T])ReadChannel() <- chan T {
	other := make(chan T)
	go b.copyRChannelThread(other)
	return other
}