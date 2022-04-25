package main 

import(
	"fmt"
	"sort"
)
/*
To custom sort a struct by using the 'Sort' Method from 'sort' package, we have to IMPLEMENT the 'Interface' interface by 
IMPLEMENTING (OVERRIDING) "Len", "Swap" and "Less" methods of 'Interface' interface since 'Sort' METHOD USES 'Len' method to know the length 'n' and makes 
'n log n' calls to 'Less' and 'Swap' methods of the 'Interface' interface.

(Similar to in Java where to implement the comparator we have to override the "compare" method.)

Steps:
1) To (custom) sort a slice of persons by age (person = user defined struct), we have to define a type of slice of persons called as 'ByAge'.

2) In order to sort the slice of persons (i.e.'ByAge' type) with 'Age' by using the in-built 'Sort(data Interface)' Method from sort package, 
We have to implement the 'Len','Swap' and 'Less' Methods of 'Interface' interface since 'Sort' Method accepts only a parameter of 'Interface' type as 
it uses 'Interface' method 'Len' to know the length 'n' of the slice of struct and makes 'n log n' calls to 'Less' and 'Swap' methods of the 'Interface' interface. 
Also in the 'LESS' METHOD WE DEFINE USING WHICH FIELD OF STRUCT WE HAVE TO CUSTOM SORT. 

3) Convert the slice of persons that you want to sort, to 'ByAge' type which is also 'slice of person' type and also 'ByAge' type implements 'Interface' interface 
(by overriding its 3 methods) and hence can be passed as a parameter to the 'Sort' Method from 'sort' package.
				personsByAge := ByAge(persons)

4) Sort the 'personsByAge' ByAge type (i.e. slice of person) by calling 'sort.Sort(personsByAge)'. Now 'personsByAge' variable will have the slice of persons sorted by 
Age.

('Sort' Method from sort package makes use of the utility functions 'Len', 'Swap' and 'Less' to sort the struct as it helps it to determine based on which field
the sorting has to be done and the swap needs to take place or not. Hence 'Sort' Method uses 'Merge sort' algorithm to sort the slice of struct and uses the utility 
functions 'Len','Less' and 'Swap' to determine length of slice and decide based on which field the sorting has to be done and whether swapping is needed or not)
*/
type Person struct{
	Name string
	Age int
}

func (p Person)String() string{
	return fmt.Sprintf("Name = %s Age = %d,",p.Name,p.Age)
}

//1) Create 'byAge' type tp sort by age field. 
type ByAge []Person

//2) Implement the 'Len','Swap' and 'Less' Methods of 'Interface' interface
func (a ByAge) Len() int { return len(a) }
func (a ByAge) Swap(i int, j int) { a[i],a[j] = a[j],a[i] }
func (a ByAge) Less(i int,j int) (bool) {return a[i].Age < a[j].Age}

//sort by Name
type ByName []Person 

func (b ByName) Len() int {return len(b)}
func (b ByName) Swap(i int, j int) {b[i],b[j] = b[j],b[i]}
func (b ByName) Less(i int, j int) (bool) {return b[i].Name < b[j].Name}

func main(){

    p1 := Person{Name:"John",Age:45}
	p2 := Person{Name:"Aaron",Age:34}
	p3 := Person{Name:"Mohammed",Age:25}
	p4 := Person{Name:"Aditya",Age:27}

	persons := []Person{p1,p2,p3,p4}
	personsByAge := ByAge(persons) //3
	sort.Sort(personsByAge) //4
	fmt.Println("Sorted By Age")	
	fmt.Println(personsByAge)	

	personsByName := ByName(persons)
	sort.Sort(personsByName)
	fmt.Println("Sorted by Name")
	fmt.Println(personsByName)
}