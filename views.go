package main

import (
	"log"
	"net/http"
)

func tasksGet(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Get method called"))
	if err != nil {
		log.Fatal("Error writing http response", err)
	}
}

func tasksDelete(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Delete method called"))
	if err != nil {
		log.Fatal("Error writing http response", err)
	}
}

func methodNotSupported(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Method " + r.Method + " is not allowed"))
	if err != nil {
		log.Fatal("Error writing http response", err)
	}
}

func tasksAPI(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tasksGet(w, r)
	case http.MethodDelete:
		tasksDelete(w, r)
	default:
		methodNotSupported(w, r)
	}
}
