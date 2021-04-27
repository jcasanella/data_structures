package main

// Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

// Cost O(n)
func containsDuplicate(nums []int) bool {
	var m map[int]int

	for _, value := range nums {
		_, found := m[value]
		if found {
			return true
		}

		if m == nil {
			m = make(map[int]int)
		}
		m[value] = value
	}

	return false
}
