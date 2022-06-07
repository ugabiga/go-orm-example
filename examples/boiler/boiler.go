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
	"math/rand"
	"time"
)

func Execute() {
	// Connection
	ctx := context.Background()
	conn, err := sql.Open(config.GetDBInfo())
	internal.LogFatal(err)

	crud(ctx, conn)
	queryWithRelation(ctx, conn)
	seed(100, ctx, conn)

	// Aggregation
	// Pagination
	// Transform
	// Raw Query
	// Event
}

func seed(count int, ctx context.Context, conn *sql.DB) {
	if count == 0 {
		return
	}

	rand.Seed(time.Now().UnixNano())

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
	gotTask, err := models.Tasks().One(ctx, conn)
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