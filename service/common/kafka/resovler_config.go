package kafka


type ResovlerConfig struct {
	SenderConfig *SenderConfig `json:"senderConfig"`
	RecvConfig *RecvConfig `json:"recvConfig"`
}