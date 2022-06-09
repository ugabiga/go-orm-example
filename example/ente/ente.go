package ente

import (
	"ariga.io/atlas/sql/migrate"
	"context"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"fmt"
	"github.com/ugabiga/go-orm-example/config"
	"github.com/ugabiga/go-orm-example/example/ente/ent"
	"github.com/ugabiga/go-orm-example/example/ente/ent/task"
	"github.com/ugabiga/go-orm-example/example/ente/ent/user"
	"github.com/ugabiga/go-orm-example/internal"
	"log"
	"math/rand"
	"time"
)

func GenerateMigration() {
	basePath := "./example/ente"
	graph, err := entc.LoadGraph(basePath+"/ent/schema", &gen.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	tbls, err := graph.Tables()
	if err != nil {
		log.Fatalln(err)
	}

	dir, err := migrate.NewLocalDir(basePath + "/migrations")
	if err != nil {
		log.Fatalln(err)
	}

	dlct, err := sql.Open(config.GetDBInfo())
	if err != nil {
		log.Fatalln(err)
	}

	m, err := schema.NewMigrate(dlct, schema.WithDir(dir))
	if err != nil {
		log.Fatalln(err)
	}

	if err := m.Diff(context.Background(), tbls...); err != nil {
		log.Fatalln(err)
	}
}

func Execute() {
	ctx := context.Background()
	conn := makeConnection()

	fmt.Println("Run CRUD")
	crud(ctx, conn)
	fmt.Println()

	fmt.Println("Run QueryWithRelation")
	queryWithRelation(ctx, conn)
	fmt.Println()

	fmt.Println("Run Seed")
	seed(ctx, conn, 100)
	fmt.Println()

	//fmt.Println("Run Aggregation")
	//aggregate(ctx, conn)
	//fmt.Println()
	//
	//fmt.Println("Run Pagination")
	//pagination(ctx, conn)
	//fmt.Println()
	//
	//fmt.Println("Run Transform")
	//transform(ctx, conn)
	//fmt.Println()
	//
	//fmt.Println("Run Raw Query")
	//rawQuery(ctx, conn)
	//fmt.Println()
	//
	//fmt.Println("Run Hook")
	//hook(ctx, conn)
	//fmt.Println()
}

func seed(ctx context.Context, c *ent.Client, count int) {
	if count <= 0 {
		return
	}

	rand.Seed(time.Now().UnixNano())

	var bulkUsers []*ent.UserCreate
	for i := 0; i < count; i++ {
		bulkUsers = append(bulkUsers, c.User.Create().
			SetFirstName("Sample").
			SetLastName("User").
			SetBirthday(time.Now().AddDate(-30, 0, 0)),
		)
	}
	newUsers, err := c.User.CreateBulk(bulkUsers...).Save(ctx)
	internal.LogFatal(err)

	var bulkTasks []*ent.TaskCreate
	for i := 0; i < count; i++ {
		bulkTasks = append(bulkTasks, c.Task.Create().
			SetUser(newUsers[rand.Intn(len(newUsers))]).
			SetTitle(fmt.Sprintf("Task %d", i)).
			SetNote(fmt.Sprintf("Note %d", i)).
			SetStatus(task.StatusTodo),
		)
	}
	_, err = c.Task.CreateBulk(bulkTasks...).Save(ctx)
	internal.LogFatal(err)

	log.Println("Seed finished")
}

func queryWithRelation(ctx context.Context, c *ent.Client) {
	// Create with relationship
	newUser, err := c.User.Create().
		SetFirstName("Sample").
		SetLastName("User").
		SetBirthday(time.Now().AddDate(-30, 0, 0)).
		Save(ctx)
	internal.LogFatal(err)
	internal.PrintJSONLog(newUser)

	newTask, err := c.Task.Create().
		SetUser(newUser).
		SetTitle("task 1").
		SetNote("note 1").
		SetStatus(task.StatusTodo).
		Save(ctx)
	internal.LogFatal(err)
	internal.PrintJSONLog(newTask)

	// Read with relationship
	gotTask, err := c.Task.Get(ctx, newTask.ID)
	internal.LogFatal(err)
	internal.PrintJSONLog(gotTask)

	gotTaskUser, err := gotTask.QueryUser().First(ctx)
	internal.LogFatal(err)
	internal.PrintJSONLog(gotTaskUser)

	// Eager Loading
	tasks, err := c.Task.Query().
		WithUser().
		All(ctx)
	internal.LogFatal(err)
	internal.PrintJSONLog(tasks)
}

func crud(ctx context.Context, c *ent.Client) {
	// Create
	newUser, err := c.User.
		Create().
		SetFirstName("Sample").
		SetLastName("User").
		SetBirthday(time.Now().AddDate(-30, 0, 0)).
		Save(ctx)
	internal.LogFatal(err)
	internal.PrintJSONLog(newUser)

	// Read
	gotUser, err := c.User.
		Query().
		Where(user.FirstNameEQ("Sample")).
		First(ctx)
	internal.LogFatal(err)
	internal.PrintJSONLog(gotUser)

	// Read List
	gotUsers, err := c.User.Query().Where(user.FirstNameEqualFold("Sample")).All(ctx)
	internal.LogFatal(err)
	internal.PrintJSONLog(gotUsers)

	// Update
	updateRowsAff, err := c.User.Update().
		Where(user.LastNameEQ("User")).
		SetLastName("Unknown").
		Save(ctx)
	internal.LogFatal(err)
	internal.PrintJSONLog(updateRowsAff)

	// Delete
	deleteRowsAff, err := c.User.Delete().
		Exec(ctx)
	internal.LogFatal(err)
	internal.PrintJSONLog(deleteRowsAff)
}

func makeConnection() *ent.Client {
	conn, err := ent.Open(config.GetDBInfo())
	internal.LogFatal(err)

	return conn
}
