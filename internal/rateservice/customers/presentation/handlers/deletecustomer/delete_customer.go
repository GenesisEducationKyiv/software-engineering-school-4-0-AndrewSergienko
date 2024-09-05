package deletecustomer

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go_service/internal/rateservice/customers/adapters"
	"go_service/internal/rateservice/customers/presentation"
	"go_service/internal/rateservice/customers/services/deletecustomer"
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

// HandleRequest @Summary Delete a customer
// @Description Delete a customer with the provided email
// @Tags customer
// @Accept json
// @Produce json
// @Param email body string true "Customer email"
// @Success 200 {string} string "OK"
// @Failure 400 {object} map[string]string "error"
// @Failure 500 {object} map[string]string "error"
// @Router /customer [delete]
func (h *Handler) HandleRequest(c *fiber.Ctx) error {
	var requestData struct {
		Email string `json:"email"`
	}

	if err := c.BodyParser(&requestData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	transactionID := uuid.New().String()
	inputData := deletecustomer.InputData{Email: requestData.Email, TransactionID: &transactionID, IsRollback: false}
	result := h.container.DeleteCustomer().Handle(inputData)

	if result.Err != nil {
		return fiber.ErrInternalServerError
	}

	messages := h.eventGateway.GetMessages(transactionID, 2)
	for _, msg := range messages {
		switch msg.Title {
		case "SubscriberDeleted":
			return c.SendStatus(fiber.StatusOK)
		case "SubscriberDeletedError":
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	}
	// TODO: Add handling of emit error
	_ = h.eventGateway.Emit(
		"SubscriberDeletedTimeout",
		map[string]interface{}{"email": requestData.Email},
		&transactionID,
	)
	return c.SendStatus(fiber.StatusInternalServerError)
}
