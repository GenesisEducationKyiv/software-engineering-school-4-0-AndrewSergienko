package get_all

import (
	"go_service/internal/infrastructure/database/models"
)

type InputDTO struct{}

type OutputDTO struct {
	subscribers []models.Subscriber
}

type SubscriberGateway interface {
	GetAll() []models.Subscriber
}

type GetAllHandler struct {
	gateway SubscriberGateway
}

func NewGetAllHandler(gateway SubscriberGateway) *GetAllHandler {
	return &GetAllHandler{gateway: gateway}
}

func (s *GetAllHandler) Handle(data InputDTO) OutputDTO {
	return OutputDTO{s.gateway.GetAll()}
}
