package main

import (
	"src/management/admission"
)

func main() {
	println("University Student Management Console")
	var admittedStudentInfo, err = admission.RunAdmissionForNewStudent(
		"Alex Hilbert",
		27,
		"1-8/A, 5th Avenue, San Fransisco",
		"UAN8819921")
	if err != nil {
		println("There was an error processing the student's admission:", err.Error())
	} else {
		println("Admission successful. Admitted student ID: ", admittedStudentInfo.StudentId.String())
	}
}
