package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"net/http"
)

var db *pgx.Conn

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// * Main Function of the program
func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	// Connect to DB
	connectDb()

	// Close DB connection at the end of the program
	defer db.Close(context.Background())

	// Start Server
	server()

	fmt.Println("Program is closed")
}

// * Server
func server() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", rootHandler)
	mux.HandleFunc("GET /health", healthHandler)

	mux.HandleFunc("POST /users", createUserHandler)
	mux.HandleFunc("GET /users", getUsersHandler)
	mux.HandleFunc("GET /users/{id}", getUserById)
	mux.HandleFunc("PATCH /users/{id}", updateUserById)
	mux.HandleFunc("DELETE /users/{id}", deleteUserById)

	fmt.Println("Server is running on http://localhost:5000")
	err := http.ListenAndServe(":5000", mux)

	if err != nil {
		fmt.Println("Server error:", err)
	}
}
