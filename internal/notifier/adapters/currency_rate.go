package adapters

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"net/url"
)

type Response struct {
	Rate float32 `json:"rate"`
}

type CurrencyRateAdapter struct {
	currencyApp *fiber.App
}

func NewCurrencyRateAdapter(currencyApp *fiber.App) CurrencyRateAdapter {
	return CurrencyRateAdapter{currencyApp: currencyApp}
}

func (adapter CurrencyRateAdapter) GetCurrencyRate(from string, to string) (float32, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.Header.SetMethod(fasthttp.MethodGet)
	req.SetRequestURI(fmt.Sprintf("/?from=%s&to=%s", url.QueryEscape(from), url.QueryEscape(to)))

	ctx := &fasthttp.RequestCtx{}
	req.CopyTo(&ctx.Request)

	adapter.currencyApp.Handler()(ctx)

	response := string(ctx.Response.Body())

	var res Response
	if err := json.Unmarshal([]byte(response), &res); err != nil {
	}
	return res.Rate, nil
}
