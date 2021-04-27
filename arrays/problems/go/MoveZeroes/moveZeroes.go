package main

/*
Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Note that you must do this in-place without making a copy of the array.
*/

// O(n^2)
func moveZeroes(nums []int) {
	for idx := 0; idx < len(nums)-1; idx++ {
		if nums[idx] == 0 {
			idx2 := idx + 1
			for idx2 < len(nums) {
				if nums[idx2] != 0 {
					aux := nums[idx2]
					nums[idx2] = nums[idx]
					nums[idx] = aux
					break
				} else {
					idx2++
				}
			}
		}
	}
}

// O(n)
func moveZeroesImproved(nums []int) {
	var aux []int

	for idx := 0; idx < len(nums); idx++ {
		if nums[idx] != 0 {
			aux = append(aux, nums[idx])
		}
	}

	nonZero := len(aux)
	totalNums := len(nums)
	zeros := totalNums - nonZero

	for idx := 0; idx < zeros; idx++ {
		aux = append(aux, 0)
	}

	for idx := 0; idx < len(nums); idx++ {
		nums[idx] = aux[idx]
	}
}
