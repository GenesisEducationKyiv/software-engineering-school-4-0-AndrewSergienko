package handlers

import (
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

func (h *GetAllHandler) HandleRequest(c *fiber.Ctx) error {
	interactor := h.container.GetAll()
	result := interactor.Handle(get_all.InputDTO{})

}
