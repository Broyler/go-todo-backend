package main

import (
	"log"
	"net/http"
)

func main() {
	tasks := make([]Task, 0)
	tasksMgr := TasksMgr{
		Tasks: tasks,
		Count: 0,
	}
	middleware := Chain(
		ContentTypeMiddleware("application/json"),
		MaxBodySizeMiddleware(1<<20),
	)
	apiMux := http.NewServeMux()   // global mux - глобальный роутер
	tasksMux := http.NewServeMux() // tasks API mux - роутер API задач

	tasksMux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		tasksAPI(w, r, &tasksMgr) // passing a pointer to slice of tasks - передаем ссылку на слайс с задачами
	})
	apiMux.Handle("/api/", http.StripPrefix("/api", middleware(tasksMux)))

	log.Printf("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", apiMux))
}
