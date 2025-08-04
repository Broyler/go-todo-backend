package tasks

import (
	"log"
	"net/http"
)

func put(w http.ResponseWriter, r *http.Request, m *TaskMgr) {
	task, err := getTaskPostData(r)
	if err != nil {
		log.Printf("Error parsing JSON: %s", err)
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}
	aux := m.PutTask(task)
	writeTask(w, aux)
}
