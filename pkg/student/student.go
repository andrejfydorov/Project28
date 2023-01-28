package student

import (
	"fmt"
	"strconv"
)

type Student struct {
	Name  string
	Age   int
	Grade int
}

func NewStudent(name string, age int, grade int) *Student {
	return &Student{
		Name:  name,
		Age:   age,
		Grade: grade,
	}
}

func (s Student) StudentToString() string {
	return fmt.Sprintf("Студент: %v, возраст: %v, класс: %v", s.Name, strconv.Itoa(s.Age), strconv.Itoa(s.Grade))
}
