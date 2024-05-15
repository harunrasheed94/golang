package main

import (
	"fmt"
	"sort"
)

func main() {
	intSlice := []int{23, 1, 45, 12, 8, 7}

	fmt.Println("Unsorted slice = ", intSlice)
	sort.Ints(intSlice) //sort.Ints sorts the slice but has NO RETURN VALUE
	fmt.Println("Sorted slice = ", intSlice)

	names := []string{"Said", "Tim", "John", "Aaron", "Mo"}

	fmt.Println("Unsorted names = ", names)
	sort.Strings(names)
	fmt.Println("Sorted names = ", names)

}
