package client

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ainokiseki/go-binance/v2"

	"github.com/ainokiseki/go-binance/v2/futures"
)

type FutureClient struct {
	*futures.Client
}

func NewFutureClient(config binance.ClientCreateConfig) *FutureClient {
	return &FutureClient{
		futures.NewProxiedClient(config.APIKey, config.SecretKey, config.Proxy),
	}
}

func (c *FutureClient) GetDepth(ctx context.Context, chao, stable string) (float64, float64) {
	res, err := c.NewDepthService().Limit(5).Symbol(chao + stable).Do(ctx)
	if err != nil {
		fmt.Printf("get stable depth error：%s", err.Error())
		return 0, 0
	}
	buy, _ := strconv.ParseFloat(res.Bids[0].Price, 64)
	sell, _ := strconv.ParseFloat(res.Asks[0].Price, 64)

	return buy, sell
}
