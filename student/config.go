package student

import "github.com/oklog/ulid/v2"

type Student struct {
	StudentId    ulid.ULID
	Name         string
	Age          int
	Address      string
	GovernmentId string
}

var UniversityStudents []Student
