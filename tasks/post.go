package tasks

import (
	"encoding/json"
	"log"
	"net/http"
)

func getTaskPostData(r *http.Request) (Task, error) {
	/* A function for getting task data from request body.
	   Функция для получения данных о задаче из тела запроса. */
	var data Task
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func post(w http.ResponseWriter, r *http.Request, m *TaskMgr) {
	/* POST method for tasks API. Will create a new task with a given name.
	   Accepts a JSON body with "name" and "done" params.

	   Метод POST для API задач. Создаст новую задачу с переданным именем.
	   Принимает тело JSON с параметрами "name" и "done". */
	data, err := getTaskPostData(r)
	if err != nil {
		log.Printf("Error parsing JSON: %s", err)
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}
	aux := m.AddTask(data)
	w.WriteHeader(http.StatusCreated)
	writeTask(w, aux)
}
