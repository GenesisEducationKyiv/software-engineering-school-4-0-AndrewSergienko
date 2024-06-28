package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go_service/internal/services"
	"net/mail"
)

type SubscribersHandlers struct {
	subscriberGateway SubscriberGateway
}

type Interactor[T, U any] interface {
	Handle(data T) (U, error)
}

func NewSubscribersHandlers(subscriberGateway SubscriberGateway) *SubscribersHandlers {
	return &SubscribersHandlers{subscriberGateway}
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
		return nil
	}

	var interactor Interactor[string, error] = services.NewSubscribe()

	return c.SendStatus(fiber.StatusOK)
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
