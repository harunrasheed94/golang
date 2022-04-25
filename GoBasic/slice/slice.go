package main

import "fmt"

func main() {
	x := []int{4, 5, 6, 7, 8}

	fmt.Println("Slices ")
	for i, v := range x {
		fmt.Println(i, v)
	}

	//slicing a slice

	fmt.Println("")
	fmt.Println("Slicing ")
	fmt.Println(x[1:3])

	//appending to a slice

	fmt.Println("")
	fmt.Println("Appending")
	x = append(x, 77, 88)
	fmt.Println(x)

	//appending a slice to another slice
	fmt.Println("")
	fmt.Println("Appending slice to another slice")
	y := []int{150, 250}
	x = append(x, y...) //append syntax = append(slice, variadic parameter)
	fmt.Println(x)

	//deleting from a slice

	fmt.Println("")
	fmt.Println("Deleting from a slice")
	fmt.Println(x)
	x = append(x[:2], x[4:]...)
	fmt.Println("Deleted elements at index 2 and 3")
	fmt.Println(x)

	var newSlice *[]int
	newSlice = CreateandReturnPointerSlice()
	fmt.Println(newSlice)

}

func CreateandReturnPointerSlice() *[]int {
	var sliceInt []int
	sliceInt = append(sliceInt, 11, 12, 13)
	return &sliceInt
}
