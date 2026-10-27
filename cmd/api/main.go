package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/NotPro955/backend/internal/config"
	"github.com/NotPro955/backend/internal/db"
	"github.com/NotPro955/backend/internal/handler"
	"github.com/NotPro955/backend/internal/middleware"
)

func main() {
	conf := config.MustLoad()
	db, dberr := db.Connect(conf.Databaseurl)
	if dberr != nil {
		log.Fatalf("main db: %v", dberr)
	}
	fmt.Println("DataBase Connected")
	fmt.Println("Starting the Backend Server")

	logHandle := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelWarn,
	})
	logger := slog.New(logHandle)
	slog.SetDefault(logger)

	mux := http.NewServeMux()

	lh := handler.NewListingHandler(db)

	mux.HandleFunc("GET /healthz", handler.Health)
	mux.HandleFunc("GET /listings", lh.List)
	mux.HandleFunc("DELETE /listings/{id}", lh.Delete)
	mux.HandleFunc("POST /listings", lh.Create)

	handler := middleware.RequestId(mux)

	srv := http.Server{
		Addr:         ":" + conf.Port,
		Handler:      handler,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
