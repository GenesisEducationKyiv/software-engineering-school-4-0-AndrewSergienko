package api

import (
	"encoding/json"
	"go_service/src"
	"net/http"
)

type SubscriberGateway interface {
	src.SubscriberWriter
	src.SubscriberDeleter
	src.SubscriberReader
}

func GetSubscribersHandler(sg SubscriberGateway) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var requestData struct {
			Email string `json:"email"`
		}

		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if requestData.Email == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		subscriber := sg.GetByEmail(requestData.Email)
		if subscriber != nil {
			w.WriteHeader(http.StatusConflict)
			return
		}

		err = sg.Create(requestData.Email)
		if err != nil {
			w.WriteHeader(http.StatusAccepted)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}
}
