package main

import (
	"fmt"
	"net/http"
)

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
