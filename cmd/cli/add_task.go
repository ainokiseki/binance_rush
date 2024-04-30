/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/cobra"

	"ainokiseki/binance_rush/api"
	"ainokiseki/binance_rush/pkg/client"
)

var startTimeMilli *int64
var runTimes *int

// AddTaskCmd represents the cli command
var AddTaskCmd = &cobra.Command{
	Use:   "add [symbol] [quantity] [price] [execute_times]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		GRPCClient := client.NewGRPCClient()
		res, err := GRPCClient.CreateCoinRushTask(ctx, &api.CreateCoinRushTaskRequest{
			StartTimestampMilli: *startTimeMilli,
			Symbol:              args[0],
			BidQuantity:         args[1],
			Price:               args[2],
			ExecuteTimes:        int32(*runTimes),
		})
		if err != nil {
			log.Print(err)
		}
		log.Println(res, err)
	},
}

func init() {
	startTimeMilli = AddTaskCmd.Flags().Int64("start", 0, "task start time")
	// Here you will define your flags and configuration settings.
	runTimes = AddTaskCmd.Flags().Int("times", 0, "task start time")

	AddTaskCmd.MarkFlagRequired("start")
	AddTaskCmd.MarkFlagRequired("times")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cliCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cliCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
