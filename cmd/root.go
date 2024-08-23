package cmd

import (
	"os"

	"github.com/joseasousa/stress_test/internal/domain"
	"github.com/joseasousa/stress_test/internal/usecase"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stress_test",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		url := cmd.Flag("url").Value.String()
		concurrency, _ := cmd.Flags().GetInt("concurrency")
		totalRequests, _ := cmd.Flags().GetInt("requests")

		config := domain.Config{
			URL:           url,
			Concurrency:   concurrency,
			TotalRequests: totalRequests,
		}

		uc := usecase.NewStressTest()
		resp := uc.Execute(config)
		resp.PrintResult()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().StringP("url", "u", "", "URL to be tested")
	rootCmd.Flags().IntP("concurrency", "c", 10, "Number of concurrent requests")
	rootCmd.Flags().IntP("requests", "r", 10, "Number of requests")

	_ = rootCmd.MarkFlagRequired("url")
}
