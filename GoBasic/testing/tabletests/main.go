package main

import "fmt"

func main() {
	fmt.Println(Sum(2, 3))
	fmt.Println(Sum(2, 4, 7))
}

func Sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
