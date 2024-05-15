package main

import "fmt"

type shape interface {
	CalculateArea() float64
}

type square struct {
	side float64
}

func (s *square) CalculateArea() float64 {
	return s.side * s.side
}

type rectangle struct {
	length, breadth float64
}

func (r *rectangle) CalculateArea() float64 {
	return r.length * r.breadth
}

//Type assertion is used to retrieve the underlying concrete type value from the interface. Here

func main() {

	var sh shape
	sh = &square{side: 12}

	/*WRONG METHOD of DOING TYPE ASSERTION. Since 'sh' is actually of type 'square' and not 'rectangle', below statement throws a panic.
	rc := sh.(*rectangle)
	fmt.Println(rc)*/

	/*type assertion of shape interface to square.
	Type assertion returns two values, converted-type 'sq' and boolean variable 'ok' to indicate if the interface actually contains the concrete type that you are converting to. */
	sq, ok := sh.(*square)
	if ok {
		fmt.Println(sq)
	}

	rc, ok := sh.(*rectangle)
	if ok {
		fmt.Println(rc)
	}

}
