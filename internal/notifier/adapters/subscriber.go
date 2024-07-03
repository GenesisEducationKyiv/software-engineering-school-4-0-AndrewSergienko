package adapters

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type Subscriber struct {
	Email string `json:"Email"`
}

type SubscriberAdapter struct {
	subscriberApp *fiber.App
}

func NewSubscriberAdapter(subscriberApp *fiber.App) SubscriberAdapter {
	return SubscriberAdapter{subscriberApp: subscriberApp}
}

func (adapter SubscriberAdapter) GetAll() ([]string, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.Header.SetMethod(fasthttp.MethodGet)
	req.SetRequestURI("/")

	ctx := &fasthttp.RequestCtx{}
	req.CopyTo(&ctx.Request)

	adapter.subscriberApp.Handler()(ctx)

	response := string(ctx.Response.Body())

	var res []Subscriber
	if err := json.Unmarshal([]byte(response), &res); err != nil {
		return nil, err
	}
	emails := make([]string, len(res))
	for i, subscriber := range res {
		emails[i] = subscriber.Email
	}
	return emails, nil
}
