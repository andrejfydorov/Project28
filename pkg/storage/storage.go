package storage

import (
	"Project28/pkg/student"
	"fmt"
)

type StudentsRepo map[string]*student.Student

func NewStudentsRepo() StudentsRepo {
	return make(map[string]*student.Student)
}

func (s StudentsRepo) Put(student *student.Student) {
	s[student.Name] = student
}

func (s StudentsRepo) Get(studentName string) (*student.Student, error) {
	ss, err := s[studentName]
	if !err {
		return nil, fmt.Errorf("no such student")
	}
	return ss, nil
}
