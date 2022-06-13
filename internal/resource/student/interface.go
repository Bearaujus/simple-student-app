package student

import (
	studentModel "github.com/Bearaujus/simple-student-app/internal/model/student"
)

type ResourceItf interface {
	GetStudentByID(string) (*studentModel.Student, error)
	GetStudents() ([]studentModel.Student, error)
	CreateStudent(*studentModel.Student) (string, error)
	UpdateStudent(*studentModel.Student) (*studentModel.Student, error)
	DeleteStudent(string) error
}
