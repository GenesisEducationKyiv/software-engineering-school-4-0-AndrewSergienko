package createcustomer

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go_service/internal/rateservice/customers/adapters"
	"go_service/internal/rateservice/customers/presentation"
	"go_service/internal/rateservice/customers/services"
	"go_service/internal/rateservice/customers/services/createcustomer"
	"log/slog"
	"net/mail"
)

type EventGateway interface {
	GetMessages(transactionID string, batchSize int) []adapters.Message
	Emit(name string, data map[string]interface{}, transactionID *string) error
}

type Handler struct {
	container    presentation.InteractorFactory
	eventGateway EventGateway
}

func New(container presentation.InteractorFactory, eventGateway EventGateway) *Handler {
	return &Handler{container, eventGateway}
}

func (h *Handler) HandleRequest(c *fiber.Ctx) error {
	var requestData struct {
		Email string `json:"email"`
	}

	if err := c.BodyParser(&requestData); err != nil {
		slog.Warn(fmt.Sprintf("Cannot parse JSON: %v", err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if _, err := mail.ParseAddress(requestData.Email); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Email"})
	}

	transactionID := uuid.New().String()
	inputData := createcustomer.InputData{Email: requestData.Email, TransactionID: &transactionID, IsRollback: false}
	result := h.container.CreateCustomer().Handle(inputData)

	if result.Err != nil {
		if errors.Is(result.Err, &services.EmailConflictError{}) {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Email already exists"})
		}
		return fiber.ErrInternalServerError
	}

	messages := h.eventGateway.GetMessages(transactionID, 2)
	for _, msg := range messages {
		switch msg.Title {
		case "SubscriberCreated":
			return c.SendStatus(fiber.StatusOK)
		case "SubscriberCreatedError":
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	}
	// TODO: Add handling of emit error
	_ = h.eventGateway.Emit(
		"SubscriberCreatedTimeout",
		map[string]interface{}{"email": requestData.Email},
		&transactionID,
	)
	return c.SendStatus(fiber.StatusInternalServerError)
}
