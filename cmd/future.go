/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"ainokiseki/binance_rush/pkg/client"
	"fmt"

	"github.com/spf13/cobra"
)

// futureCmd represents the future command
var futureCmd = &cobra.Command{
	Use:   "future",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("future called")
		stable := cmd.Flag("stable").Value.String()
		chao := cmd.Flag("symbol").Value.String()
		bClient := &client.BinanceClient{
			Client: c,
		}
		client.RunFuture(bClient, f, chao, stable)
	},
}

func init() {
	rootCmd.AddCommand(futureCmd)

	futureCmd.Flags().String("symbol", "", "chao coin")
	futureCmd.Flags().String("stable", "", "stable coin, must have no fee between usdt")

	futureCmd.MarkFlagRequired("symbol")
	futureCmd.MarkFlagRequired("stable")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// futureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// futureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
