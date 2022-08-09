package kafka

import "google.golang.org/protobuf/proto"

type SendMessage interface {
	Err(err error)
	Ok(topic string,msessage proto.Message)
}

type sendMessageImpl struct {
	err error
	topic string
	message proto.Message
}

func (s *sendMessageImpl)Err(err error) {
	s.err = err
}

func (s *sendMessageImpl)Ok(topic string,message proto.Message) {
	s.message = message
	s.topic = topic
}
