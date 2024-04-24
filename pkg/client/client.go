package client

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ainokiseki/go-binance/v2"
)

type BinanceClient struct {
	*binance.Client
}

func NewClient(config binance.ClientCreateConfig) *BinanceClient {
	return &BinanceClient{
		binance.NewClientWithConfig(config),
	}
}

func (c *BinanceClient) GetStableDepth(ctx context.Context, symbol string) (float64, float64) {

	res, err := c.NewDepthService().Symbol(symbol + USDT).Limit(1).Do(ctx)
	if err != nil {
		fmt.Printf("get stable depth error：%s", err.Error())
		return 0, 0
	}
	buy, _ := strconv.ParseFloat(res.Bids[0].Price, 64)
	sell, _ := strconv.ParseFloat(res.Asks[0].Price, 64)

	return buy, sell
}
func (c *BinanceClient) getStableDepth(ctx context.Context, symbol, usdt string) (float64, float64) {

	res, err := c.NewDepthService().Symbol(symbol + usdt).Limit(1).Do(ctx)
	if err != nil {
		fmt.Printf("get stable depth error：%s", err.Error())
		return 0, 0
	}
	buy, _ := strconv.ParseFloat(res.Bids[0].Price, 64)
	sell, _ := strconv.ParseFloat(res.Asks[0].Price, 64)

	return buy, sell
}
func (c *BinanceClient) GetChaoDepth(ctx context.Context, chao, stable string) (*binance.DepthResponse, error) {
	return c.NewDepthService().Symbol(chao + stable).Limit(2).Do(ctx)
}

func (c *BinanceClient) LimitTakerBuyOrder(ctx context.Context, chao, stable, price, quantity string) error {
	_, err := c.NewCreateOrderService().Symbol(chao + stable).Side("BUY").
		Type(binance.OrderTypeLimit).
		Quantity(quantity).
		Price(price).
		TimeInForce(binance.TimeInForceTypeFOK).
		NewOrderRespType(binance.NewOrderRespTypeACK).
		Do(context.Background())
	return err
}
func (c *BinanceClient) MarketSellOrder(ctx context.Context, chao, stable, price, quantity string) error {
	_, err := c.NewCreateOrderService().Symbol(chao + stable).Side("SELL").
		Type(binance.OrderTypeMarket).
		Quantity(quantity).
		NewOrderRespType(binance.NewOrderRespTypeACK).
		Do(context.Background())
	return err
}
func (c *BinanceClient) MarketSellStableOrder(ctx context.Context, stable, quantity string, side binance.SideType) error {
	_, err := c.NewCreateOrderService().Symbol(stable + USDT).Side(side).
		Type(binance.OrderTypeMarket).
		Quantity(quantity).
		NewOrderRespType(binance.NewOrderRespTypeACK).
		Do(context.Background())
	return err
}
