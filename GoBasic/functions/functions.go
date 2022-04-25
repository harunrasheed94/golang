package main

import (
	"fmt"
)

/*
Functions.
Variadic Parameters.
Unfurling a slice
*/
func main() {
	foo("Rasheed", 1, 2, 3, 4, 5, 6, 7, 8, 9)

	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	//foo(xi) will throw an error bcoz cannot pass a slice as variadic parameter (vp)

	//convert slice to vp i.e. UNFURLING the SLICE (xi) to VP
	foo("Rasheed", xi...)

	foo("Rasheed")
}

/*func foo (x ...int, s string) will throw an error for foo("Rasheed"). Thus
VP ALWAYS has to be the LAST PARAMETER as IT CAN BE EMPTY but OTHER PARAMETERS
that come before it CAN'T be EMPTY
*/

func foo(s string, x ...int) {

	sum := 0

	//vp is converted to slice in foo automatically but when we are
	//passing as argument conversion will have to be done by us

	fmt.Println(x)

	for _, v := range x {
		sum += v
	}

	fmt.Println("Hi ", s, "sum is ", sum)
}
