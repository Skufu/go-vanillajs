package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/skufu/movies/logger"
)

func initializeLogger() *logger.Logger {
	logInstance, err := logger.NewLogger("movie.log")
	defer logInstance.Close()
	if err != nil {
		log.Fatalf("Error initializing logger: %v", err)
	}
	return logInstance
}

func main() {

	logInstance := initializeLogger()

	http.Handle("/", http.FileServer(http.Dir("public")))
	const addr = "localhost:8080"
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		log.Fatalf("Error starting server: %v", err)
		logInstance.Error("Server Failed", err)
	}
	fmt.Println("Server is running on port 8080")
}
