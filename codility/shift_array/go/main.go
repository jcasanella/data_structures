package main

import "fmt"

func shiftNumbers(numbers []int, shift int) []int {
	size := len(numbers)
	if len(numbers) == 1 {
		return numbers
	}

	var resp = make([]int, size)
	for idx := 0; idx < size; idx++ {
		posic := calcIndex(idx, shift, size)
		resp[posic] = numbers[idx]
	}

	return resp
}

func calcIndex(index int, shift int, length int) int {
	posic := index + shift
	if posic < length {
		return posic
	} else {
		idx := length - posic
		if idx < 0 {
			return idx * -1
		} else {
			return idx
		}
	}
}

func main() {
	var numbers = [...]int{3, 8, 9, 7, 6}
	shift := 3

	output := shiftNumbers(numbers[:], shift)
	fmt.Println(output)
}
