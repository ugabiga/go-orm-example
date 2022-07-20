package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ugabiga/go-orm-example/example/gorme"
)

func init() {
	gormeCmd.AddCommand(gormeRunCmd)
}

var gormeCmd = &cobra.Command{
	Use:   "gorm",
	Short: "gorm example",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gorm example")
	},
}

var gormeRunCmd = &cobra.Command{
	Use:   "run",
	Short: "run gorm example",
	Run: func(cmd *cobra.Command, args []string) {
		gorme.Execute()
	},
}
