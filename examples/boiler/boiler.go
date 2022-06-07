package boiler

import (
	"context"
	"database/sql"
	"github.com/ugabiga/go-orm-example/config"
	"github.com/ugabiga/go-orm-example/examples/boiler/models"
	"github.com/ugabiga/go-orm-example/internal"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"time"
)

func Execute() {
	// Connection
	ctx := context.Background()
	conn, err := sql.Open(config.GetDBInfo())
	internal.LogFatal(err)

	// Create
	newUser := &models.User{
		FirstName: "Sample",
		LastName:  "Man",
		Birthday:  time.Now().AddDate(-30, 0, 0),
	}
	err = newUser.Insert(ctx, conn, boil.Infer())
	internal.LogFatal(err)
	internal.PrintJSONLog(newUser)

	// Read One
	gotUser, err := models.Users(models.UserWhere.FirstName.EQ("Sample")).One(ctx, conn)
	internal.LogFatal(err)
	internal.PrintJSONLog(gotUser)

	// Read List
	gotUsers, err := models.Users(models.UserWhere.FirstName.EQ("Sample")).All(ctx, conn)
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

	// CRUD with Relationship
	// Eager Loading
	// Aggregation
	// Pagination
	// Transform
	// Raw Query
	// Event
}
