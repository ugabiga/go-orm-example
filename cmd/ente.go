package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ugabiga/go-orm-example/example/ente"
	"github.com/ugabiga/go-orm-example/internal"
)

func init() {
	enteCmd.AddCommand(enteMigrateUpCmd)
	enteCmd.AddCommand(enteMigrateDownCmd)
	enteCmd.AddCommand(enteRunCmd)
	enteCmd.AddCommand(enteGenerateMigration)
}

var enteMigrationPath = "file://example/boiler/migrations"

var enteCmd = &cobra.Command{
	Use:   "en",
	Short: "entgo example",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sqlboiler example")
	},
}

var enteGenerateMigration = &cobra.Command{
	Use:   "gen",
	Short: "run migrate up",
	Run: func(cmd *cobra.Command, args []string) {
		ente.GenerateMigration()
	},
}

var enteMigrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "run migrate up",
	Run: func(cmd *cobra.Command, args []string) {
		internal.UpMigration(enteMigrationPath)
	},
}

var enteMigrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "run migrate down",
	Run: func(cmd *cobra.Command, args []string) {
		internal.DownMigration(migrationPath)
	},
}

var enteRunCmd = &cobra.Command{
	Use:   "run",
	Short: "run sqlboiler example",
	Run: func(cmd *cobra.Command, args []string) {
		ente.Execute()
	},
}
