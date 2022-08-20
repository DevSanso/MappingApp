package kafka


type ResovlerConfig struct {
	SenderConfig *SenderConfig `json:"senderConfig"`
	RecvConfig *RecvConfig `json:"recvConfig"`
}

func NewResolverConfig() *ResovlerConfig {
	return &ResovlerConfig{
		SenderConfig: &SenderConfig{},
		RecvConfig: &RecvConfig{},
	}
}