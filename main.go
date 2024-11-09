package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	db "github.com/Zync1402/ZyncifyServer/database"
	"github.com/Zync1402/ZyncifyServer/models"
	"github.com/gorilla/mux"
)

func main() {
	db := db.Connect()

	r := mux.NewRouter()

	r.HandleFunc("/api/getTodos", func(w http.ResponseWriter, r *http.Request) {
		var todos []models.Todos
		db.Where(models.Todos{UserID: "ch123"}).Find(&todos)
		json.NewEncoder(w).Encode(todos)
	}).Methods("GET")

	r.HandleFunc("/api/getTodo", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "Missing 'id' query parameter", http.StatusBadRequest)
			return
		}
		var todos []models.Todos
		db.Where(models.Todos{UserID: "ch123", Id: id}).First(&todos)
		json.NewEncoder(w).Encode(todos)
	}).Methods("GET")

	r.HandleFunc("/api/addTodo", func(w http.ResponseWriter, r *http.Request) {
		var Todo models.Todos
		decoder := json.NewDecoder(r.Body)
		res := decoder.Decode(&Todo)
		if res != nil {
			http.Error(w, res.Error(), http.StatusBadRequest)
			return
		}
		db.Create(&Todo)
	}).Methods("POST")

	r.HandleFunc("/api/updateTodo", func(w http.ResponseWriter, r *http.Request) {
		var Todo models.Todos
		decoder := json.NewDecoder(r.Body)
		res := decoder.Decode(&Todo)
		if res != nil {
			http.Error(w, res.Error(), http.StatusBadRequest)
			return
		}
		db.Where(models.Todos{UserID: "ch123"}).Save(&Todo)
	}).Methods("PATCH")

	r.HandleFunc("/api/deleteTodo", func(w http.ResponseWriter, r *http.Request) {
		var Todo models.Todos
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "Missing 'id' query parameter", http.StatusBadRequest)
			return
		}
		db.Where(models.Todos{UserID: "ch123", Id: id}).Delete(&Todo)
	}).Methods("DELETE")

	fmt.Println("Server running on port 3001")
	log.Fatal(http.ListenAndServe(":3001", r))
}
