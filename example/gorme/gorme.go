package gorme

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	"time"

	"github.com/ugabiga/go-orm-example/internal"
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
		//Should make your own migration down ddl method
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
	ctx := context.Background()
	db, err := makeConnection()
	if err != nil {
		log.Fatal(err)
	}

	crud(ctx, db)

}

func crud(ctx context.Context, db *gorm.DB) {
	tx := db.WithContext(ctx)
	//Create
	newUser := User{
		FirstName: "Sample",
		LastName:  "User",
		Birthday:  sql.NullTime{Time: time.Now().AddDate(-30, 0, 0), Valid: true},
	}

	if err := tx.Create(&newUser).Error; err != nil {
		internal.LogFatal(err)
	}
	internal.PrintJSONLog(newUser)

	// Read
	var user User
	if err := tx.First(&user, &User{FirstName: "Sample"}).Error; err != nil {
		internal.LogFatal(err)
	}
	internal.PrintJSONLog(user)

	// Read List
	var users []User
	if err := tx.Find(&users, &User{FirstName: "Sample"}).Error; err != nil {
		internal.LogFatal(err)
	}
	internal.PrintJSONLog(users)

	// Update
	r := tx.Model(&User{}).Where(&User{LastName: "User"}).Updates(&User{FirstName: "Unknown"})
	if err := r.Error; err != nil {
		internal.LogFatal(err)
	}
	internal.PrintJSONLog(r.RowsAffected)
	internal.PrintJSONLog(fmt.Sprintf("Updated rows : %d", r.RowsAffected))

	// Delete
	r = tx.Where("id > 0").Delete(&User{})
	if err := r.Error; err != nil {
		internal.LogFatal(err)
	}
	internal.PrintJSONLog(fmt.Sprintf("Deleted rows : %d", r.RowsAffected))
}
