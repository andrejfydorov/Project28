package main

import (
	"Project28/pkg/storage"
	"Project28/pkg/student"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	students := storage.NewStudentsRepo()

	var name string
	var age int
	var grade int

	var input string
	fmt.Print("Введите студента в формате: имя возраст класс\n")

	buf := bufio.NewScanner(os.Stdin)

	for buf.Scan() != false {
		input = buf.Text()
		inputStrings := strings.Split(input, " ")
		//fmt.Println(inputStrings)

		name = inputStrings[0]
		age, _ = strconv.Atoi(inputStrings[1])
		grade, _ = strconv.Atoi(inputStrings[2])

		s := student.NewStudent(name, age, grade)
		students.Put(s)
		fmt.Print("Введите студента в формате: имя возраст класс\n")
	}

	for _, s := range students {
		fmt.Println(s.StudentToString())
	}

}
