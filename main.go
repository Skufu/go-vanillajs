package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("public")))
	const addr = "localhost:8080"
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	fmt.Println("Server is running on port 8080")
}
