package main

import "fmt"

/*callback = is passing function as an argument to another function which will execute the 'callback' function in its body.
In the below example, function 'sum' is passed as a callback to 'calcevensum' and is used to calculate the sum of even numbers. Callback function 'sum' is executed
by 'calcevensum' function in its body once the even numbers from the mainSlice have been found out. Then the callback function 'sum' return the sum of the
even numbers.
*/

/*
1) Passing the 'sum' func as a callback to 'calcevensum' function along with the slice with both even and odd numbers.
2) CALLBACK function will be executed by the 'calcevensum' function after it has found the even numbers
*/

func main() {
	allNumSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(calcevensum(sum, allNumSlice...))
}

//function to calculate the sum
func sum(x ...int) (total int) {

	for _, v := range x {
		total += v
	}
	return total
}

/*
function that accepts 'sum' function as a callback (i.e. an argument) and calls the callback once it has found out the even numbers
1) find all the even numbers of the slice from the main function (i.e. 'mainSlice') and append it to evenSlice
2) pass the even slice to the callback function 'sum' which is stored in 'f' and 'sum' will return the sum of even numbers. Finally return the evensum
*/
func calcevensum(f func(x ...int) int, mainSlice ...int) (evensum int) {

	var evenSlice []int

	for _, v := range mainSlice {
		if v%2 == 0 {
			evenSlice = append(evenSlice, v)
		}
	}

	evensum = f(evenSlice...)
	return evensum
}
