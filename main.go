package main

import (
	"log"
	"net/http"
	"todoBackend/tasks"
)

func main() {
	data := make([]tasks.Task, 0)
	tasksMgr := tasks.TaskMgr{
		Tasks: data,
		Count: 0,
	}
	middleware := tasks.Chain(
		tasks.ContentTypeMiddleware("application/json"),
		tasks.MaxBodySizeMiddleware(1<<20),
	)
	apiMux := http.NewServeMux()   // global mux - глобальный роутер
	tasksMux := http.NewServeMux() // tasks API mux - роутер API задач

	tasksMux.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		tasks.HandleTasks(w, r, &tasksMgr)
	})
	apiMux.Handle("/api/", http.StripPrefix("/api", middleware(tasksMux)))

	log.Printf("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", apiMux))
}
