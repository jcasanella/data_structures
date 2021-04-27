package main

// Complexity O(n)
func rotate(nums []int, k int) {

	// Trivial case
	if k == 0 || len(nums) == k {
		return
	}

	arr1 := make([]int, len(nums))
	for idx := 0; idx < len(nums); idx++ {
		arr1[(idx+k)%len(nums)] = nums[idx]
	}

	copy(nums, arr1)
}
