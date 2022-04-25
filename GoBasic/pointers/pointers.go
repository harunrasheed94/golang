package main

import "fmt"

var print = fmt.Println

/*
Go is ONLY PASS BY VALUE.
In the below example I am PASSING AN ADDRESS (of 'x') by VALUE to the pointer 'y' so that the POINTER 'y' in the foo function also POINTS to the SAME VALUE
POINTED by x. And then in foo() I AM CHANGING the VALUE THAT IS BEING STORED IN 'x' by SETTING *y = 84.

Hence BY PASSING ADDRESS of a VARIABLE BY VALUE to a POINTER ('*y'), I can change the value of the variable ('x' here) WHOSE ADDRESS
is being passed as a VALUE to the pointer and is pointed by the pointer 'y'.

BUT ADDRESS of 'x' WHICH IS BEING PASSED as a VALUE if MODIFIED in foo (i.e set to some other variable's address) WILL NOT BE REFLECTED in the MAIN FUNCTION.
For example if we change the address of y which initially holds the address of 'x' being passed from main, by removing the commented lines and
making it equal to &b (i.e. y = &b) then only '*y' pointer will stop pointing to x and instead point to 'b' but 'x' will still have the old address and
the old value and its value will not be modified now as '*y' is pointing to 'b' now and not 'x' as it was pointing initially.
HENCE golang ONLY has CALL BY VALUE.

*/
func main() {
	x := 42
	print(x)
	foo(&x)
	print(x)
}

func foo(y *int) {
	// b := 155
	// y = &b
	*y = 84
}
