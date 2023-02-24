package router

import (
	"github.com/gorilla/mux"
	"github.com/krishnalagad/golang-react-todo/middlewares"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	
	router.HandleFunc("/api/task", middlewares.GetAllTasks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/tasks", middlewares.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/tasks/{id}", middlewares.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undotask/{id}", middlewares.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deletetask/{id}", middlewares.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/deleteAllTasks", middlewares.DeleteAllTasks).Methods("DELETE", "OPTIONS")

	return router
}
