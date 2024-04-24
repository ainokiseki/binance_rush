package cmd

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ainokiseki/go-binance/v2"
)

func TestBi(t *testing.T) {

	c = binance.NewClientWithConfig(binance.ClientCreateConfig{
		Proxy:     "http://127.0.0.1:7890",
		Signature: nil,
		APIKey:    AK,
		SecretKey: SK,
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// res, _ := c.NewDepthService().Limit(4).Symbol("APTFDUSD").Do(ctx)
	// res, _ := c.NewExchangeInfoService().Symbol("ETHFDUSD").Do(ctx)
	// res, err := c.NewTradeFeeService().Symbol("APTUSDT").Do(ctx)
	res, err := c.NewCreateOrderService().Symbol("ETHFDUSD").Side("BUY").
		Type(binance.OrderTypeLimit).
		Quantity("0.01").
		Price("3000").
		TimeInForce(binance.TimeInForceTypeFOK).
		NewOrderRespType(binance.NewOrderRespTypeACK).
		Do(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", res)

}
