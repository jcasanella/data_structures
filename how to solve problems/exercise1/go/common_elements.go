package main

import "fmt"

func hasCommonElements(array1 []byte, size1 int, array2 []byte, size2 int) bool {
	hashSetValues := map[byte]struct{}{}
	for i := 0; i < size1; i++ {
		hashSetValues[array1[i]] = struct{}{}
	}

	for i := 0; i < size2; i++ {
		if _, found := hashSetValues[array2[i]]; found {
			return true
		}
	}

	return false
}

func main() {
	// Given 2 arrays, create a function than let's user know (true, false) whether these 2 arrays contains
	// any common items
	// For example:
	// array1 = ['a', 'b', 'c', 'x']
	// array2 = ['z', 'y', 'i']
	// should return false

	// array1 = ['a', 'b', 'c', 'x']
	// array2 = ['z', 'y', 'x']
	// should return true
	array1 := []byte{'a', 'b', 'c', 'x'}
	array2 := []byte{'z', 'y', 'i'}

	var res1 bool = hasCommonElements(array1, len(array1), array2, len(array2))
	fmt.Println(res1)

	array3 := []byte{'z', 'y', 'x'}
	var res2 bool = hasCommonElements(array1, len(array1), array3, len(array3))
	fmt.Println(res2)
}
