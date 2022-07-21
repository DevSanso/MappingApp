package kafka



type ByteEncodeAndWriter interface {
	EncodeAndWrite([]byte) error
}