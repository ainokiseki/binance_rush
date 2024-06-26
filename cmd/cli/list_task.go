/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cli

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"ainokiseki/binance_rush/api"
	"ainokiseki/binance_rush/pkg/client"
)

// ListTaskCmd represents the cli command
var ListTaskCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		GRPCClient := client.NewGRPCClient()
		fmt.Println("try client")
		res, err := GRPCClient.ListTask(context.Background(), &api.ListTaskRequest{})
		if err != nil {
			log.Println(err)
		}
		for _, i := range res.GetTasks() {
			log.Println(i)
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cliCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cliCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
