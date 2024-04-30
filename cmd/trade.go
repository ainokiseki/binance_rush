/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/ainokiseki/go-binance/v2"
	"github.com/spf13/cobra"

	"ainokiseki/binance_rush/common"
	"ainokiseki/binance_rush/util"
)

// tradeCmd represents the trade command
var tradeCmd = &cobra.Command{
	Use:   "trade",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		ctx := context.Background()
		mean, _ := util.GetDelay(ctx, c)
		fmt.Println("平均延时：", mean)
		var startTime int64 = 1712836800000
		fireTime := startTime - int64(mean)
		fmt.Println("点火时间：", fireTime)

		ch := make(chan struct{})
		type orderRes struct {
			resp *binance.CreateOrderResponse
			err  error
		}
		resChan := make(chan orderRes, 120)
		common.GOMultiProcessWithChan(20, func() {
			res3, err := c.NewCreateOrderService().Symbol("TAOUSDT").Side("BUY").
				Type(binance.OrderTypeLimit).
				Quantity("1").
				Price("20").
				TimeInForce(binance.TimeInForceTypeGTC).
				NewOrderRespType(binance.NewOrderRespTypeACK).
				Do(context.Background())
			resChan <- orderRes{res3, err}
		}, ch)

		common.GOMultiProcessWithChan(30, func() {
			res3, err := c.NewCreateOrderService().Symbol("TAOUSDT").Side("BUY").
				Type(binance.OrderTypeLimit).
				Quantity("1").
				Price("30").
				TimeInForce(binance.TimeInForceTypeGTC).
				NewOrderRespType(binance.NewOrderRespTypeACK).
				Do(context.Background())
			resChan <- orderRes{res3, err}
		}, ch)

		ticker := time.NewTicker(time.Millisecond * 2)

		for {
			t := <-ticker.C
			if t.UnixMilli() > fireTime {
				close(ch)
				break
			}
		}

		for i := 0; i < 120; i++ {
			r := <-resChan
			if r.err != nil {
				fmt.Println(r.err)
			} else if r.resp != nil {
				fmt.Println(*r.resp)

			}
		}

	},
}

func init() {
	rootCmd.AddCommand(tradeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tradeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tradeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
