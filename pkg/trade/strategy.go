package trade

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/ainokiseki/go-binance/v2"
	"golang.org/x/sync/errgroup"
)

var USDT = "USDT"
var PriceGap = 4
var Precission = 0.01
var ChaoBuyNumStr = "1"
var QuantNum = 1.0
var ChaoFormatPrice = "%.4f"
var ChaoPriceRoundingNum = 0.00005

var ProfitLimit = 10003

var MaxTransTimeLimit = 100

var total int

type tri struct {
	s2c, c2s, s2s float64
}

func InitMaxTransTime(x int) {
	MaxTransTimeLimit = x
}

func InitChaoBuyNumStr(buyNum string) {
	ChaoBuyNumStr = buyNum
	QuantNum = getFloat(buyNum)
}

func InitChaoPriceConfig(precision int) {
	var data = map[int]struct {
		string
		float64
	}{
		1: {
			"%.1f", 0.05,
		},
		2: {
			"%.2f", 0.005,
		},
		3: {
			"%.3f", 0.0005,
		},
		4: {
			"%.4f", 0.00005,
		},
	}
	ChaoFormatPrice = data[precision].string
	ChaoPriceRoundingNum = data[precision].float64

}

func RunTriangle(c *BinanceClient, chao, stable string) {
	depthTicker := time.NewTicker(time.Millisecond * 100)
	defer depthTicker.Stop()
	stableTicker := time.NewTicker(time.Second * 5)
	defer stableTicker.Stop()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var s2c, u2c tri

	go func() {
		for {
			buy, sell := c.GetStableDepth(ctx, stable)
			s2c.s2s = 1 / sell
			u2c.s2s = buy

			select {
			case <-stableTicker.C:
			case <-ctx.Done():
			}
		}
	}()

	for {
		g, ctx2 := errgroup.WithContext(ctx)
		select {
		case <-depthTicker.C:
			if total > 1 {
				return
			}
			g.Go(func() error {
				res, err := c.GetChaoDepth(ctx2, chao, stable)
				if err != nil {
					fmt.Println("get price error", chao, stable, err)

					return err
				}
				s2c.s2c = 1 / getFloat(res.Asks[0].Price)
				u2c.c2s = getFloat(res.Bids[0].Price)
				return nil
			})
			g.Go(func() error {
				res, err := c.GetChaoDepth(ctx2, chao, USDT)
				if err != nil {
					fmt.Println("get price error", chao, USDT, err)
					return err
				}
				s2c.c2s = getFloat(res.Bids[0].Price)
				u2c.s2c = 1 / getFloat(res.Asks[0].Price)
				return nil
			})
			err := g.Wait()
			if err != nil {
				fmt.Println("get price error", err)
				return
			}
			go calcTriangle(c, s2c, "s2c", stable, chao)

			go calcTriangle(c, u2c, "u2c", stable, chao)

		}

	}
}

func calcTriangle(c *BinanceClient, price tri, name, stable, chao string) {
	if total >= MaxTransTimeLimit {
		return
	}

	start := 10000.0
	cc := start * price.s2c * price.c2s * price.s2s * 0.99925

	if cc > float64(ProfitLimit) {
		fmt.Println("do strategy", name, cc, fmt.Sprint("%+v", price))
		if name == "s2c" {

			err := c.LimitTakerBuyOrder(context.Background(), chao, stable, fmt.Sprintf(ChaoFormatPrice, 1/price.s2c+ChaoPriceRoundingNum), ChaoBuyNumStr)
			if err == nil {
				c.MarketSellOrder(context.Background(), chao, USDT, fmt.Sprintf(ChaoFormatPrice, price.c2s+ChaoPriceRoundingNum), ChaoBuyNumStr)
				total++
				if err == nil {
					err := c.MarketSellStableOrder(context.Background(), stable, fmt.Sprintf("%d", int(1/price.c2s*QuantNum+0.5)), binance.SideTypeBuy)
					fmt.Println("sell fdusd", err)
				}
			} else {
				fmt.Println("limit buy error:", err)
			}
		} else {
			err := c.LimitTakerBuyOrder(context.Background(), chao, USDT, fmt.Sprintf(ChaoFormatPrice, 1/price.s2c+ChaoPriceRoundingNum), ChaoBuyNumStr)
			if err == nil {
				err = c.MarketSellOrder(context.Background(), chao, stable, fmt.Sprintf(ChaoFormatPrice, price.c2s+ChaoPriceRoundingNum), ChaoBuyNumStr)
				fmt.Println("market sell error:", err)
				total++
				if err == nil {

					err := c.MarketSellStableOrder(context.Background(), stable, fmt.Sprintf("%d", int(price.c2s*QuantNum+0.5)), binance.SideTypeSell)
					fmt.Println("sell fdusd", err)

				}
			} else {
				fmt.Println("limit buy error:", err)
			}
		}

	}

}
func getFloat(num string) float64 {
	res, _ := strconv.ParseFloat(num, 64)
	return res
}

func RunFuture(c *BinanceClient, f *FutureClient, chao, stable string) {
	ticker := time.NewTicker(time.Millisecond * 100)
	ctx := context.Background()
	defer ticker.Stop()
	var buy, sell float64
	var fbuy, fsell float64

	for {

		<-ticker.C
		wg := sync.WaitGroup{}
		wg.Add(2)
		go func() {
			defer wg.Done()
			buy, sell = c.getStableDepth(ctx, chao, stable)
		}()
		go func() {
			defer wg.Done()
			fbuy, fsell = f.GetDepth(ctx, chao, stable)
		}()
		wg.Wait()
		calculateSellFuture(buy, fsell, chao, stable)
		calculateBuyFuture(fbuy, sell, chao, stable)
	}
}
func calculateSellFuture(buy, fsell float64, chao, stable string) {
	start := 1000.0
	rate := start / buy * fsell * 0.999 / 1000
	if rate > 1.0003 {
		fmt.Println("do future sell", rate, buy, fsell)
	} else {
		fmt.Println("not do future sell", rate, buy, fsell)

	}
}
func calculateBuyFuture(fbuy, sell float64, chao, stable string) {
	start := 1000.0
	rate := start / fbuy * sell * 0.999 / 1000
	if rate > 1.0003 {
		fmt.Println("do future buy", rate, fbuy, sell)
	} else {
		fmt.Println("not do future buy", rate, fbuy, sell)

	}
}
