package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// * Api Routes Handler
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

	// newUser.Id = len(users) + 1
	// users = append(users, newUser)

	query := `
		insert into users (name, email, age)
		values ($1, $2, $3)
		returning id
	`
	err = db.QueryRow(context.Background(), query,
		newUser.Name, newUser.Email, newUser.Age).Scan(&newUser.Id)

	if err != nil {
		fmt.Println(err)

		res := Response{
			Success: false,
			Message: "Failed to create user",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

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

	// b, err := json.Marshal(res)

	// if err != nil {
	// 	http.Error(w, "failed to encode response", http.StatusInternalServerError)
	// 	return
	// }

	// pros: memory efficient
	// encoder := json.NewEncoder(w)
	// encoder.Encode(res)

	query := `
		select id, name, email, age
		from users
	`

	rows, err := db.Query(context.Background(), query)

	if err != nil {
		res := Response{
			Success: false,
			Message: "Failed to fetch users",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Age)

		if err != nil {
			fmt.Println(err)

			res := Response{
				Success: false,
				Message: "Failed to scan users",
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(res)
			return
		}

		users = append(users, user)
	}

	res := Response{
		Success: true,
		Message: "Users fetched successfully",
		Data:    users,
	}

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

	// for _, user := range users {
	// 	if user.Id == id {
	// 		res := Response{
	// 			Success: true,
	// 			Message: "User fetched successfully",
	// 			Data:    user,
	// 		}

	// 		w.Header().Set("Content-Type", "application/json")
	// 		w.WriteHeader(http.StatusOK)
	// 		json.NewEncoder(w).Encode(res)
	// 		return
	// 	}
	// }

	var user User

	query := `
		select id, name, email, age
		from users
		where id = $1
	`

	err = db.QueryRow(context.Background(), query, id).Scan(
		&user.Id, &user.Name, &user.Email, &user.Age)

	if err != nil {
		fmt.Println(err)

		res := Response{
			Success: false,
			Message: "User not found",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(res)
		return
	}

	res := Response{
		Success: true,
		Message: "User fetched successfully",
		Data:    user,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
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

	// for idx, user := range users {
	// 	if user.Id == id {
	// 		userData.Id = id
	// 		users[idx] = userData

	// 		res := Response{
	// 			Success: true,
	// 			Message: "User fetched successfully",
	// 			Data:    userData,
	// 		}

	// 		w.Header().Set("Content-Type", "application/json")
	// 		w.WriteHeader(http.StatusOK)
	// 		json.NewEncoder(w).Encode(res)
	// 		return
	// 	}
	// }

	query := `
		update users
		set name = $1, email = $2, age = $3
		where id = $4
		returning id, name, email, age
	`

	err = db.QueryRow(context.Background(), query,
		userData.Name, userData.Email, userData.Age, id).Scan(
		&userData.Id, &userData.Name, &userData.Email, &userData.Age)

	if err != nil {
		fmt.Println(err)

		res := Response{
			Success: false,
			Message: "Failed to update user",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
		return
	}

	res := Response{
		Success: true,
		Message: "User updated successfully",
		Data:    userData,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func deleteUserById(w http.ResponseWriter, r *http.Request) {
	id, err := getId(w, r)

	if err != nil {
		return
	}

	// for idx, user := range users {
	// 	if user.Id == id {
	// 		// users = append(users[:idx], users[idx+1:]...)
	// 		users = slices.Delete(users, idx, idx+1)

	// 		res := Response{
	// 			Success: true,
	// 			Message: "User fetched successfully",
	// 			Data:    user,
	// 		}

	// 		w.Header().Set("Content-Type", "application/json")
	// 		w.WriteHeader(http.StatusOK)
	// 		json.NewEncoder(w).Encode(res)
	// 		return
	// 	}
	// }

	// var deletedUser User

	query := `
		delete from users
		where id = $1
		returning id, name, email, age
	`

	// err = db.QueryRow(context.Background(), query, id).Scan(
	// 	&deletedUser.Id, &deletedUser.Name, &deletedUser.Email, &deletedUser.Age)

	tag, err := db.Exec(context.Background(), query, id)

	if err != nil {
		fmt.Println(err)

		res := Response{
			Success: false,
			Message: "Failed to delete user",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(res)
	}

	if tag.RowsAffected() != 1 {
		res := Response{
			Success: false,
			Message: "User not found",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(res)
		return
	}

	res := Response{
		Success: true,
		Message: "User deleted successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
