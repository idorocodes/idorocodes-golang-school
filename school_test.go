package main

import "testing"

func TestNewSchool(t *testing.T) {
	school, err := newSchool("Idoro School Of Golang")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if school.SchoolName != "Idoro School Of Golang" {
		t.Fatalf("school name mismatch")
	}

	if school.SchoolId == 0 {
		t.Fatalf("school id not generated")
	}
}

func TestEditSchoolName(t *testing.T) {
	school, _ := newSchool("Old Name")

	ok, err := school.editSchoolName("New Name")
	if err != nil || !ok {
		t.Fatalf("expected edit to succeed")
	}

	if school.SchoolName != "New Name" {
		t.Fatalf("school name not updated")
	}
}


func TestAddAndRemoveDepartment(t *testing.T) {
	school, _ := newSchool("Test School")

	_, err := school.addDepartment("Computer Science")
	if err != nil {
		t.Fatalf("failed to add department")
	}

	deptID, err := school.getDepartmentId("Computer Science")
	if err != nil {
		t.Fatalf("department not found")
	}

	ok, err := school.removeDepartment(deptID)
	if err != nil || !ok {
		t.Fatalf("failed to remove department")
	}

	if len(school.SchoolDepartments) != 0 {
		t.Fatalf("department not removed")
	}
}


func TestEditDepartment(t *testing.T) {
	school, _ := newSchool("Test School")
	school.addDepartment("cs")

	deptID, _ := school.getDepartmentId("cs")

	ok, err := school.editDepartment(deptID, "COMPUTER SCIENCE")
	if err != nil || !ok {
		t.Fatalf("failed to edit department")
	}

	if school.SchoolDepartments[0].DepartmentName != "COMPUTER SCIENCE" {
		t.Fatalf("department name not updated")
	}
}

func TestStudentLifecycle(t *testing.T) {
	school, _ := newSchool("Test School")
	school.addDepartment("CS")

	dept := &school.SchoolDepartments[0]

	_, err := dept.addStudent("Idoro", "CSC001", 20)
	if err != nil {
		t.Fatalf("failed to add student")
	}

	studID, err := dept.getStudentId("CSC001")
	if err != nil {
		t.Fatalf("student not found")
	}

	_, err = dept.rusticateStudent(studID)
	if err != nil {
		t.Fatalf("failed to rusticate student")
	}

	if !dept.DepartmentStudents[0].SudentRusticated {
		t.Fatalf("student should be rusticated")
	}
}


func TestRusticatedStudentCannotAddCourse(t *testing.T) {
	school, _ := newSchool("Test School")
	school.addDepartment("CS")

	dept := &school.SchoolDepartments[0]
	dept.addStudent("Idoro", "CSC001", 20)

	student := &dept.DepartmentStudents[0]
	dept.rusticateStudent(student.StudentId)

	_, err := student.addCourse("CSC301", "Algorithms")
	if err == nil {
		t.Fatalf("expected error when rusticated student adds course")
	}
}