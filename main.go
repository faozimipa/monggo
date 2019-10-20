package main

import (
	"net/http"

	"github.com/faozimipa/monggo/api"
	"github.com/gorilla/mux"
)

/* main func
 */
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/students", api.GetAllStudents).Methods("GET")
	r.HandleFunc("/api/students/{id}", api.GetStudent).Methods("GET")
	r.HandleFunc("/api/students", api.PostStudent).Methods("POST")
	r.HandleFunc("/api/students/edit/{id}", api.UpdateStudent).Methods("POST")
	r.HandleFunc("/api/students/delete/{id}", api.DeleteStudent).Methods("DELETE")
	http.ListenAndServe(":3200", r)
}
