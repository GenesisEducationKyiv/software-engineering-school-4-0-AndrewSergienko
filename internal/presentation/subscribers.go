package presentation

import (
	"github.com/gofiber/fiber/v2"
	"go_service/internal/common"
	"net/mail"
)

type SubscriberGateway interface {
	common.SubscriberWriter
	common.SubscriberDeleter
	common.SubscriberReader
}

type SubscribersHandlers struct {
	subscriberGateway SubscriberGateway
}

func InitSubscribersHandlers(subscriberGateway SubscriberGateway) *SubscribersHandlers {
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
		return fiber.ErrBadRequest
	}

	if sh.subscriberGateway.GetByEmail(requestData.Email) != nil {
		return fiber.ErrConflict
	}

	if sh.subscriberGateway.Create(requestData.Email) != nil {
		return fiber.ErrInternalServerError
	}
	return c.SendStatus(fiber.StatusOK)
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
