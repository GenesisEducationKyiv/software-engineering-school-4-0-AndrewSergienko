package deletesubscriber

type InputData struct {
	Email string
}

type OutputData struct {
	Err error
}

type SubscriberGateway interface {
	Delete(email string) error
}

type DeleteSubscriber struct {
	subscriberGateway SubscriberGateway
}

func NewDeleteSubscriber(sg SubscriberGateway) *DeleteSubscriber {
	return &DeleteSubscriber{subscriberGateway: sg}
}

func (s *DeleteSubscriber) Handle(data InputData) OutputData {
	return OutputData{s.subscriberGateway.Delete(data.Email)}
}
