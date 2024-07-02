package presentation

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go_service/internal/subscribers/services"
	"net/mail"
)

type Subscribe interface {
	Handle(data services.SubscribeInputDTO) services.SubscribeOutputDTO
}

type SubscribersHandlers struct {
	container InteractorFactory
}

func NewSubscribersHandlers(container InteractorFactory) *SubscribersHandlers {
	return &SubscribersHandlers{container}
}

func (sh *SubscribersHandlers) AddSubscriber(c *fiber.Ctx) error {
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
	result := interactor.Handle(services.SubscribeInputDTO{Email: requestData.Email})

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
