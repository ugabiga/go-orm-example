package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	sqcCmd.AddCommand(sqcRunCmd)
}

var sqcCmd = &cobra.Command{
	Use:   "sc",
	Short: "sqlc example",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("sqlc example")
	},
}

var sqcRunCmd = &cobra.Command{
	Use:   "run",
	Short: "run sqlc example",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("run sqlc example")
	},
}
