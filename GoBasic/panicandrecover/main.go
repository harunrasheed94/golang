package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f()")
}

func f() {

	fmt.Println("Calling g()")
	g(0)
	fmt.Println("Returned normally from g()")
}

func g(i int) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered in g()", r)
		}
	}()
	for i = 0; i < 5; i++ {
		fmt.Println("Printing in g()", i)
		if i > 3 {
			fmt.Println("Panicking!")
			panic(fmt.Sprintf("%v", i))
		}
	}
	fmt.Println("continue in g()")
}
