package adapters

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"net/url"
)

type SubscriberResponse struct {
	Emails []string `json:"emails"`
}

type SubscriberAdapter struct {
	subscriberApp *fiber.App
}

func NewSubscriberAdapter(subscriberApp *fiber.App) SubscriberAdapter {
	return SubscriberAdapter{subscriberApp: subscriberApp}
}

func (adapter SubscriberAdapter) GetCurrencyRate(from string, to string) (float32, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.Header.SetMethod(fasthttp.MethodGet)
	req.SetRequestURI(fmt.Sprintf("/", url.QueryEscape(from), url.QueryEscape(to)))

	ctx := &fasthttp.RequestCtx{}
	req.CopyTo(&ctx.Request)

	adapter.subscriberApp.Handler()(ctx)

	response := string(ctx.Response.Body())

	var res SubscriberResponse
	if err := json.Unmarshal([]byte(response), &res); err != nil {
	}
	return res.Rate, nil
}
