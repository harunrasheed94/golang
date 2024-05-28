package main

import "fmt"

type Numbers interface {
	int64 | float64
}

/*
Generic function 'SumIntsorFloats' to add numeric values.
TYPE PARAMETERS are 'K' and 'V'.
TYPE CONSTRAINT for 'K' is comparable.
TYPE CONSTRAINT for 'V' is Numbers interface which has int64 or float64 type constraints.
*/
func SumIntsorFloats[K comparable, V Numbers](values map[K]V) V {
	var sum V
	for _, val := range values {
		sum += val
	}
	return sum
}

func main() {
	intVals := make(map[string]int64, 0)
	intVals["a1"] = 89
	intVals["a2"] = 19
	intVals["a3"] = 30
	intVals["a4"] = 40
	intVals["a5"] = 50

	floatVals := make(map[string]float64, 0)
	floatVals["a1"] = 29.23
	floatVals["a2"] = 81.33
	floatVals["a3"] = 49.11
	floatVals["a4"] = 28.81
	floatVals["a5"] = 59.11

	fmt.Println("Sum of ints ", SumIntsorFloats(intVals))
	fmt.Println("Sum of floats ", SumIntsorFloats(floatVals))

	//TYPE ARGUMENTS is being passed but as shown above it is NOT MANDATORY.
	// fmt.Println("Sum of ints ", SumIntsorFloats[string, int64](intVals))
	// fmt.Println("Sum of floats ", SumIntsorFloats[string, float64](floatVals))

	//If you uncomment this it will result in COMPILE-TIME ERROR.
	// stringVals := make(map[string]string, 0)
	// stringVals["aa"] = "bb"
	//fmt.Println("Sum of floats ", SumIntsorFloats(stringVals))
}

/*
OUTPUT:
Sum of ints  228
Sum of floats  247.59000000000003
*/
