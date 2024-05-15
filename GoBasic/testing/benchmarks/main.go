package main

import "fmt"

func main() {
	fmt.Println(Sum(5, 4, 3, 1, 1))
}

func Sum(x ...int) int {
	total := 0

	for _, v := range x {
		total += v
	}
	return total
}
