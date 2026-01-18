package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/Eraguzy/PolyNew/httpmanager"
)

func main() {
	mux := http.NewServeMux()
	mux = httpmanager.TestRoutes(mux)

	fmt.Println("starting server on :3010")

	err := http.ListenAndServe(":3010", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

// 	// limit := 500
// 	// for offset := 0; ; offset += 500 {
// 	// 	args := []polymarket.GetArgs{
// 	// 		{Param: polymarket.URLParamActive, Value: "true"},
// 	// 		{Param: polymarket.URLParamClosed, Value: "false"},
// 	// 		{Param: polymarket.URLParamLimit, Value: fmt.Sprintf("%d", limit)},
// 	// 		{Param: polymarket.URLParamOffset, Value: fmt.Sprintf("%d", offset)},
// 	// 	}

// 	// 	body, err := polymarket.SendGetRequest(
// 	// 		polymarket.URLGammaAPI,
// 	// 		polymarket.PathEvents,
// 	// 		args,
// 	// 	)
// 	// 	if err != nil {
// 	// 		fmt.Println("error:", err)
// 	// 		break
// 	// 	}

// 	// 	fmt.Println(string(body))
// 	// 	fmt.Println("offset:", offset)

// 	// 	if string(body) == "[]" {
// 	// 		fmt.Println("done")
// 	// 		break
// 	// 	}
// 	// }
// }

// func main() {
// 	godotenv.Load() // Load .env file

// 	dbpool, err := storage.ConnectToDB()
// 	if err != nil {
// 		log.Fatalf("Unable to create connection pool: %v\n", err)
// 	}
// 	defer dbpool.Close()

// 	var greeting string
// 	err = dbpool.QueryRow(context.Background(), "select 'Hello, db!'").Scan(&greeting)
// 	if err != nil {
// 		log.Fatalf("QueryRow failed: %v\n", err)
// 	}

// 	fmt.Println(greeting)
// }
