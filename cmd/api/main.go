package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/NotPro955/backend/internal/config"
	"github.com/NotPro955/backend/internal/db"
	"github.com/NotPro955/backend/internal/handler"
)

func main() {
	conf := config.MustLoad()
	db, dberr := db.Connect(conf.Databaseurl)
	if dberr != nil {
		log.Fatalf("main db: %v", dberr)
	}
	fmt.Println("DataBase Connected")
	fmt.Println("Starting the Backend Server")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", handler.Health)
	mux.HandleFunc("GET /listings", handler.List(db))

	srv := http.Server{
		Addr:         ":" + conf.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
