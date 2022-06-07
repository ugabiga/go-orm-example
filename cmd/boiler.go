package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ugabiga/go-orm-example/internal"
)

func init() {
	boilerCmd.AddCommand(boilerMigrateUpCmd)
	boilerCmd.AddCommand(boilerMigrateDownCmd)
	boilerCmd.AddCommand(boilerSeedCmd)
	boilerCmd.AddCommand(boilerRunCmd)
}

var migrationPath = "file://examples/boiler/migrations"

var boilerCmd = &cobra.Command{
	Use:   "bo",
	Short: "sqlboiler example",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sqlboiler example")
	},
}

var boilerMigrateUpCmd = &cobra.Command{
	Use:   "up",
	Short: "run migrate up",
	Run: func(cmd *cobra.Command, args []string) {
		internal.UpMigration(migrationPath)
	},
}

var boilerMigrateDownCmd = &cobra.Command{
	Use:   "down",
	Short: "run migrate down",
	Run: func(cmd *cobra.Command, args []string) {
		internal.DownMigration(migrationPath)
	},
}

var boilerSeedCmd = &cobra.Command{
	Use:   "seed",
	Short: "seed data",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("app")
	},
}

var boilerRunCmd = &cobra.Command{
	Use:   "run",
	Short: "run sqlboiler example",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("app")
	},
}
