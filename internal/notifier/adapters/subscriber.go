package adapters

import (
	"encoding/json"
	"go_service/internal/notifier/infrastructure"
)

type Subscriber struct {
	Email string `json:"Email"`
}

type SubscriberAdapter struct {
	subscriberServiceSettings *infrastructure.SubscriberServiceAPISettings
}

func NewSubscriberAdapter(subscriberServiceSettings *infrastructure.SubscriberServiceAPISettings) SubscriberAdapter {
	return SubscriberAdapter{subscriberServiceSettings: subscriberServiceSettings}
}

func (adapter SubscriberAdapter) GetAll() ([]string, error) {
	url := adapter.subscriberServiceSettings.Host + adapter.subscriberServiceSettings.GetSubscribersURL
	body, err := ReadHTTP(url)

	if err != nil {
		return nil, err
	}

	var response []Subscriber
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	emails := make([]string, len(response))
	for i, subscriber := range response {
		emails[i] = subscriber.Email
	}
	return emails, nil
}
