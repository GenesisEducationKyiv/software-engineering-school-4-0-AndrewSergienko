package deletesubscriber

type InputData struct {
	Email         string
	TransactionID *string
}

type OutputData struct {
	Err error
}

type SubscriberGateway interface {
	Delete(email string) error
}

type EventEmitter interface {
	Emit(name string, data map[string]interface{}, transactionID *string) error
}

type DeleteSubscriber struct {
	subscriberGateway SubscriberGateway
	eventEmitter      EventEmitter
}

func New(sg SubscriberGateway, em EventEmitter) *DeleteSubscriber {
	return &DeleteSubscriber{subscriberGateway: sg, eventEmitter: em}
}

func (s *DeleteSubscriber) Handle(data InputData) OutputData {
	err := s.subscriberGateway.Delete(data.Email)

	var event string
	if err == nil {
		event = "SubscriberDeleted"
	} else {
		event = "SubscriberDeletedError"
	}
	_ = s.eventEmitter.Emit(
		event,
		map[string]interface{}{"email": data.Email},
		data.TransactionID,
	)
	// TODO: Add transactional outbox pattern
	return OutputData{err}
}
