package handlers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"go_service/internal/customers/presentation"
	"go_service/internal/customers/services/getall"
)

type GetAll interface {
	Handle(data getall.InputDTO) getall.OutputDTO
}

type GetAllHandler struct { // nolint: exported
	container presentation.InteractorFactory
}

func NewGetAllHandler(container presentation.InteractorFactory) *GetAllHandler {
	return &GetAllHandler{container: container}
}

func (h *GetAllHandler) HandleRequest(c *fiber.Ctx) error {
	interactor := h.container.GetAll()
	result := interactor.Handle(getall.InputDTO{})

	data, err := json.Marshal(result.Subscribers)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.Send(data)
}
