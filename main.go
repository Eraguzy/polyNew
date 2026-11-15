package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Eraguzy/PolyNew/internal/actions"
	"github.com/Eraguzy/PolyNew/internal/httpmanager"
	"github.com/Eraguzy/PolyNew/internal/storage"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load() // Load .env file

	// open db connection
	db, err := storage.ConnectToDB()
	if err != nil {
		log.Fatalf("error connecting to db: %s", err)
	}

	// setup http server with routes
	var user = os.Getenv("BASICAUTH_USER")
	var pass = os.Getenv("BASICAUTH_PASS")
	if user == "" || pass == "" {
		log.Fatal("BASICAUTH_USER or BASICAUTH_PASS not set")
	}
	mux := http.NewServeMux()
	mux = httpmanager.TestRoutes(mux)
	mux = httpmanager.UpdaterRoutes(
		mux,
		httpmanager.DBHandler{
			Db: storage.DBManager{Pool: db},
		},
		user,
		pass,
	)

	// self polling
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	actions.Polling(ctx, storage.DBManager{Pool: db})

	fmt.Println("starting server on :3010")
	err = http.ListenAndServe(":3010", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		log.Fatalf("error when starting server: %s", err)
	}
}
