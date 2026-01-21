package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/Eraguzy/PolyNew/httpmanager"
	"github.com/Eraguzy/PolyNew/internal/storage"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load() // Load .env file

	db, err := storage.ConnectToDB() // open db connection
	if err != nil {
		fmt.Printf("error connecting to db: %s\n", err)
		os.Exit(1)
	}

	mux := http.NewServeMux()
	mux = httpmanager.TestRoutes(mux)
	mux = httpmanager.UpdaterRoutes(
		mux,
		httpmanager.DBHandler{
			Db: storage.DBManager{Pool: db},
		})
	fmt.Println("starting server on :3010")

	err = http.ListenAndServe(":3010", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error when starting server: %s\n", err)
		os.Exit(1)
	}
}

//  limit := 500
//  for offset := 0; ; offset += 500 {
//  	args := []messenger.GetArgs{
//  		{Param: messenger.URLParamActive, Value: "true"},
//  		{Param: messenger.URLParamClosed, Value: "false"},
//  		{Param: messenger.URLParamLimit, Value: fmt.Sprintf("%d", limit)},
//  		{Param: messenger.URLParamOffset, Value: fmt.Sprintf("%d", offset)},
//  	}

//  	body, err := messenger.SendGetRequest(
//  		messenger.URLGammaAPI,
//  		messenger.PathEvents,
//  		args,
//  	)
//  	if err != nil {
//  		fmt.Println("error:", err)
//  		break
//  	}

//  	fmt.Println(string(body))
//  	fmt.Println("offset:", offset)

//  	if string(body) == "[]" {
//  		fmt.Println("done")
//  		break
//  	}
//  }
//
