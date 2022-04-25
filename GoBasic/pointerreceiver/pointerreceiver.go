package main

import "fmt"

/* Example to show Pointer Receiver vs Value Receiver.
Pointer Receiver = *p, value Receiver = p1

When pointer receiver is used in the method to modify the receiver's properties, changes are reflected in the original object that is used as the receiver.
For example, when 'personPointerObj' is used as the receiver to call the 'setPersonDetailsPointer' method, changes set in the method are directly reflected
in personPointerObj in the main function without returning the object from the  'setPersonDetailsPointer' method.

When value receiver is used in the method to modify the receiver's properties, changes are NOT REFLECTED in the original object that is used as the receiver.
For example, when 'personValueObj' is used as the receiver to call the 'setPersonDetailsValue' method, changes set in the method are NOT reflected
in 'personValueObj' in the main function and 'personValueObj' has to be returned from the 'setPersonDetailsValue' method in order for the changes made in the
method to be reflected in the main function also. If it is not returned and assigned to 'personValueObj' then changes won't be reflected in the main function.
*/

type person struct {
	firstName string
	lastName  string
	age       int64
}

//Pointer Receiver method to set person details. Changes are reflected normally without returning the object due to the use of pointer receiver.
func (p *person) setPersonDetailsPointer(firstName string, lastName string, age int64) {
	p.firstName = firstName
	p.lastName = lastName
	p.age = age
}

//Value Receiver method to set person details. Changes are NOT reflected in the main function and object has to be returned and assigned for the changes to be reflected in the main function
func (p1 person) setPersonDetailsValue(firstName string, lastName string, age int64) person {
	p1.firstName = firstName
	p1.lastName = lastName
	p1.age = age

	return p1
}

func main() {

	//calling the pointer receiver 'setPersonDetailsPointer' method, Changes are reflected normally without returning and assigning the object to 'personPointerObj'
	personPointerObj := &person{}
	personPointerObj.setPersonDetailsPointer("Mohammed", "Haroon", 27)
	fmt.Println(personPointerObj)

	//calling the value receiver 'setPersonDetailsValue' method, Changes are not reflected unless the object is returned and assigned to 'personValueObj'.
	personValueObj := person{}
	personValueObj = personValueObj.setPersonDetailsValue("Rasheed", "Rahman", 27)
	fmt.Println(personValueObj)

	//ERROR CASE
	//calling the value receiver 'setPersonDetailsValue' method using a POINTER TYPE *Person (i.e.personValueObj) will throw an ERROR as the VALUE RECIEVER p1 of person type cannot handle a POINTER TYPE *Person (i.e.)
	// personPointrObj1 := &person{}
	// personPointrObj1 = personPointrObj1.setPersonDetailsValue("Rasheed", "Rahman", 27)
	// fmt.Println(personPointrObj1)
}
