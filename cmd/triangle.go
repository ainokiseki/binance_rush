/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"ainokiseki/binance_rush/pkg/trade"

	"github.com/spf13/cobra"
)

type triangleConfig struct {
	limit         *int
	chaoPrecision *int
	quant         *string
	profit        *int
}

var triConfig triangleConfig

// triangleCmd represents the triangle command
var triangleCmd = &cobra.Command{
	Use:   "triangle",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("triangle called")

		stable := cmd.Flag("stable").Value.String()
		chao := cmd.Flag("symbol").Value.String()
		bClient := &trade.BinanceClient{
			Client: c,
		}
		trade.InitChaoBuyNumStr(*triConfig.quant)
		trade.InitChaoPriceConfig(*triConfig.chaoPrecision)
		trade.InitMaxTransTime(*triConfig.limit)
		trade.ProfitLimit = *triConfig.profit
		trade.RunTriangle(bClient, chao, stable)

	},
}

func init() {
	rootCmd.AddCommand(triangleCmd)
	triangleCmd.Flags().String("symbol", "", "chao coin")
	triangleCmd.Flags().String("stable", "", "stable coin, must have no fee between usdt")
	triConfig.quant = triangleCmd.Flags().String("quant", "", "how many coin to buy once a time")
	triConfig.limit = triangleCmd.Flags().Int("limit", 100, "after how many times of transaction before exit")
	triConfig.chaoPrecision = triangleCmd.Flags().Int("chao_precision", -1, "how many digit behind dot of price of chao")
	triConfig.profit = triangleCmd.Flags().Int("profit", 10003, "profit")

	usdt := triangleCmd.Flags().String("usdt", "USDT", "coin as usdt")
	trade.USDT = *usdt

	triangleCmd.MarkFlagRequired("symbol")
	triangleCmd.MarkFlagRequired("stable")
	triangleCmd.MarkFlagRequired("chao_precision")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// triangleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// triangleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
