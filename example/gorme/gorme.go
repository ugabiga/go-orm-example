package gorme

import (
	"database/sql"
	"log"

	"time"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	FirstName string
	LastName  string
	Birthday  sql.NullTime
	UpdatedAt time.Time
	CreatedAt time.Time
}

func Execute() {
	dsn := "host=localhost user=user password=pass dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&User{})

	db.Create(&User{FirstName: "Sample", LastName: "User"})

	var user User
	db.First(&user, 1)
	log.Println(user)
}
