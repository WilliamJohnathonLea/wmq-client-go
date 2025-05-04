package model

type AssignConsumerCommand struct {
	Type string `json:"type"` // "AssignConsumer"
	ID   string `json:"id"`
}

func NewAssignConsumer(id string) AssignConsumerCommand {
	return AssignConsumerCommand{
		Type: "AssignConsumer",
		ID:   id,
	}
}

type StartConsumerCommand struct {
	Type string `json:"type"` // "StartConsumer"
	ID   string `json:"id"`
}

func NewStartConsumer(id string) StartConsumerCommand {
	return StartConsumerCommand{
		Type: "StartConsumer",
		ID:   id,
	}
}

type AssignProducerCommand struct {
	Type string `json:"type"` // "AssignProducer"
	ID   string `json:"id"`
}

func NewAssignProducer(id string) AssignProducerCommand {
	return AssignProducerCommand{
		Type: "AssignProducer",
		ID:   id,
	}
}

type DeclareQueueCommand struct {
	Type string `json:"type"` // "DeclareQueue"
	Name string `json:"name"`
	Size int    `json:"size"`
}

func NewDeclareQueue(name string, size int) DeclareQueueCommand {
	return DeclareQueueCommand{
		Type: "DeclareQueue",
		Name: name,
		Size: size,
	}
}

type AssignQueueCommand struct {
	Type       string `json:"type"` // "AssignQueue"
	ConsumerID string `json:"consumer_id"`
	Queue      string `json:"queue"`
}

func NewAssignQueue(consumerID, queue string) AssignQueueCommand {
	return AssignQueueCommand{
		Type:       "AssignQueue",
		ConsumerID: consumerID,
		Queue:      queue,
	}
}

type SendMessageCommand struct {
	Type       string `json:"type"` // "SendMessage"
	Queue      string `json:"queue"`
	ProducerID string `json:"producer_id"`
	Msg        string `json:"msg"` // base64-encoded
}

func NewSendMessage(queue, producerID, msg string) SendMessageCommand {
	return SendMessageCommand{
		Type:       "SendMessage",
		Queue:      queue,
		ProducerID: producerID,
		Msg:        msg,
	}
}
