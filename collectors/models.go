package collectors

// Reads from pub//sub source when awake
type Collector interface {
	Sleep()
	Wake()
	GetMessages() ([]MessageWrapper, error)
	PublishMessage(message *MessageWrapper, delaySeconds int64, errFlag bool) error
	AckMessage(receiptHandle string) error
}

type Result struct {
	Err   error
	Retry bool
	Fatal bool
}

type MessageWrapper struct {
	MessageBody   []byte
	ReceiptHandle string
	Retries       int
	Retry         bool
	Fatal         bool
	Message       interface{}
}

type CollectorOptions struct {
	PollingPeriod     int
	MaxPollingPeriod  int
	MaxRetries        int
	RetryIntervalSecs int64
	SourceTopic       string
	ErrorTopic        string
	Region            string
	AccountID         string
	BusinessProcessor func([]byte) *Result
	GetMessages       func() ([]MessageWrapper, error)
	PublishMessage    func(message *MessageWrapper, delaySeconds int64, errFlag bool) error
	AckMessage        func(receiptHandle string) error
	Sleep             func()
	Wake              func()
}
