package api

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

//func GetSubscribersHandler(sg SubscriberGateway) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		if r.Method != http.MethodPost {
//			w.WriteHeader(http.StatusMethodNotAllowed)
//			return
//		}
//
//		var requestData struct {
//			Email string `json:"email"`
//		}
//
//		err := json.NewDecoder(r.Body).Decode(&requestData)
//		if err != nil {
//			w.WriteHeader(http.StatusBadRequest)
//			return
//		}
//
//		if requestData.Email == "" {
//			w.WriteHeader(http.StatusBadRequest)
//			return
//		}
//
//		subscriber := sg.GetByEmail(requestData.Email)
//		if subscriber != nil {
//			w.WriteHeader(http.StatusConflict)
//			return
//		}
//
//		err = sg.Create(requestData.Email)
//		if err != nil {
//			w.WriteHeader(http.StatusAccepted)
//			return
//		}
//		w.WriteHeader(http.StatusOK)
//	}
//}

type SubscribersHandlers struct {
	subscriberGateway SubscriberGateway
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
