package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go_service/internal/rateservice/customers/presentation"
	"go_service/internal/rateservice/customers/services"
	"go_service/internal/rateservice/customers/services/subscribe"
	"net/mail"
)

type Subscribe interface {
	Handle(data subscribe.InputDTO) subscribe.OutputDTO
}

type SubscribeHandler struct {
	container presentation.InteractorFactory
}

func NewSubscriberHandler(container presentation.InteractorFactory) *SubscribeHandler {
	return &SubscribeHandler{container}
}

func (sh *SubscribeHandler) HandleRequest(c *fiber.Ctx) error {
	var requestData struct {
		Email string `json:"email"`
	}

	if err := c.BodyParser(&requestData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if !isValidEmail(requestData.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Email",
		})
	}

	interactor := sh.container.Subscribe()
	result := interactor.Handle(subscribe.InputDTO{Email: requestData.Email})

	if result.Err != nil {
		if errors.Is(result.Err, &services.EmailConflictError{}) {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": "Email already exists",
			})
		}
		return fiber.ErrInternalServerError
	}
	return c.SendStatus(fiber.StatusOK)
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
