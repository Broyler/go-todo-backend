package main

import (
	"log"
	"net/http"
)

func main() {
	// tasks := make([]Task, 0)
	apiMux := http.NewServeMux()
	tasksMux := http.NewServeMux()

	tasksMux.HandleFunc("/tasks", tasksAPI)
	apiMux.Handle("/api/", http.StripPrefix("/api", tasksMux))

	log.Printf("Listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", apiMux))
}
