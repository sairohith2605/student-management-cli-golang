package admission

import (
	"github.com/oklog/ulid/v2"
	"src/management/student"
)

func RunAdmissionForNewStudent(newStudent student.Student) *student.Student {
	newStudentId := ulid.Make()
	newStudent.StudentId = newStudentId
	_ = append(student.UniversityStudents, newStudent)
	return &newStudent
}
