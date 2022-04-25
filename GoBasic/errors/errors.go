package main

import (
	"fmt"
)

type MyError struct {
	Message string
}

func (e *MyError) Error() string {
	return e.Message
}

func (e *MyError) GetInfo() string {
	return "Get info method"
}

func GetNewMyError(errorMsg string) *MyError {
	return &MyError{
		Message: errorMsg,
	}
}

func main() {
	var myErr error
	myErr = GetNewMyError("My new error")
	fmt.Println(myErr.Error())
}
