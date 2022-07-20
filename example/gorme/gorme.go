package gorme

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	"time"

	"github.com/ugabiga/go-orm-example/config"
	"github.com/ugabiga/go-orm-example/internal"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func GenerateMigration(basePath string) {
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

	writeGeneratedDDL(basePath, statements)
}

func Execute() {
	ctx := context.Background()
	db, err := makeConnection()
	if err != nil {
		log.Fatal(err)
	}

	clear(ctx, db)
	crud(ctx, db)
	queryWithRelation(ctx, db)
}

func queryWithRelation(ctx context.Context, db *gorm.DB) {
	tx := db.WithContext(ctx)
	// Create with relation
	for i := 0; i < 10; i++ {
		newUser := User{
			FirstName: "John",
			LastName:  "Doe",
			Birthday:  sql.NullTime{Time: time.Now().AddDate(-30, 0, 0), Valid: true},
			Tasks: []Task{
				{Title: "Task 1", Note: "Task 1 note"},
				{Title: "Task 2", Note: "Task 2 note"},
			},
		}
		if err := tx.Create(&newUser).Error; err != nil {
			internal.LogFatal(err)
		}
		internal.PrintJSONLog(newUser)
	}

	// Read with relationship
	var user User
	if err := tx.Model(&User{}).Preload("Tasks").First(&user, &User{FirstName: "John"}).Error; err != nil {
		internal.LogFatal(err)
	}
	internal.PrintJSONLog(user)

	// Eager Loading
	var tasks []Task
	r := tx.Model(&Task{}).Preload("User").Find(&tasks, &Task{Title: "Task 1"})
	if err := r.Error; err != nil {
		internal.LogFatal(err)
	}
	internal.PrintJSONLog(tasks)
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

func clear(ctx context.Context, db *gorm.DB) {
	if err := db.WithContext(ctx).Exec("TRUNCATE TABLE project_tasks, projects, tasks, users").Error; err != nil {
		internal.LogFatal(err)
	}
}

func models() []interface{} {
	return []interface{}{
		&User{},
		&Task{},
		&Project{},
	}
}

func makeConnection() (*gorm.DB, error) {
	sqlDB, err := sql.Open(config.GetDBInfo())
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: sqlDB}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, err
}

func writeGeneratedDDL(basePath string, statements []string) {
	now := time.Now()
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
