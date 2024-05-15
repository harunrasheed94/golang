package main

import (
	"fmt"
	"reflect"
)

type user struct {
	name string
	age  int
}

/*Below program shows how reflection is used to find the type and value of a variable. Especially when the variable is of 'struct' kind, 'reflect.ValueOf()'
and 'reflect.TypeOf()' is used to find the number of fields in the struct and find each field's name and value.*/
func FindInfo(i interface{}) {

	fmt.Println("Kind of variable 'u' is = ", reflect.ValueOf(i).Kind())
	fmt.Println("Type of variable 'u' is = ", reflect.TypeOf(i).Name())

	if reflect.TypeOf(i).Kind() == reflect.Struct {
		valueInfo := reflect.ValueOf(i)
		typeInfo := reflect.TypeOf(i)
		numFields := valueInfo.NumField()
		fmt.Println("Num of fields in struct is = ", numFields)
		for k := 0; k < numFields; k++ {
			fmt.Println("Field Number = ", k)
			fmt.Println("Field Name = ", typeInfo.Field(k).Name)
			fmt.Println("Field Value = ", valueInfo.Field(k))
		}

	}
}

func main() {
	u := user{
		name: "Haroon",
		age:  27,
	}

	FindInfo(u)
}
