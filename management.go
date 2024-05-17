package main

import (
	"src/management/admission"
	"src/management/student"
)

func main() {
	println("University Student Management Console")
	var newStudent = student.Student{
		Name:         "Alex Hilbert",
		Age:          27,
		Address:      "1-8/A, 5th Avenue, San Fransisco",
		GovernmentId: "UAN8819921",
	}
	var admittedStudentInfo = admission.RunAdmissionForNewStudent(newStudent)
	println("Admission successful. Admitted student info: ", admittedStudentInfo.StudentId.String())
}
