package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	quickstartURL := os.Getenv("STELLAR_QUICKSTART_URL")
	fmt.Println("Quickstart URL: ", quickstartURL)
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
