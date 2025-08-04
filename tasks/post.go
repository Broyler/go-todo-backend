package tasks

import (
	"encoding/json"
	"log"
	"net/http"
)

func getTaskPostData(w http.ResponseWriter, r *http.Request) (TaskPost, error) {
	/* A function for getting task data from request body.
	   Функция для получения данных о задаче из тела запроса. */
	var data TaskPost
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("Error encoding JSON response: %s", err)
		http.Error(w, "Error decoding request body JSON", http.StatusBadRequest)
		return data, err
	}
	return data, nil
}

func post(w http.ResponseWriter, r *http.Request, m *TaskMgr) {
	/* POST method for tasks API. Will create a new task with a given name.
	   Accepts a JSON body with "name" param.

	   Метод POST для API задач. Создаст новую задачу с переданным именем.
	   Принимает тело JSON с параметром "name". */
	data, err := getTaskPostData(w, r)
	if err != nil {
		log.Printf("Error parsing JSON: %s", err)
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}
	m.AddTask(data.Name)
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("201 Created"))
	if err != nil {
		log.Printf("Error writing http response: %s", err)
		http.Error(w, "Error writing http response", http.StatusInternalServerError)
		return
	}
}
