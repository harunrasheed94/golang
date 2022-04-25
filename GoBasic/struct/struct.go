package main

import(
	"fmt"
)

type person struct{
	first string
	last string
	age int
}

type secretagent struct{
	person
	ltk bool
}

func main(){
	p1 := person{
		first: "James",
		last: "Bond",
		age: 35,
	}

	sa := secretagent{
		person: p1,
		ltk: true,
	}

	fmt.Println(p1.first)

	fmt.Println(sa.first, sa.last)
}