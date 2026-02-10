package main

// This is a code wrriten by idorocodes, first go project, did this because i wanted to.
// This code explains go in the simplex way by using a school system as an analogy

import (
	"errors"
	"math/rand"
	"slices"
)

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
	SudentRusticated  bool
	StudentCourses    []Course
}

type Course struct {
	CourseId          int
	CourseCode        string
	CourseTitle       string
	DepartmentOffered Department
}

func newSchool(name string) (School, error) {
	if len(name) <= 0 {
		return School{}, errors.New("School Name not supplied !")
	}
	schoolId := rand.Int()

	var emptyDepartment []Department

	schoolData := School{SchoolName: name, SchoolId: schoolId, SchoolDepartments: emptyDepartment}

	return schoolData, nil
}

func (s *School) editSchoolName(newname string) (bool, error) {
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

	if len(currentDepartments) >= 10 {
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

func (s *School) removeDepartment(deptId int) (bool, error) {
	currentDepartments := s.SchoolDepartments

	if len(currentDepartments) == 0 {
		return false, errors.New("All departments removed !")
	}

	index := slices.IndexFunc(s.SchoolDepartments, func(d Department) bool {
		return d.DepartmentId == deptId
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

func (d *Department) addStudent(studentName string, matricNo string, age int) (bool, error) {

	if len(studentName) == 0 {
		return false, errors.New("Name not supplied!")
	}

	if age <= 0 {
		return false, errors.New("Invalid age!")
	}

	if len(matricNo) == 0 {
		return false, errors.New("Matric Number not supplied! ")
	}

	currentStudents := d.DepartmentStudents

	if len(currentStudents) > 100 {
		return false, errors.New("A Department cannot have more than 100 students")
	}

	studId := rand.Int()

	dept := *d

	var newCourse []Course

	newStudent := Student{
		StudentName:       studentName,
		StudentId:         studId,
		StudentMatricNo:   matricNo,
		StudentAge:        age,
		StudentDepartment: dept,
		SudentRusticated:  false,
		StudentCourses:    newCourse,
	}

	d.DepartmentStudents = append(d.DepartmentStudents, newStudent)

	return true, nil

}

func (d *Department) addStudents(newStudents []Student) (bool, error) {

	if len(newStudents) == 0 {
		return false, errors.New("Students not supplied!")
	}

	currentStudents := d.DepartmentStudents

	if len(currentStudents) > 100 {
		return false, errors.New("A Department cannot have more than 100 students")
	}

	sliceOfNewStudents := newStudents

	d.DepartmentStudents = append(d.DepartmentStudents, sliceOfNewStudents...)

	return true, nil

}

func (d *Department) removeStudent(studId int) (bool, error) {
	currentStudents := d.DepartmentStudents

	if len(currentStudents) == 0 {
		return false, errors.New("All students removed !")
	}

	index := slices.IndexFunc(d.DepartmentStudents, func(d Student) bool {
		return d.StudentId == studId
	})

	if index == -1 {
		return false, errors.New("student not found")
	}

	d.DepartmentStudents = slices.Delete(d.DepartmentStudents, index, index+1)

	return true, nil
}

func (d *Department) editStudent(studId int, newName string) (bool, error) {
	currentStudents := d.DepartmentStudents

	if len(currentStudents) == 0 {
		return false, errors.New("All students removed !")
	}

	index := slices.IndexFunc(d.DepartmentStudents, func(d Student) bool {
		return d.StudentId == studId
	})

	if index == -1 {
		return false, errors.New("student not found")
	}

	d.DepartmentStudents[index].StudentName = newName

	return true, nil
}

func (d *Department) rusticateStudent(studId int) (bool, error) {
	currentStudents := d.DepartmentStudents

	if len(currentStudents) == 0 {
		return false, errors.New("All students removed !")
	}

	index := slices.IndexFunc(d.DepartmentStudents, func(d Student) bool {
		return d.StudentId == studId
	})

	if index == -1 {
		return false, errors.New("student not found")
	}

	d.DepartmentStudents[index].SudentRusticated = true

	return true, nil
}

func (d *Department) reinstateStudent(studId int) (bool, error) {
	currentStudents := d.DepartmentStudents

	if len(currentStudents) == 0 {
		return false, errors.New("All students removed !")
	}

	index := slices.IndexFunc(d.DepartmentStudents, func(d Student) bool {
		return d.StudentId == studId
	})

	if index == -1 {
		return false, errors.New("student not found")
	}

	d.DepartmentStudents[index].SudentRusticated = false

	return true, nil
}

func (d *Department) changeAge(studId int, newAge int) (bool, error) {
	currentStudents := d.DepartmentStudents

	if len(currentStudents) == 0 {
		return false, errors.New("All students removed !")
	}

	index := slices.IndexFunc(d.DepartmentStudents, func(d Student) bool {
		return d.StudentId == studId
	})

	if index == -1 {
		return false, errors.New("student not found")
	}

	d.DepartmentStudents[index].StudentAge = newAge

	return true, nil
}

func (s *Student) addCourse(code string, title string) (bool, error) {

	currentCourses := s.StudentCourses
	if len(currentCourses) > 10 {
		return false, errors.New("Student cannot offer more than 10 courses!!")
	}

	if len(code) == 0 {
		return false, errors.New("Code is not supplied!")
	}

	if len(title) == 0 {
		return false, errors.New("Title is not supplied!")
	}

	rusticatedStatus := s.SudentRusticated

	if rusticatedStatus {
		return false, errors.New("Rusticated Students cannot add new courses!")
	}

	newCourseId := rand.Int()
	newCourse := Course{
		CourseCode:        code,
		CourseTitle:       title,
		CourseId:          newCourseId,
		DepartmentOffered: s.StudentDepartment,
	}

	s.StudentCourses = append(s.StudentCourses, newCourse)

	return true, nil

}

func (s *Student) removeCourse(cId int) (bool, error) {

	currentCourses := s.StudentCourses

	if len(currentCourses) == 0 {
		return false, errors.New("All courses removed!")
	}

	rusticatedStatus := s.SudentRusticated

	if rusticatedStatus {
		return false, errors.New("Rusticated Students cannot add new courses!")
	}
	index := slices.IndexFunc(s.StudentCourses, func(d Course) bool {
		return d.CourseId == cId
	})

	if index == -1 {
		return false, errors.New("course not found")
	}

	s.StudentCourses = slices.Delete(s.StudentCourses, index, index+1)

	return true, nil

}

func (s *Student) editCourse(cId int, newCode string, newTitle string) (bool, error) {

	if len(newCode) == 0 || len(newTitle) == 0 {
		return false, errors.New("Course Code or Title not Found!")
	}

	currentCourses := s.StudentCourses

	if len(currentCourses) == 0 {
		return false, errors.New("All courses removed!")
	}

	rusticatedStatus := s.SudentRusticated

	if rusticatedStatus {
		return false, errors.New("Rusticated Students cannot add new courses!")
	}
	index := slices.IndexFunc(s.StudentCourses, func(d Course) bool {
		return d.CourseId == cId
	})

	if index == -1 {
		return false, errors.New("student not found")
	}

	s.StudentCourses[index].CourseCode = newCode
	s.StudentCourses[index].CourseTitle = newTitle

	return true, nil

}

func main() {

	// I know you want to see the tests 
	// 
	// I moved it into school-test.go check it there 

}

func (s *School) getDepartmentId(deptName string) (int, error) {
	for _, d := range s.SchoolDepartments {
		if d.DepartmentName == deptName {
			return d.DepartmentId, nil
		}
	}
	return 0, errors.New("department not found")
}

func (d *Department) getStudentId(studentMatricNo string) (int, error) {
	for _, d := range d.DepartmentStudents {
		if d.StudentMatricNo == studentMatricNo {
			return d.StudentId, nil
		}
	}
	return 0, errors.New("student not found")
}

func (s *Student) getCourseId(courseCode string) (int, error) {
	for _, d := range s.StudentCourses {
		if d.CourseCode == courseCode {
			return d.CourseId, nil
		}
	}
	return 0, errors.New("course not found")
}
