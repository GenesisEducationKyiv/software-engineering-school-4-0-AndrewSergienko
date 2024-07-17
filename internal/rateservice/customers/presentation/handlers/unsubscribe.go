package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go_service/internal/rateservice/customers/presentation"
	"go_service/internal/rateservice/customers/services/unsubscribe"
)

type UnsuscribeHandler struct {
	container presentation.InteractorFactory
}

func NewUnsubscriberHandler(container presentation.InteractorFactory) *UnsuscribeHandler {
	return &UnsuscribeHandler{container}
}

func (uh *UnsuscribeHandler) HandleRequest(c *fiber.Ctx) error {
	var requestData struct {
		Email string `json:"email"`
	}

	if err := c.BodyParser(&requestData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	interactor := uh.container.Unsubscribe()
	result := interactor.Handle(unsubscribe.InputDTO{Email: requestData.Email})

	if result.Err != nil {
		return fiber.ErrInternalServerError
	}
	return c.SendStatus(fiber.StatusOK)
}
