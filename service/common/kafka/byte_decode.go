package kafka




type ByteDecoder interface {
	Decode() []byte
}