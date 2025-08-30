package main

import (
	"fmt"
	"log"
	"net/http"

	"database/sql"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/skufu/movies/data"
	"github.com/skufu/movies/handlers"
	"github.com/skufu/movies/logger"
)

func initializeLogger() *logger.Logger {
	logInstance, err := logger.NewLogger("movie.log")
	if err != nil {
		log.Fatalf("Error initializing logger: %v", err)
	}
	// Remove defer as it will close the logger immediately, not when main() exits
	return logInstance
}

func main() {
	//initialize logger
	logInstance := initializeLogger()
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbConnStr := os.Getenv("DATABASE_URL")
	if dbConnStr == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	//initialize database connection
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	//initialize movie repository
	movieRepo, err := data.NewMovieRepository(db, logInstance)
	if err != nil {
		log.Fatalf("Failed to initialize movie repository: %v", err)
	}
	//initialize movie handler
	movieHandler := handlers.NewMovieHandler(movieRepo, logInstance)
	// authHandler := handlers.NewAuthHandler(userStorage, jwt, logInstance)

	// Set up routes - specific routes first
	http.HandleFunc("/api/movies/random/", movieHandler.GetRandomMovies) // Handle trailing slash
	http.HandleFunc("/api/movies/top/", movieHandler.GetTopMovies)       // Handle trailing slash
	http.HandleFunc("/api/movies/search/", movieHandler.SearchMovies)    // Handle trailing slash
	http.HandleFunc("/api/movies/", movieHandler.GetMovie)               // This should be last - it's the catch-all
	http.HandleFunc("/api/genres", movieHandler.GetGenres)
	// TODO: Implement proper auth handlers
	// http.HandleFunc("/api/account/register", authHandler.Register)
	// http.HandleFunc("/api/account/authenticate", authHandler.Authenticate)

	//handle static files(frontend)
	http.Handle("/", http.FileServer(http.Dir("public")))
	fmt.Println("Serving Files")
	const addr = "localhost:8080"

	if err := http.ListenAndServe(addr, nil); err != nil {
		logInstance.Error("Server Failed", err)
		log.Fatalf("Error starting server: %v", err)
	}

	fmt.Println("Server is running on port 8080")
	// Test comment for Air rebuild - updated
}
