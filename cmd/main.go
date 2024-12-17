package main

import (
	"fmt"
	"log"
	"time"

	stress "github.com/dpcamargo/fullcycle-stress-test/internal"
	"github.com/spf13/cobra"
)

func main() {
	initialTime := time.Now().UnixMilli()
	defer func() {
		finalTime := time.Now().UnixMilli()
		fmt.Printf("Total time: %.2f seconds", float32(finalTime-initialTime)/1000)
	}()
	var rootCmd = &cobra.Command{
		Use:   "stress-test",
		Short: "FullCycle Stress Test",
		Run: func(cmd *cobra.Command, args []string) {
			url, _ := cmd.Flags().GetString("url")
			requests, _ := cmd.Flags().GetInt("requests")
			concurrency, _ := cmd.Flags().GetInt("concurrency")

			stress.Start(url, requests, concurrency)
		},
	}

	rootCmd.Flags().StringP("url", "u", "", "The URL to process (required)")
	rootCmd.Flags().IntP("requests", "r", 1, "Total number of requests (required)")
	rootCmd.Flags().IntP("concurrency", "c", 1, "Number of concurrent operations (required)")

	rootCmd.MarkFlagRequired("url")
	rootCmd.MarkFlagRequired("requests")
	rootCmd.MarkFlagRequired("concurrency")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
