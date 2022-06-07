package boiler

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ugabiga/go-orm-example/config"
	"github.com/ugabiga/go-orm-example/examples/boiler/models"
	"github.com/ugabiga/go-orm-example/internal"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"log"
	"math/rand"
	"time"
)

func Execute() {
	ctx := context.Background()

	// Connection
	conn, err := sql.Open(config.GetDBInfo())
	internal.LogFatal(err)

	crud(ctx, conn)
	queryWithRelation(ctx, conn)
	seed(ctx, conn, 100)
	aggregate(ctx, conn)
	pagination(ctx, conn)

	// Transform
	// Raw Query
	// Event
}

func pagination(ctx context.Context, conn *sql.DB) {
	// Limit & Offset
	for i := 0; i < 5; i++ {
		users, err := models.Users(
			qm.Limit(3),
			qm.Offset(i*3),
		).All(ctx, conn)
		internal.LogFatal(err)
		internal.PrintJSONLog(users)
	}

	// Cursor
	lastUserID := int64(0)
	for i := 0; i < 5; i++ {
		users, err := models.Users(
			models.UserWhere.ID.GT(lastUserID),
			qm.Limit(3),
		).All(ctx, conn)
		internal.LogFatal(err)
		internal.PrintJSONLog(users)
		lastUserID = users[len(users)-1].ID
	}
}

func aggregate(ctx context.Context, conn *sql.DB) {
	var result []struct {
		UserID int64 `boil:"user_id"`
		Count  int64 `boil:"count"`
	}
	err := models.Tasks(
		qm.Select("count(*) as count", models.TaskColumns.UserID),
		qm.GroupBy(models.TaskColumns.UserID),
		qm.OrderBy("count DESC"),
		qm.Limit(10),
	).Bind(ctx, conn, &result)
	internal.LogFatal(err)
	internal.PrintJSONLog(result)
}

func seed(ctx context.Context, conn *sql.DB, count int) {
	if count == 0 {
		return
	}

	rand.Seed(time.Now().UnixNano())

	clearModels(ctx, conn)

	var newUsers []*models.User
	for i := 0; i < count; i++ {
		newUser := &models.User{
			FirstName: "Sample",
			LastName:  "Unknown",
			Birthday:  time.Now().AddDate(-30, 0, 0),
		}
		err := newUser.Insert(ctx, conn, boil.Infer())
		internal.LogFatal(err)

		newUsers = append(newUsers, newUser)
	}

	var newTasks []*models.Task
	for i := 0; i < count; i++ {
		num := i + 1
		newTask := &models.Task{
			UserID: null.Int64{Int64: newUsers[rand.Intn(len(newUsers))].ID, Valid: true},
			Title:  fmt.Sprintf("Task %d", num),
			Note:   fmt.Sprintf("Note %d", num),
			Status: models.TaskStatusTodo,
		}
		err := newTask.Insert(ctx, conn, boil.Infer())
		internal.LogFatal(err)

		newTasks = append(newTasks, newTask)
	}

	log.Println("Seed finished")
}

func clearModels(ctx context.Context, conn *sql.DB) {
	_, err := models.Tasks().DeleteAll(ctx, conn)
	internal.LogFatal(err)
	_, err = models.Users().DeleteAll(ctx, conn)
	internal.LogFatal(err)
}

func queryWithRelation(ctx context.Context, conn *sql.DB) {
	// Create with relationship
	newUser := &models.User{
		FirstName: "Sample",
		LastName:  "User",
		Birthday:  time.Now().AddDate(-30, 0, 0),
	}
	err := newUser.Insert(ctx, conn, boil.Infer())
	internal.LogFatal(err)
	internal.PrintJSONLog(newUser)

	newTask := &models.Task{
		UserID: null.Int64{Int64: newUser.ID, Valid: true},
		Title:  "task 1",
		Note:   "note 1",
		Status: models.TaskStatusTodo,
	}
	err = newTask.Insert(ctx, conn, boil.Infer())
	internal.LogFatal(err)
	internal.PrintJSONLog(newTask)

	// Read with relationship
	gotTask, err := models.FindTask(ctx, conn, newTask.ID)
	internal.LogFatal(err)
	internal.PrintJSONLog(gotTask)

	gotTaskUser, err := gotTask.User().One(ctx, conn)
	internal.LogFatal(err)
	internal.PrintJSONLog(gotTaskUser)

	// Eager Loading
	gotTask, err = models.Tasks(qm.Load(models.TaskRels.User)).One(ctx, conn)
	internal.LogFatal(err)
	internal.PrintJSONLog(gotTask)
	internal.PrintJSONLog(gotTask.R.User)
}

func crud(ctx context.Context, conn *sql.DB) {
	// Create
	newUser := &models.User{
		FirstName: "Sample",
		LastName:  "Man",
		Birthday:  time.Now().AddDate(-30, 0, 0),
	}
	err := newUser.Insert(ctx, conn, boil.Infer())
	internal.LogFatal(err)
	internal.PrintJSONLog(newUser)

	// Read One
	gotUser, err := models.Users(
		models.UserWhere.FirstName.EQ("Sample"),
	).One(ctx, conn)
	internal.LogFatal(err)
	internal.PrintJSONLog(gotUser)

	// Read List
	gotUsers, err := models.Users(
		models.UserWhere.FirstName.EQ("Sample"),
	).All(ctx, conn)
	internal.LogFatal(err)
	internal.PrintJSONLog(gotUsers)

	// Update
	updateRowsAff, err := models.Users(
		models.UserWhere.LastName.EQ("Man"),
	).UpdateAll(
		ctx,
		conn,
		models.M{models.UserColumns.LastName: "Unknown"},
	)
	internal.LogFatal(err)
	internal.PrintJSONLog(updateRowsAff)

	// Delete
	deleteRowsAff, err := models.Users().DeleteAll(ctx, conn)
	internal.LogFatal(err)
	internal.PrintJSONLog(deleteRowsAff)
}
