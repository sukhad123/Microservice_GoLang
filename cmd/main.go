/**Entry point	*/
package main

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/learn_go/internal/env"
)

/**Entry point of main function*/
func main() {

	//godotenv
	cfg := &config{
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "postgresql://neondb_owner:npg_9ifW6ZHApLjx@ep-royal-credit-ad5tp1vo-pooler.c-2.us-east-1.aws.neon.tech/neondb?sslmode=require"),
		},
	}

	//Database
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		logger.Error("db Connection error" + err.Error())
		//os exit
		os.Exit(1)
	}
	defer conn.Close(ctx)
	logger.Info("Connected to Database")

	api := &application{
		config: *cfg,
		db:     conn,
	}

	//logger: Structured logger
	//

	//mount and run
	h := api.mount()
	if err := api.run(h); err != nil {
		logger.Error("Server has failed to start")
	}
}
