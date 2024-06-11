package trade

import (
	"log"

	"github.com/ainokiseki/go-binance/v2"

	"ainokiseki/binance_rush/api"
)

type OrderInfo struct {
	symbol           string
	sideType         binance.SideType
	price            string
	quantity         string
	orderType        binance.OrderType
	timeInForce      binance.TimeInForceType
	newOrderRespType binance.NewOrderRespType
}

func NewLimitOrder(price string, quantity string, symbol string) *OrderInfo {
	return &OrderInfo{
		price:     price,
		quantity:  quantity,
		orderType: binance.OrderTypeLimit,
		symbol:    symbol,
	}
}

func NewTrade() *OrderInfo {
	return &OrderInfo{}
}

func (t *OrderInfo) SetSideType(sideType binance.SideType) *OrderInfo {
	t.sideType = sideType
	return t
}

func (t *OrderInfo) SetTimeInForce(timeInForce api.TimeInForceType) *OrderInfo {
	switch timeInForce {
	case api.TimeInForceType_GTC:
		t.timeInForce = binance.TimeInForceTypeGTC
	case api.TimeInForceType_IOC:
		t.timeInForce = binance.TimeInForceTypeIOC
	case api.TimeInForceType_FOK:
		t.timeInForce = binance.TimeInForceTypeFOK
	default:
		log.Fatalf("unsupported time in force type: %v", timeInForce)
	}
	return t
}

func (t *OrderInfo) SetNewOrderRespType(newOrderRespType binance.NewOrderRespType) *OrderInfo {
	t.newOrderRespType = newOrderRespType
	return t
}
