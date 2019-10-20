package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/faozimipa/monggo/db"
	"github.com/faozimipa/monggo/models"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

func handleError(err error, message string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf(message, err)))
}

/*GetAllStudents all
 */
func GetAllStudents(w http.ResponseWriter, req *http.Request) {
	rs, err := db.GetAllStudent()
	if err != nil {
		handleError(err, "failed to load alll students", w)
		return
	}

	bs, err := json.Marshal(rs)
	if err != nil {
		handleError(err, "Failed to marshal data", w)
		return
	}
	w.Write(bs)
}

// GetStudent returns a single database item matching given ID parameter.
func GetStudent(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	rs, err := db.GetOne(id)
	if err != nil {
		handleError(err, "Failed to read database: %v", w)
		return
	}

	bs, err := json.Marshal(rs)
	if err != nil {
		handleError(err, "Failed to marshal data: %v", w)
		return
	}

	w.Write(bs)
}

/*PostStudent poststudent
 */
func PostStudent(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var item models.Student
	item.UUID = uuid.NewV4()
	err := decoder.Decode(&item)
	if err != nil {
		panic(err)
	}
	log.Println(item.Name)

	err = db.Save(item)
	if err != nil {
		handleError(err, "Failed to save data: %v", w)
		return
	}

	w.Write([]byte("OK"))

}

/*UpdateStudent update student
 */
func UpdateStudent(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	decoder := json.NewDecoder(req.Body)
	var item models.Student

	err := decoder.Decode(&item)
	if err != nil {
		panic(err)
	}

	err = db.Update(id, item)

	if err != nil {
		handleError(err, "Failed to save data: %v", w)
		return
	}

	w.Write([]byte("OK"))

}

/*DeleteStudent delete student
 */
func DeleteStudent(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	err := db.Delete(id)

	if err != nil {
		handleError(err, "Failed to delete data: %v", w)
		return
	}

	w.Write([]byte("OK"))
}
