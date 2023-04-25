package routes

import "github.com/gorilla/mux"
import "github.com/Msmpunk/go-server/controllers"

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", controllers.GetAllTasks).Methods("GET")
	router.HandleFunc("/api/tasks", controllers.CreateTask).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", controllers.UpdateTask).Methods("PUT")
	router.HandleFunc("/api/tasks/{id}", controllers.DeleteTask).Methods("DELETE")
	router.HandleFunc("/api/get-doc", controllers.GetDoc).Methods("GET")

	return router
}
