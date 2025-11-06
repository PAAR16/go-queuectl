package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "queuectl",
	Short: "A simple CLI-based background job queue system.",
	Long: `queuectl is a tool to manage background jobs, run workers,
and handle retries with a dead-letter queue.`,
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// func init() {
// 	// Here you will initialize your global flags if any.
// }