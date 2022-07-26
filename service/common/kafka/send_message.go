package kafka


type SendMessage interface {
	Err(err error)
	Ok(topic string,msessage []byte)
}

type sendMessageImpl struct {
	err error
	topic string
	message []byte
}

func (s *sendMessageImpl)Err(err error) {
	s.err = err
}

func (s *sendMessageImpl)Ok(topic string,message []byte) {
	s.message = message
	s.topic = topic
}
