package tasks

import (
	"log"
	"net/http"
)

func _delete(w http.ResponseWriter, m *TaskMgr, id int) {
	m.Lock()
	defer m.Unlock()

	for i, task := range m.Tasks {
		if task.ID == id {
			// Todo: rewrite with linked list for performance
			m.Tasks = append(m.Tasks[:i], m.Tasks[i+1:]...)
			if _, err := w.Write([]byte("Ok")); err != nil {
				log.Printf("Error writing http response: %s", err)
				http.Error(w, "Error writing http response", http.StatusInternalServerError)
			}
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	if _, err := w.Write([]byte("A task with such ID does not exist.")); err != nil {
		log.Printf("Error writing http response: %s", err)
		http.Error(w, "Error writing http response", http.StatusInternalServerError)
	}
}
