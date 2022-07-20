package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ugabiga/go-orm-example/example/gorme"
	"github.com/ugabiga/go-orm-example/internal"
)

func init() {
	gormeCmd.AddCommand(gormeRunCmd)
	gormeCmd.AddCommand(gormeGenerateMigration)
	gormeCmd.AddCommand(gormeMigrateUpCmd)
	gormeCmd.AddCommand(gormeMigrateDownCmd)
}

var migrationPath = "file://example/gorme/migrations"

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

var gormeGenerateMigration = &cobra.Command{
	Use:   "gen",
	Short: "generate gorm migration",
	Run: func(cmd *cobra.Command, args []string) {
		gorme.GenerateMigration()
	},
}

var gormeMigrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "run migrate up",
	Run: func(cmd *cobra.Command, args []string) {
		internal.UpMigration(migrationPath)
	},
}

var gormeMigrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "run migrate down",
	Run: func(cmd *cobra.Command, args []string) {
		internal.DownMigration(migrationPath)
	},
}
