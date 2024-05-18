package admission

import (
	"src/management/student"
)

func RunAdmissionForNewStudent(name string, age int, address string, governmentId string) (*student.Student, error) {
	newStudent, err := student.NewStudent(name, age, address, governmentId)
	if err != nil {
		return nil, err
	}
	return newStudent, nil
}
