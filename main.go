package main

import (
	"fmt"
	"net/http"
)

type Todo struct {
	ID   int
	Text string
	Done bool
}

var todos = []Todo{
	{ID: 1, Text: "Learning HTMX", Done: false},
	{ID: 1, Text: "Learning Templ", Done: false},
}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/todos", handleTodos)
	http.HandleFunc("/toggle/", handleToggle)

	http.ListenAndServe(":6543", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	component := home(todos)

	component.Render(r.Context(), w)
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()

		newTodo := Todo{
			ID:   len(todos) + 1,
			Text: r.FormValue("text"),
		}

		todos = append(todos, newTodo)
		todoItem(newTodo).Render(r.Context(), w)
		return
	}
}

func handleToggle(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/toggle/"):]
	for i := range todos {
		if fmt.Sprint(todos[i].ID) == id {
			todos[i].Done = !todos[i].Done
			todoItem(todos[i]).Render(r.Context(), w)
			return
		}
	}
}
