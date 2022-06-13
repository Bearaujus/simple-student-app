package student

import "net/http"

type HandlerItf interface {
	HandleGetStudentByID(w http.ResponseWriter, r *http.Request) error
	HandleGetStudents(w http.ResponseWriter, r *http.Request) error
	HandleCreateStudent(w http.ResponseWriter, r *http.Request) error
	HandleUpdateStudent(w http.ResponseWriter, r *http.Request) error
	HandleDeleteStudent(w http.ResponseWriter, r *http.Request) error
}
