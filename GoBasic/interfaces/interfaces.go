package main

import (
	"fmt"
)

type professor struct {
	name   string
	dept   string
	salary int
}

type student struct {
	name  string
	dept  string
	batch string
}

type human interface {
	speak()
}

//student implements speak method. Hence student is of type human
func (s student) speak() {
	fmt.Println("Student can speak")
	fmt.Println("Name = ", s.name, " Department = ", s.dept, " Batch = ", s.batch)
}

//professor implements speak method. Hence professor is of type human
func (p professor) speak() {
	fmt.Println("Professor can speak")
	fmt.Println("Name = ", p.name, " Department = ", p.dept, " Salary = ", p.salary)
}

/*
Since both professor and student implement human interface, therefore both the student and professor is
of type human
*/

func main() {
	fmt.Println("Interfaces demo")

	s1 := student{
		name:  "Haroon Rasheed",
		dept:  "CISE",
		batch: "Fall 2019",
	}

	p1 := professor{
		name:   "Tom Shrimpton",
		dept:   "CISE",
		salary: 300000,
	}

	s1.speak()
	p1.speak()

}
