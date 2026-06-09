package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// * Utils
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
