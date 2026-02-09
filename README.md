 # Go School System (Learning Go by Analogy)

A beginner-friendly Go project that explains core Go concepts using a school system analogy.

This project models a real-world school structure (School → Departments → Students → Courses) to demonstrate how Go handles structs, methods, interfaces, slices, pointers, errors, and basic data modeling.

Built as a hands-on learning exercise while practicing Go fundamentals.


## Core Concepts Demonstrated

This project intentionally touches the following Go concepts:

- Structs and nested structs
- Methods attached to structs
- Pointer receivers for mutation
- Error handling with `error`
- Slice operations (`append`, `slices.IndexFunc`, `slices.Delete`)
- Interfaces (basic introduction)
- Package imports and organization
- Random ID generation
- Defensive programming (validation + limits)


## Data Model Overview

### 1. School

Represents a school entity.

```go
type School struct {
    SchoolName        string
    SchoolId          int
    SchoolDepartments []Department
}
```

**Responsibilities:**

- Create a new school
- Edit school name
- Add departments (max 10)
- Remove departments
- Edit departments


### 2. Department

Represents an academic department.

```go
type Department struct {
    DepartmentName     string
    DepartmentId       int
    DepartmentStudents []Student
}
```

**Responsibilities:**

- Add a student (single or batch)
- Remove students
- Edit student details
- Rusticate / reinstate students
- Enforce max student limit (100)


### 3. Student

Represents a student.

```go
type Student struct {
    StudentName       string
    StudentId         int
    StudentMatricNo   string
    StudentAge        int
    StudentDepartment Department
    SudentRustsicated bool
    StudentFavCourse  []Course
}
```

**Features:**

- Belongs to a department
- Can be rusticated or reinstated
- Stores basic student metadata

### 4. Course

Represents a course offered by departments.

```go
type Course struct {
    CourseCode        string
    CourseTitle       string
    DepartmentOffered []Department
}
```

## Key Operations Implemented

### School-Level Operations

- `newSchool(name)`
- `editSchool(newName)`
- `addDepartment(deptName)`
- `removeDepartment(deptId)`
- `editDepartment(deptId, newName)`


### Department-Level Operations

- `addStudent(name, matricNo, age)`
- `addStudents([]Student)`
- `removeStudent(studentId)`
- `editStudent(studentId, newName)`
- `changeAge(studentId, newAge)`
- `rusticateStudent(studentId)`
- `reinstateStudent(studentId)`

All operations:

- Validate inputs
- Return `(bool, error)`
- Fail safely with descriptive errors

## Why a School Analogy?

Because **abstractions stick better when they map to reality**.

| Go Concept | School Analogy |
|------------|----------------|
| Struct | Real-world object |
| Slice | Group of entities |
| Method | Action performed |
| Pointer receiver | Mutating state |
| Error handling | Real-life validation |
| Interface | Behavior contract |

This makes Go concepts easier to visualize, remember, and extend.


## How to Run

```bash
go run main.go
```


## Author

**idorocodes**  
First Go project. Built out of curiosity and discipline.
