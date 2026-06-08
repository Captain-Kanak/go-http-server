package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strconv"
)

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

var users = []User{
	{
		Id:    1,
		Name:  "John Doe",
		Email: "jD0Hw@example.com",
		Age:   30,
	},
	{
		Id:    2,
		Name:  "Jane Doe",
		Email: "2b4e9@example.com",
		Age:   25,
	},
	{
		Id:    3,
		Name:  "Bob Smith",
		Email: "r9B2o@example.com",
		Age:   35,
	},
}

func main() {
	// Routes Handler
	server()

}

// Utils
func getId(w http.ResponseWriter, r *http.Request) (int, error) {
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		fmt.Println(err)

		res := Response{
			Success: false,
			Message: "Invalid id",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return 0, err
	}

	return id, nil
}

// Server
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

func rootHandler(w http.ResponseWriter, r *http.Request) {
	res := Response{
		Success: true,
		Message: "Welcome to Go server. Server is running...",
	}

	// b, err := json.Marshal(res)

	// if err != nil {
	// 	http.Error(w, "failed to encode response", http.StatusInternalServerError)
	// 	return
	// }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	res := Response{
		Success: true,
		Message: "Server is Healthy",
	}

	// b, err := json.Marshal(res)

	// if err != nil {
	// 	http.Error(w, "failed to encode response", http.StatusInternalServerError)
	// 	return
	// }

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		fmt.Println(err)

		res := Response{
			Success: false,
			Message: "Invalid request body",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	// fmt.Println(newUser)

	newUser.Id = len(users) + 1
	users = append(users, newUser)

	res := Response{
		Success: true,
		Message: "User created successfully",
		Data:    newUser,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	res := Response{
		Success: true,
		Message: "Users fetched successfully",
		Data:    users,
	}

	// b, err := json.Marshal(res)

	// if err != nil {
	// 	http.Error(w, "failed to encode response", http.StatusInternalServerError)
	// 	return
	// }

	// pros: memory efficient
	// encoder := json.NewEncoder(w)
	// encoder.Encode(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func getUserById(w http.ResponseWriter, r *http.Request) {
	id, err := getId(w, r)

	if err != nil {
		return
	}

	// fmt.Printf("the value of id is: %v and type is: %T", id, id)

	for _, user := range users {
		if user.Id == id {
			res := Response{
				Success: true,
				Message: "User fetched successfully",
				Data:    user,
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(res)
			return
		}
	}

	res := Response{
		Success: false,
		Message: "User not found",
		Data:    nil,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(res)
}

func updateUserById(w http.ResponseWriter, r *http.Request) {
	id, err := getId(w, r)

	if err != nil {
		return
	}

	var userData User

	err = json.NewDecoder(r.Body).Decode(&userData)

	if err != nil {
		fmt.Println(err)

		res := Response{
			Success: false,
			Message: "Invalid request body",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	for idx, user := range users {
		if user.Id == id {
			userData.Id = id
			users[idx] = userData

			res := Response{
				Success: true,
				Message: "User fetched successfully",
				Data:    userData,
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(res)
			return
		}
	}

	res := Response{
		Success: false,
		Message: "User not found",
		Data:    nil,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(res)
}

func deleteUserById(w http.ResponseWriter, r *http.Request) {
	id, err := getId(w, r)

	if err != nil {
		return
	}

	for idx, user := range users {
		if user.Id == id {
			// users = append(users[:idx], users[idx+1:]...)
			users = slices.Delete(users, idx, idx+1)

			res := Response{
				Success: true,
				Message: "User fetched successfully",
				Data:    user,
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(res)
			return
		}
	}

	res := Response{
		Success: false,
		Message: "User not found",
		Data:    nil,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(res)
}
