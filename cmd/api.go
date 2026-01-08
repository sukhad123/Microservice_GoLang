package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	products "github.com/learn_go/internal"
	repo "github.com/learn_go/internal/adapters/postgresql.sqlc"
)

var logger = slog.Default()

type application struct {
	config config
	db     *pgx.Conn
	//logger

	//db driver
}

// run -> graceful shutdown
// mount
func (app *application) mount() http.Handler {
	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID) //very important for rate limiting
	r.Use(middleware.RealIP)    // import for rate limiting and analytics and tracing
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) // recover from crashes

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))
	//health router
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("I have a health server for now"))
	})

	productService := products.NewService(repo.New(app.db))
	productHandler := products.NewHandler(productService)
	r.Get("/products", productHandler.ListProducts)
	r.Post("/products", productHandler.ListProducts)

	return r
}

// run
func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 30,
	}

	logger.Info("Starting server on port", "addr", app.config.addr)
	//

	return srv.ListenAndServe()

}

type config struct {
	addr string
	db   dbConfig
}

// type for db
type dbConfig struct {
	dsn string //username:password@protocol(address)/dbname?param=value
}
