package main

import (
	"fmt"

	"github.com/jcasanella/algorithms/sorting"

)

func main() {
	values := [...]int{6, 1, 8, 5, 9}
	values2 := sorting.Bubble(values[:])

	fmt.Printf("%v\n", values2)
}
