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

var todo []Todo

func main() {
	router := mux.NewRouter()

	todo = append(todo, Todo{ID: "1", TaskName: "Learn Programming", CompleteFlg: false})
	todo = append(todo, Todo{ID: "2", TaskName: "Diet", CompleteFlg: true})
	todo = append(todo, Todo{ID: "3", TaskName: "Play the Violin", CompleteFlg: false})

	router.HandleFunc("/todo", GetTodo).Methods("GET")
	router.HandleFunc("/todo/{id}", GetTodo).Methods("GET")
	router.HandleFunc("/todo/{id}", CreateTodo).Methods("POST")
	router.HandleFunc("/todo/{id}", DeleteTodo).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todo)
}

func GetTodos(w http.ResponseWriter, r *http.Request)   {}
func CreateTodo(w http.ResponseWriter, r *http.Request) {}
func DeleteTodo(w http.ResponseWriter, r *http.Request) {}
