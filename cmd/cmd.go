package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	runCmd.AddCommand(boilerCmd)
	runCmd.AddCommand(enteCmd)
	runCmd.AddCommand(gormeCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run app",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("app")
	},
}

func Execute() {
	err := runCmd.Execute()
	if err != nil {
		panic(err)
	}
}
