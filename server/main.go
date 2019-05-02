package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/datastore"
	"test-golang-cache.appspot.com/hello"
)

func main() {
	// Use Cloud Datastore, even if we don't need it, to artificially increase the time required to fetch dependencies.
	ctx := context.Background()
	_, err := datastore.NewClient(ctx, "jevoteclimat")
	if err != nil {
		panic(err)
	}

	// Setup HTTP handlers
	http.HandleFunc("/", handleIndex)

	// Start HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(hello.Hello()))
}
