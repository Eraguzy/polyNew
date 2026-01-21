package main

import (
	"context"
	"errors"
	"fmt"
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
		fmt.Printf("error connecting to db: %s\n", err)
		os.Exit(1)
	}

	// setup http server with routes
	mux := http.NewServeMux()
	mux = httpmanager.TestRoutes(mux)
	mux = httpmanager.UpdaterRoutes(
		mux,
		httpmanager.DBHandler{
			Db: storage.DBManager{Pool: db},
		},
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
		fmt.Printf("error when starting server: %s\n", err)
		os.Exit(1)
	}
}
