package student

import (
	"database/sql"
	"errors"

	studentModel "github.com/Bearaujus/simple-student-app/internal/model/student"
	studentResource "github.com/Bearaujus/simple-student-app/internal/resource/student"

	uuid "github.com/satori/go.uuid"
)

type usecase struct {
	resource studentResource.ResourceItf
}

func NewUsecase(db *sql.DB) UsecaseItf {
	return &usecase{
		resource: studentResource.NewResource(db),
	}
}

func (usecase *usecase) GetStudentByID(sid string) (*studentModel.Student, error) {
	// Validate sid
	if sid == "" {
		return nil, errors.New("sid is empty")
	}

	// Call resource
	student, err := usecase.resource.GetStudentByID(sid)
	if err != nil {
		return nil, err
	}

	return student, nil
}

func (usecase *usecase) GetStudents() ([]studentModel.Student, error) {
	// Call resource
	students, err := usecase.resource.GetStudents()
	if err != nil {
		return nil, err
	}

	return students, nil
}

func (usecase *usecase) CreateStudent(student *studentModel.Student) (string, error) {
	// Validate name
	if student.Name == "" {
		return "", errors.New("name is empty")
	}

	if len(student.Name) < 3 {
		return "", errors.New("name cannot < 3 characters")
	}

	if len(student.Name) > 30 {
		return "", errors.New("name cannot > 30 characters")
	}

	// Validate age
	if student.Age < 1 {
		return "", errors.New("age cannot < 1")
	}

	// Validate grade
	if student.Grade > 4 {
		return "", errors.New("grade cannot > 4")
	}

	if student.Grade < 0 {
		return "", errors.New("grade cannot < 0")
	}

	// Generate student id
	student.SID = uuid.NewV4().String()

	// Call resource
	sid, err := usecase.resource.CreateStudent(student)
	if err != nil {
		return "", err
	}

	return sid, nil
}

func (usecase *usecase) UpdateStudent(student *studentModel.Student) (*studentModel.Student, error) {
	// Validate sid
	if student.SID == "" {
		return nil, errors.New("sid is empty")
	}

	// Validate name
	if student.Name == "" {
		return nil, errors.New("name is empty")
	}

	if len(student.Name) < 3 {
		return nil, errors.New("name cannot < 3 characters")
	}

	if len(student.Name) > 30 {
		return nil, errors.New("name cannot > 30 characters")
	}

	// Validate age
	if student.Age < 1 {
		return nil, errors.New("age cannot < 1")
	}

	// Validate grade
	if student.Grade > 4 {
		return nil, errors.New("grade cannot > 4")
	}

	if student.Grade < 0 {
		return nil, errors.New("grade cannot < 0")
	}

	// Call resource
	student, err := usecase.resource.UpdateStudent(student)
	if err != nil {
		return nil, err
	}

	return student, nil
}

func (usecase *usecase) DeleteStudent(sid string) error {
	// Validate sid
	if sid == "" {
		return errors.New("sid is empty")
	}

	// Call resource
	err := usecase.resource.DeleteStudent(sid)
	if err != nil {
		return err
	}

	return nil
}
