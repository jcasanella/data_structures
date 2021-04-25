package main

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
*/

// Complexity O^2
func twoSum(nums []int, target int) []int {
	for idx := 0; idx < len(nums)-1; idx++ {

		expected := target - nums[idx]

		for idx2 := idx + 1; idx2 < len(nums); idx2++ {

			if nums[idx2] == expected {
				output := []int{idx, idx2}
				return output
			}
		}
	}

	return nil
}

// Time complexity O(n)
func twoSumOptimized(nums []int, target int) []int {
	values := make(map[int]int)

	for index, value := range nums {
		values[value] = index
	}

	for idx := 0; idx < len(nums)-1; idx++ {
		lookFor := target - nums[idx]
		if lookFor != nums[idx] {

			value, found := values[lookFor]

			if found {
				output := []int{idx, value}
				return output
			}
		}
	}

	return nil
}

func main() {

}
