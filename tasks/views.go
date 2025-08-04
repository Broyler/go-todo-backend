package tasks

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

func HandleTasks(w http.ResponseWriter, r *http.Request, m *TaskMgr) {
	path := strings.TrimPrefix(r.URL.Path, "/tasks")
	path = strings.TrimPrefix(path, "/")
	if path == "" {
		CheckMethods(w, r, m)
		return
	}

	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}
	ByID(w, r, m, id)
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

func CheckMethods(w http.ResponseWriter, r *http.Request, m *TaskMgr) {
	/* A function to switch between request tasks for tasks API.
	   Функция для выбора между методами запросов для API задач. */
	switch r.Method {
	case http.MethodGet:
		get(w, r, m)
	case http.MethodPost:
		post(w, r, m)
	default:
		methodNotSupported(w, r)
	}
}

func ByID(w http.ResponseWriter, r *http.Request, m *TaskMgr, id int) {
	/* A function to switch between requests for ID specific requests
	   Функция для выбора между метода для запросов для конкретного ID */
	switch r.Method {
	case http.MethodGet:
		taskGet(w, r, m, id)
	default:
		methodNotSupported(w, r)
	}
}
