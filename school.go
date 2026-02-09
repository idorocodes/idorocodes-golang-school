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

func (School) newSchool(name string) (School, error) {
	if len(name) <= 0 {
		return School{}, errors.New("School Name not supplied !")
	}
	schoolId := rand.Int()

	var emptyDepartment []Department

	schoolData := School{SchoolName: name, SchoolId: schoolId, SchoolDepartments: emptyDepartment}

	return schoolData, nil
}

func (s *School) editSchool(newname string) (bool, error) {
	if len(newname) <= 0 {
		return false, errors.New("School Name not supplied !")
	}
	s.SchoolName = newname
	return true, nil
}

func (s *School) addDepartment(deptName string) (bool, error) {

	if len(deptName) == 0 {
		return false, errors.New("department name not supplied!")
	}
	currentDepartments := s.SchoolDepartments
	
	if len(currentDepartments) > 10 {
		return false, errors.New("A School cannot have more than 10 departments")
	}
	
	deptId := rand.Int()

	var emptyStudent []Student

	newDept := Department{
		DepartmentName:     deptName,
		DepartmentId:       deptId,
		DepartmentStudents: emptyStudent,
	}
	s.SchoolDepartments = append(s.SchoolDepartments, newDept)

	return true, nil
}

func (s *School) removeDepartment(preferedDepartment Department) (bool, error) {
	currentDepartments := s.SchoolDepartments

	if len(currentDepartments) == 0 {
		return false, errors.New("All departments removed !")
	}

	index := slices.IndexFunc(s.SchoolDepartments, func(d Department) bool {
		return d.DepartmentId == preferedDepartment.DepartmentId
	})

	if index == -1 {
		return false, errors.New("department not found")
	}

	s.SchoolDepartments = slices.Delete(s.SchoolDepartments, index, index+1)

	return true, nil
}

func (s *School) editDepartment(departmentId int, newName string) (bool, error) {
	currentDepartments := s.SchoolDepartments

	if len(currentDepartments) == 0 {
		return false, errors.New("All departments removed !")
	}

	if len(newName) == 0 {
		return false, errors.New("Name not supplied! !")
	}
	index := slices.IndexFunc(s.SchoolDepartments, func(d Department) bool {
		return d.DepartmentId == departmentId
	})

	if index == -1 {
		return false, errors.New("department not found")
	}

	s.SchoolDepartments[index].DepartmentName = newName

	return true, nil
}

func (d *Department) addStudent(studentName string, matricNo string, age int) (bool,error) {
	
	if len(studentName) == 0 {
		return false, errors.New("Name not supplied!")
	}
	
	if age <= 0 {
			return false, errors.New("Invalid age!")
	}
	
	if len(matricNo)  == 0 {
			return false, errors.New("Matric Number not supplied! ")
	}
	
	currentStudents := d.DepartmentStudents
	
	if len(currentStudents) > 100 {
		return false, errors.New("A Department cannot have more than 100 students")
	}
	
	studId := rand.Int()
	
	dept := *d
	
	var newCourse  []Course
	
	newStudent := Student{
		StudentName: studentName,
		StudentId: studId,
		StudentMatricNo:matricNo,
		StudentAge: age,
		StudentDepartment: dept,
		SudentRustsicated: false,
		StudentFavCourse: newCourse,
	}
	
	d.DepartmentStudents = append(d.DepartmentStudents, newStudent)
	
	return true, nil
	
}

func main() {

	fmt.Println("Hello World!, this is my school!")
}
