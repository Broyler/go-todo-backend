package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func tasksGet(w http.ResponseWriter, r *http.Request, m *TasksMgr) {
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

func tasksPost(w http.ResponseWriter, r *http.Request, m *TasksMgr) {
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

func methodNotSupported(w http.ResponseWriter, r *http.Request) {
	/* A function to write a response that the method is not supported.
	   Функция для записи в ответ информации о недоступности метода запроса. */
	_, err := w.Write([]byte("Method " + r.Method + " is not allowed"))
	if err != nil {
		log.Printf("Error writing http response: %s", err)
		http.Error(w, "Error writing http response", http.StatusInternalServerError)
		return
	}
}

func tasksAPI(w http.ResponseWriter, r *http.Request, m *TasksMgr) {
	/* A function to switch between request methods for tasks API.
	   Функция для выбора между методами запросов для API задач. */
	switch r.Method {
	case http.MethodGet:
		tasksGet(w, r, m)
	case http.MethodPost:
		tasksPost(w, r, m)
	default:
		methodNotSupported(w, r)
	}
}
