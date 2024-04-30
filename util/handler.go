package util

import (
	"context"
	"sync"
	"time"

	"github.com/ainokiseki/go-binance/v2"
	"github.com/grd/statistics"

	"ainokiseki/binance_rush/common"
)

var CalculateDelayOnce sync.Once

var requestDelayMean, requestDelaySd float64

func GetDelay(ctx context.Context, c *binance.Client) (mean, sd float64) {
	CalculateDelayOnce.Do(func() {
		requestDelayMean, requestDelaySd = CalculateDelay(ctx, c)
	})
	return requestDelayMean, requestDelaySd
}

func CalculateDelay(ctx context.Context, c *binance.Client) (mean, sd float64) {
	resChan := make(chan int64, 60)

	ch := common.GOMultiProcess(60, func() {
		res, _ := c.NewServerTimeService().Do(ctx)
		resChan <- res
	})
	time.Sleep(time.Second * 3)
	t := time.Now().UnixMilli()
	close(ch)
	delay := make([]int64, 60)
	for i := 0; i < 60; i++ {
		res := <-resChan
		delay[i] = res - t
	}
	d := statistics.Int64(delay)
	sd = statistics.Sd(&d)
	mean = statistics.Mean(&d)
	return mean, sd
}
