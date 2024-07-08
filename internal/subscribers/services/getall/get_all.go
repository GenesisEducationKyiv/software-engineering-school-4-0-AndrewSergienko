package getall

import (
	"go_service/internal/subscribers/infrastructure/database/models"
)

type InputDTO struct{}

type OutputDTO struct {
	Subscribers []models.Subscriber
}

type SubscriberGateway interface {
	GetAll() []models.Subscriber
}

type GetAll struct {
	gateway SubscriberGateway
}

func NewGetAllHandler(gateway SubscriberGateway) *GetAll {
	return &GetAll{gateway: gateway}
}

func (s *GetAll) Handle(data InputDTO) OutputDTO { //nolint:all
	return OutputDTO{s.gateway.GetAll()}
}
