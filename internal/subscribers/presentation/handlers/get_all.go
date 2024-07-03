package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"go_service/internal/subscribers/presentation"
	"go_service/internal/subscribers/services/get_all"
)

type GetAll interface {
	Handle(data get_all.InputDTO) get_all.OutputDTO
}

type GetAllHandler struct {
	container presentation.InteractorFactory
}

func NewGetAllHandler(container presentation.InteractorFactory) *GetAllHandler {
	return &GetAllHandler{container: container}
}

func (h *GetAllHandler) HandleRequest(c *fiber.Ctx) error {
	interactor := h.container.GetAll()
	result := interactor.Handle(get_all.InputDTO{})

	data, err := json.Marshal(result.Subscribers)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.Send(data)
}
