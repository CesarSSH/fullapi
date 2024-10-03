package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/CesarSSH/go-pg-connection/db"
	"github.com/CesarSSH/go-pg-connection/types"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const SERVER_PORT string = ":3690"

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading the .env file.")
	}

	// Open connection to Database
	database, err := sql.Open("postgres", db.GetDSN())
	if err != nil {
		panic(err.Error())
	}
	defer database.Close()

	// Check success connection to Database
	err = database.Ping()
	if err != nil {
		panic(err.Error())
	}

	// Define handlers to endpoints
	http.HandleFunc("/api/v1/user", func(w http.ResponseWriter, r *http.Request) {
		types.HandleUser(w, r, database)
	})

	println("Server listening on port" + SERVER_PORT + " ...")
	http.ListenAndServe(SERVER_PORT, nil)
}
