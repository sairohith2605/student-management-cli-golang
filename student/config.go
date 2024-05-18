package student

import (
	"errors"
	"github.com/oklog/ulid/v2"
)

type Student struct {
	StudentId    ulid.ULID
	Name         string
	Age          int
	Address      string
	GovernmentId string
}

func (s *Student) validate() error {
	if s.Age < 18 {
		return errors.New("student of age less than 18 does not qualify for the admission")
	}
	if s.GovernmentId == "" {
		return errors.New("the governmentId is required and cannot be empty")
	}
	if s.Address == "" {
		return errors.New("the address is required and cannot be empty")
	}
	return nil
}

func NewStudent(name string, age int, address string, governmentId string) (*Student, error) {
	newStudent := Student{
		Name:         name,
		Age:          age,
		Address:      address,
		GovernmentId: governmentId,
	}
	err := newStudent.validate()
	if err != nil {
		return nil, err
	}
	newStudent.StudentId = ulid.Make()
	_ = append(UniversityStudents, newStudent)
	return &newStudent, nil
}

var UniversityStudents []Student
