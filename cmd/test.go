/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"ainokiseki/binance_rush/handler"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test called")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		t := time.Now().UnixMilli()
		res, err := c.NewServerTimeService().Do(ctx)
		t2 := time.Now().UnixMilli()

		if err == nil {
			fmt.Println("网络延迟测试成功", "请求延时", res-t, "响应延时", t2-res)
		} else {
			fmt.Println("网络延迟测试失败", err)
			return
		}

		_, err = c.NewGetUserAsset().Do(ctx)
		if err == nil {
			fmt.Println("ak/sk测试成功")
		} else {
			fmt.Println(err)
			fmt.Println("ak/sk测试失败")
			return
		}

		mean, variance := handler.CalculateDelay(ctx, c)
		fmt.Println("平均延时：", mean, "标准差：", variance)

	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
