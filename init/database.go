package init

import (
	"database/sql"

	studentResource "github.com/Bearaujus/simple-student-app/internal/resource/student"
)

func InitDatabase() (*sql.DB, error) {
	// Initialize database
	db, err := sql.Open("sqlite3", "./.db")
	if err != nil {
		return nil, err
	}

	// Initialize table student
	_, err = db.Exec(string(studentResource.StudentTable))
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(string(studentResource.StudentTruncate))
	if err != nil {
		return nil, err
	}

	return db, nil
}
