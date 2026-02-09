package main

// This is a code wrriten by idorocodes, first go project, did this because i wanted to.
// This code explains go in the simplex way by using a school system as an analogy

import (
	"errors"
	"fmt"
	"math/rand"
	"slices"
)

// interface
var s interface {
	remove() bool
	add() bool
}

// Data Structures
type School struct {
	SchoolName        string
	SchoolId          int
	SchoolDepartments []Department
}

type Department struct {
	DepartmentName     string
	DepartmentId       int
	DepartmentStudents []Student
}

type Student struct {
	StudentName       string
	StudentId         int
	StudentMatricNo   string
	StudentAge        int
	StudentDepartment Department
	SudentRustsicated bool
	StudentFavCourse  []Course
}

type Course struct {
	CourseCode        string
	CourseTitle       string
	DepartmentOffered []Department
}

func (School) addSchool(name string) (School, error) {
	if len(name) <= 0 {
		return School{}, errors.New("School Name not supplied !")
	}
	schoolId := rand.Int()

	var emptyDepartment []Department

	schoolData := School{SchoolName: name, SchoolId: schoolId, SchoolDepartments: emptyDepartment}

	return schoolData, nil
}

func (s *School) addDepartment(newDepartment Department) (bool, error) {
	currentDepartment := s.SchoolDepartments
	if len(currentDepartment) > 10 {
		return false, errors.New("A School cannot have more than 10 departments")
	}
	currentDepartment = append(currentDepartment, newDepartment)

	return true, nil
}

func (s *School) removeDepartment(preferedDepartment Department) (bool, error) {
	currentDepartments := s.SchoolDepartments
	if len(currentDepartments) < 0 {
		return false, errors.New("All departments removed !")
	}

	index := slices.Index(s.SchoolDepartments, preferedDepartment)

	s.SchoolDepartments = slices.Delete(s.SchoolDepartments, index, index+1)

	return true, nil
}

func main() {

	fmt.Println("Hello World!, this is my school!")
}
