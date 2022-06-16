package sqc

import (
	"context"
	"database/sql"
	"time"

	"github.com/ugabiga/go-orm-example/config"
	db "github.com/ugabiga/go-orm-example/example/sqc/db/sqlc"
	"github.com/ugabiga/go-orm-example/internal"
)

func Execute() {
	ctx := context.Background()
	conn, err := sql.Open(config.GetDBInfo())
	internal.LogFatal(err)

	query := db.New(conn)

	newUser, err := query.CreateUser(ctx, db.CreateUserParams{
		FirstName: "Sample",
		LastName:  "User",
		Birthday:  time.Now().AddDate(-30, 0, 0),
	})
	internal.LogFatal(err)
	internal.PrintJSONLog(newUser)

	user, err := query.GetUser(ctx, newUser.ID)
	internal.LogFatal(err)
	internal.PrintJSONLog(user)
}
