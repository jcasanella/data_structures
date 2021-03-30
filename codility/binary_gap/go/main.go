package main

import "fmt"

func calcBinary(num int) []int {
	var digits []int

	for num > 0 {
		digit := num % 2
		digits = append(digits, digit)
		num /= 2
	}

	return digits
}

func maxBinaryGaps(digits []int) int {
	var (
		startToCount          bool = false
		maxNumGaps, countGaps int  = 0, 0
	)
	for i := 1; i < len(digits); i++ {
		if digits[i-1] == 1 && digits[i] == 0 {
			startToCount = true
		}

		if startToCount {
			if digits[i] == 0 {
				countGaps++
			} else {
				if maxNumGaps < countGaps {
					maxNumGaps = countGaps
				}

				countGaps = 0
				startToCount = false
			}
		}
	}

	return maxNumGaps
}

func main() {
	var num int = 5
	digits := calcBinary(num)

	fmt.Printf("%v\n", digits)

	gaps := maxBinaryGaps(digits)
	fmt.Printf("Number gaps: %d\n", gaps)
}
