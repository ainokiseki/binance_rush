/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"time"

	"ainokiseki/binance_rush/pkg/client"

	"github.com/ainokiseki/go-binance/v2"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "binance",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}
var c *binance.Client
var f *client.FutureClient

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.binance.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	px := rootCmd.PersistentFlags().StringP("proxy", "p", "", "proxy url if you want to use vpn")
	tm := rootCmd.PersistentFlags().Int64("time", 0, "sleep until this timestamp")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		t := time.Unix(*tm, 0)
		k := time.NewTimer(time.Until(t))
		<-k.C

		c = binance.NewClientWithConfig(binance.ClientCreateConfig{
			Proxy:     *px,
			Signature: nil,
			APIKey:    AK,
			SecretKey: SK,
		})
		f = client.NewFutureClient(binance.ClientCreateConfig{
			Proxy:     *px,
			Signature: nil,
			APIKey:    AK,
			SecretKey: SK,
		})
	}
}
