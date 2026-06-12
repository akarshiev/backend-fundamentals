package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "Ali"},
	{ID: 2, Name: "Bob"},
	{ID: 3, Name: "Charlie"},
}

func main() {
	http.HandleFunc("/users", handleUsers)
	http.HandleFunc("/users/", handleUserByID)

	fmt.Println("REST API server listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		// Pagination
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		size, _ := strconv.Atoi(r.URL.Query().Get("size"))

		if page < 1 {
			page = 1
		}
		if size < 1 {
			size = 10
		}

		start := (page - 1) * size
		end := start + size
		if end > len(users) {
			end = len(users)
		}

		json.NewEncoder(w).Encode(users[start:end])

	case http.MethodPost:
		var user User
		json.NewDecoder(r.Body).Decode(&user)
		user.ID = len(users) + 1
		users = append(users, user)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid id"})
		return
	}

	var found *User
	for i := range users {
		if users[i].ID == id {
			found = &users[i]
			break
		}
	}

	if found == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "user not found"})
		return
	}

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(found)

	case http.MethodPut:
		var user User
		json.NewDecoder(r.Body).Decode(&user)
		user.ID = id
		*found = user
		json.NewEncoder(w).Encode(user)

	case http.MethodPatch:
		var updates map[string]string
		json.NewDecoder(r.Body).Decode(&updates)
		if name, ok := updates["name"]; ok {
			found.Name = name
		}
		json.NewEncoder(w).Encode(found)

	case http.MethodDelete:
		users = append(users[:id-1], users[id:]...)
		w.WriteHeader(http.StatusNoContent)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
