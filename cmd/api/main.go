package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *slog.Logger
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Nothing to see here."))
}

// Retrieve origination context for a specific user interaction with your
// products user-interface
func getOriginsByOriginId(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("getting origin data..."))
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment [development|staging|production]")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		config: cfg,
		logger: logger,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	//mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)
	//mux.HandleFunc("/partner/v2/origins/:origin_id", getOriginsByOriginId)
	//mux.HandleFunc("/partner/v2/transactions", getAllTransactions)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", "addr", srv.Addr, "env", cfg.env)

	err := srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)

}
