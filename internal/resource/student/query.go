package student

type query string

const (
	StudentTable    query = "CREATE TABLE IF NOT EXISTS student (sid TEXT PRIMARY KEY NOT NULL, name TEXT NOT NULL, age INTEGER NOT NULL, grade INTEGER NOT NULL);"
	StudentTruncate query = "DELETE FROM student;"

	GetStudentByID query = "SELECT sid, name, age, grade FROM student WHERE sid=?;"
	GetStudents    query = "SELECT sid, name, age, grade FROM student;"
	CreateStudent  query = "INSERT INTO student (sid, name, age, grade) VALUES (?, ?, ?, ?);"
	UpdateStudent  query = "UPDATE student SET sid=?, name=?, age=?, grade=? WHERE sid=?;"
	DeleteStudent  query = "DELETE FROM student WHERE sid=?;"
)
