package student

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	studentModel "github.com/Bearaujus/simple-student-app/internal/model/student"
	studentUsecase "github.com/Bearaujus/simple-student-app/internal/usecase/student"
	util "github.com/Bearaujus/simple-student-app/util"
	"github.com/go-chi/chi"
)

type handler struct {
	usecase studentUsecase.UsecaseItf
}

func NewHandler(db *sql.DB) HandlerItf {
	return &handler{
		usecase: studentUsecase.NewUsecase(db),
	}
}

func (handler *handler) HandleGetStudentByID(w http.ResponseWriter, r *http.Request) error {
	// Parse sid
	sid := chi.URLParam(r, "sid")

	// Call usecase
	student, err := handler.usecase.GetStudentByID(sid)
	if err != nil {
		return err
	}

	// Generate return json
	resp, err := util.ParseResponseToJSON(true, nil, student)
	if err != nil {
		return err
	}

	// Write result
	w.Write([]byte(resp))
	return nil
}

func (handler *handler) HandleGetStudents(w http.ResponseWriter, r *http.Request) error {
	// Call usecase
	students, err := handler.usecase.GetStudents()
	if err != nil {
		return err
	}

	// Generate return json
	resp, err := util.ParseResponseToJSON(true, nil, students)
	if err != nil {
		return err
	}

	// Write result
	w.Write([]byte(resp))
	return nil
}

func (handler *handler) HandleCreateStudent(w http.ResponseWriter, r *http.Request) error {
	// Parse name
	name := r.URL.Query().Get("name")

	// Parse age
	age, err := strconv.Atoi(r.URL.Query().Get("age"))
	if err != nil {
		return errors.New("cannot parse age")
	}

	// Parse grade
	grade, err := strconv.Atoi(r.URL.Query().Get("grade"))
	if err != nil {
		return errors.New("cannot parse grade")
	}

	// Call usecase
	student := studentModel.Student{
		Name:  name,
		Age:   age,
		Grade: grade,
	}

	sid, err := handler.usecase.CreateStudent(&student)
	if err != nil {
		return err
	}

	// Generate return json
	resp, err := util.ParseResponseToJSON(true, nil, studentModel.Student{
		SID: sid,
	})
	if err != nil {
		return err
	}

	// Write result
	w.Write([]byte(resp))
	return nil
}

func (handler *handler) HandleUpdateStudent(w http.ResponseWriter, r *http.Request) error {
	// Parse sid
	sid := chi.URLParam(r, "sid")

	// Parse name
	name := r.URL.Query().Get("name")

	// Parse age
	age, err := strconv.Atoi(r.URL.Query().Get("age"))
	if err != nil {
		return errors.New("cannot parse age")
	}

	// Parse grade
	grade, err := strconv.Atoi(r.URL.Query().Get("grade"))
	if err != nil {
		return errors.New("cannot parse grade")
	}

	// Call usecase
	student := studentModel.Student{
		SID:   sid,
		Name:  name,
		Age:   age,
		Grade: grade,
	}

	nStudent, err := handler.usecase.UpdateStudent(&student)
	if err != nil {
		return err
	}

	// Generate return json
	resp, err := util.ParseResponseToJSON(true, nil, nStudent)
	if err != nil {
		return err
	}

	// Write result
	w.Write([]byte(resp))
	return nil
}

func (handler *handler) HandleDeleteStudent(w http.ResponseWriter, r *http.Request) error {
	// Parse sid
	sid := chi.URLParam(r, "sid")

	// Call usecase
	err := handler.usecase.DeleteStudent(sid)
	if err != nil {
		return err
	}

	// Generate return json
	resp, err := util.ParseResponseToJSON(true, nil, nil)
	if err != nil {
		return err
	}

	// Write result
	w.Write([]byte(resp))
	return nil
}
