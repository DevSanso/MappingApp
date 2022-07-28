package buffer_test

import (
	"testing"
	"common/buffer"
	//"sync"
)



func TestBuffer(t *testing.T) {
	var buf = buffer.NewBuffer[string]()
	buf.Push("hello world")
	buf.Push("hello world")
	buf.Push("hello world")
	buf.Push("hello world")
	buf.Push("hello world")

	var read = buf.Read()

	if len(read) != 0 {
		t.Fatal("not read buffer size 0")
	}

	buf.SwapBuffer()

	read = buf.Read()
	if len(read) == 0 {
		t.Fatal("read buffer size 0")
	}

	for _,val := range read {
		t.Log(val)
	}
}
func TestAsyncBuffer(t *testing.T) {

	
	var buf = buffer.NewBuffer[string]()
	
	go func() {
		buf.Push("hello world")
		buf.Push("hello world")
		buf.Push("hello world")
		buf.Push("hello world")
		buf.Push("hello world")

	}()

	buf.Push("hello")
	buf.Push("hello")
	buf.Push("hello")
	buf.Push("hello")
	buf.SwapBuffer()
	buf.Push("hello")


	r := buf.Read()
	
	for _,val := range r {
		t.Log(val)
	}
	t.Log("---------------------")
	buf.SwapBuffer()
	r = buf.Read()
	
	for _,val := range r {
		t.Log(val)
	}
	
}