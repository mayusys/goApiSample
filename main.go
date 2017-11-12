package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Todo struct {
	ID          string `json:"id,omitempty"`
	TaskName    string `json:"taskname,omitempty"`
	CompleteFlg bool   `json:"completeflg,omitempty"`
}

var todos []Todo

func main() {
	router := mux.NewRouter()

	todos = append(todos, Todo{ID: "1", TaskName: "Learn Programming", CompleteFlg: false})
	todos = append(todos, Todo{ID: "2", TaskName: "Diet", CompleteFlg: true})
	todos = append(todos, Todo{ID: "3", TaskName: "Play the Violin", CompleteFlg: false})

	router.HandleFunc("/todo", GetTodos).Methods("GET")
	router.HandleFunc("/todo/{id}", GetTodo).Methods("GET")
	router.HandleFunc("/todo/{id}", CreateTodo).Methods("POST")
	router.HandleFunc("/todo/{id}", DeleteTodo).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todos)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range todos {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)

	todo.ID = params["id"]
	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todos)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range todos {
		if item.ID == params["id"] {
			todos = append(todos[:index], todos[index+1:]...)
			newTodo := make([]Todo, len(todos))
			copy(newTodo, todos)
			break
		}
	}
	json.NewEncoder(w).Encode(todos)
}
