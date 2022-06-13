package student

import (
	"database/sql"

	studentModel "github.com/Bearaujus/simple-student-app/internal/model/student"
)

type resource struct {
	db *sql.DB
}

func NewResource(db *sql.DB) ResourceItf {
	return &resource{
		db: db,
	}
}

func (resource *resource) GetStudentByID(sid string) (*studentModel.Student, error) {
	var student studentModel.Student
	if err := resource.db.QueryRow(string(GetStudentByID), sid).Scan(&student.SID, &student.Name, &student.Age, &student.Grade); err != nil {
		return nil, err
	}
	return &student, nil
}

func (resource *resource) GetStudents() ([]studentModel.Student, error) {
	rows, err := resource.db.Query(string(GetStudents))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []studentModel.Student
	for rows.Next() {
		var student studentModel.Student
		if err := rows.Scan(&student.SID, &student.Name, &student.Age, &student.Grade); err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func (resource *resource) CreateStudent(student *studentModel.Student) (string, error) {
	stmt, err := resource.db.Prepare(string(CreateStudent))
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(student.SID, student.Name, student.Age, student.Grade)
	if err != nil {
		return "", err
	}

	return student.SID, nil
}

func (resource *resource) UpdateStudent(student *studentModel.Student) (*studentModel.Student, error) {
	_, err := resource.GetStudentByID(student.SID)
	if err != nil {
		return nil, err
	}

	stmt, err := resource.db.Prepare(string(UpdateStudent))
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(student.SID, student.Name, student.Age, student.Grade, student.SID)
	if err != nil {
		return nil, err
	}

	return student, nil
}

func (resource *resource) DeleteStudent(sid string) error {
	_, err := resource.GetStudentByID(sid)
	if err != nil {
		return err
	}

	stmt, err := resource.db.Prepare(string(DeleteStudent))
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(sid)
	if err != nil {
		return err
	}

	return nil
}
