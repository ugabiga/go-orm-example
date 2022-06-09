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
	"github.com/ugabiga/go-orm-example/internal"
	"log"
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

	//fmt.Println("Run QueryWithRelation")
	//queryWithRelation(ctx, conn)
	//fmt.Println()
	//
	//fmt.Println("Run Seed")
	//seed(ctx, conn, 100)
	//fmt.Println()
	//
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

func crud(ctx context.Context, c *ent.Client) {
	newUser, err := c.User.
		Create().
		SetFirstName("Sample").
		SetLastName("User").
		SetBirthday(time.Now().AddDate(-30, 0, 0)).
		Save(ctx)
	internal.LogFatal(err)
	internal.PrintJSONLog(newUser)
}

func makeConnection() *ent.Client {
	conn, err := ent.Open(config.GetDBInfo())
	internal.LogFatal(err)

	return conn
}
