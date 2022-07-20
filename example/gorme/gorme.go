package gorme

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	"time"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Birthday  sql.NullTime `gorm:""`
	UpdatedAt time.Time    `gorm:"autoUpdateTime"`
	CreatedAt time.Time    `gorm:"autoCreateTime"`
}

type Task struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	Note      string
	Status    string
	UpdatedAt time.Time
	CreatedAt time.Time
}

func models() []interface{} {
	return []interface{}{
		&User{},
		&Task{},
	}
}

func makeConnection() (*gorm.DB, error) {
	dsn := "host=localhost user=user password=pass dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, err
}

func writeGeneratedDDL(statements []string) {
	now := time.Now()
	basePath := "./example/gorme/migrations"
	upFilePath := fmt.Sprintf("%s/%s_changes.up.sql", basePath, now.Format("20060201150415"))
	downFilePath := fmt.Sprintf("%s/%s_changes.down.sql", basePath, now.Format("20060201150415"))

	var byteString []byte
	for _, s := range statements {
		ddl := fmt.Sprintf("%s;\n", s)
		byteString = append(byteString, []byte(ddl)...)
		log.Println(ddl)
	}

	if err := ioutil.WriteFile(upFilePath, byteString, 0644); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(downFilePath, byteString, 0644); err != nil {
		log.Fatal(err)
	}
}

func GenerateMigration() {
	db, err := makeConnection()
	if err != nil {
		log.Fatal(err)
	}

	tx := db.Begin()
	var statements []string
	tx.Callback().Raw().Register("record_migration", func(tx *gorm.DB) {
		statements = append(statements, tx.Statement.SQL.String())
	})
	if err := tx.AutoMigrate(models()...); err != nil {
		panic(err)
	}
	tx.Rollback()
	tx.Callback().Raw().Remove("record_migration")

	writeGeneratedDDL(statements)
}

func Execute() {
	db, err := makeConnection()
	if err != nil {
		log.Fatal(err)
	}
	db.Create(&User{FirstName: "Sample", LastName: "User"})

	// var user User
	// db.First(&user, 1)
	// log.Println(user)
}
