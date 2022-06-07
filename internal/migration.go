package internal

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/ugabiga/go-orm-example/config"
	"log"
)

func makeMigrate(migrationPath string) *migrate.Migrate {
	db, err := sql.Open(config.GetDBInfo())
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		migrationPath,
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	return m
}

func UpMigration(migrationPath string) {
	log.Println("Up Migration")
	m := makeMigrate(migrationPath)
	err := m.Up()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Complete Up Migration")
}

func DownMigration(migrationPath string) {
	log.Println("Up Migration")
	m := makeMigrate(migrationPath)
	err := m.Down()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Complete Up Migration")
}
