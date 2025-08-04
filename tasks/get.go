package tasks

import (
	"encoding/json"
	"log"
	"net/http"
)

func get(w http.ResponseWriter, _ *http.Request, m *TaskMgr) {
	/* GET method for tasks API.
	   Will return a paginated JSON response with tasks.

	   Метод GET для API задач.
	   Вернет JSON задач с пагинацией. */
	tasks := m.GetTasks()
	response := struct {
		Count   int    `json:"count"`
		Results []Task `json:"results"`
	}{
		len(tasks),
		tasks,
	}

	w.Header().Add("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding JSON response: %s", err)
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}
}

func writeTask(w http.ResponseWriter, task Task) {
	/* A function to write single task response
	   Функция для записи ответа с одной задачей*/
	res, err := json.Marshal(task)
	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		log.Printf("Error encoding task to JSON: %s", err)
		http.Error(w, "Error encoding task to JSON", http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(res); err != nil {
		log.Printf("Error writing http response: %s", err)
		http.Error(w, "Error writing http response", http.StatusInternalServerError)
	}
}

func taskGet(w http.ResponseWriter, _ *http.Request, m *TaskMgr, id int) {
	/* A function to get a single task by ID
	   Функция для получения задачи по ID */
	tasks := m.GetTasks()

	for _, task := range tasks {
		if task.ID == id {
			writeTask(w, task)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	if _, err := w.Write([]byte("No item with such id exists.")); err != nil {
		log.Printf("Error writing http response: %s", err)
		http.Error(w, "Error writing http response", http.StatusInternalServerError)
	}
}
