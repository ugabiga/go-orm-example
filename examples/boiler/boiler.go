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
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"log"
	"math/rand"
	"strings"
	"time"
)

func Execute() {
	ctx := context.Background()

	conn := makeConnection()
	crud(ctx, conn)
	queryWithRelation(ctx, conn)
	seed(ctx, conn, 100)
	aggregate(ctx, conn)
	pagination(ctx, conn)
	transform(ctx, conn)
	rawQuery(ctx, conn)
	hook(ctx, conn)
}

func hook(ctx context.Context, conn *sql.DB) {
	taskInsertHook := func(ctx context.Context, exec boil.ContextExecutor, t *models.Task) error {
		log.Println("Insert into task :", t.Title)
		return nil
	}
	models.AddTaskHook(boil.BeforeInsertHook, taskInsertHook)
	newTask := &models.Task{
		Title:  "task 1",
		Note:   "note 1",
		Status: models.TaskStatusTodo,
	}
	err := newTask.Insert(ctx, conn, boil.Infer())
	internal.LogFatal(err)
}

func rawQuery(ctx context.Context, conn *sql.DB) {
	var task struct {
		ID    int64  `json:"id"`
		Title string `json:"title"`
	}
	err := queries.Raw(
		"SELECT * FROM tasks WHERE title LIKE '%Task%'",
	).Bind(ctx, conn, &task)
	internal.LogFatal(err)
	internal.PrintJSONLog(task)
}

func transform(ctx context.Context, conn *sql.DB) {
	task, err := models.Tasks(
		models.TaskWhere.UserID.IsNotNull(),
		qm.Load(models.TaskRels.User),
	).One(ctx, conn)
	internal.LogFatal(err)

	// Struct way
	internal.PrintJSONLog(
		struct {
			*models.Task
			R interface{}
		}{
			task,
			task.R,
		},
	)

	// Functional way
	buildTaskOutput := func(t *models.Task) map[string]interface{} {
		output := map[string]interface{}{
			"id":         t.ID,
			"user_id":    t.UserID,
			"child_id":   t.ChildID,
			"title":      t.Title,
			"note":       t.Note,
			"status":     t.Status,
			"updated_at": t.UpdatedAt,
			"created_at": t.CreatedAt,
			"R":          t.R,
		}
		return output
	}
	internal.PrintJSONLog(buildTaskOutput(task))

	// Hide Column
	convertSelectStatement := func(target []string) string {
		return strings.Join(target, ", ")
	}
	taskDefaultColumns := []string{
		models.TaskColumns.Title,
		models.TaskColumns.Note,
		models.TaskColumns.Status,
	}
	task, err = models.Tasks(
		qm.Select(convertSelectStatement(taskDefaultColumns)),
	).One(ctx, conn)
	internal.PrintJSONLog(task)

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
			LastName:  "User",
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
		LastName:  "User",
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
		models.UserWhere.LastName.EQ("User"),
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

func makeConnection() *sql.DB {
	conn, err := sql.Open(config.GetDBInfo())
	internal.LogFatal(err)
	return conn
}
