package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func tasksGet(w http.ResponseWriter, r *http.Request, tasks *[]Task) {
	/* GET method for tasks API.
	   Will return a paginated JSON response with tasks.

	   Метод GET для API задач.
	   Вернет JSON задач с пагинацией. */

	taskStrings := make([]string, 0)
	for _, task := range *tasks {
		taskStrings = append(taskStrings, task.JSON())
	}
	tasksJSON := strings.Join(taskStrings, ", ")
	res := fmt.Sprintf("{\"count\": %d, \"results\": [%s]}", len(*tasks), tasksJSON)
	w.Header().Add("Content-Type", "application/json")
	_, err := w.Write([]byte(res))
	if err != nil {
		http.Error(w, "Error writing http response", http.StatusBadRequest)
		panic(err)
	}
}

func getTaskPostData(w http.ResponseWriter, r *http.Request) (TaskPost, error) {
	/* A function for getting task data from request body.
	   Функция для получения данных о задаче из тела запроса. */
	var data TaskPost
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Error decoding request body JSON", http.StatusBadRequest)
		return data, err
	}
	return data, nil
}

func tasksPost(w http.ResponseWriter, r *http.Request, mgr *TasksMgr) {
	/* POST method for tasks API. Will create a new task with a given name.
	   Accepts a JSON body with "name" param.

	   Метод POST для API задач. Создаст новую задачу с переданным именем.
	   Принимает тело JSON с параметром "name". */
	data, err := getTaskPostData(w, r)
	if err != nil {
		return
	}
	appendTask(mgr, data.Name)
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("201 Created"))
	if err != nil {
		http.Error(w, "Error writing http response", http.StatusBadRequest)
		panic(err)
	}
}

func methodNotSupported(w http.ResponseWriter, r *http.Request) {
	/* A function to write a response that the method is not supported.
	   Функция для записи в ответ информации о недоступности метода запроса. */
	_, err := w.Write([]byte("Method " + r.Method + " is not allowed"))
	if err != nil {
		panic(err)
	}
}

func tasksAPI(w http.ResponseWriter, r *http.Request, mgr *TasksMgr) {
	/* A function to switch between request methods for tasks API.
	   Функция для выбора между методами запросов для API задач. */
	switch r.Method {
	case http.MethodGet:
		tasksGet(w, r, (*mgr).Tasks)
	case http.MethodPost:
		tasksPost(w, r, mgr)
	default:
		methodNotSupported(w, r)
	}
}
