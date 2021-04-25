package main

import (
	"fmt"
	"testing"

)

func CheckValues(actual []int, expected []int, t *testing.T) {
	if len(actual) != len(expected) {
		t.Errorf("Actual len %d is different than len %d expected\n", len(actual), len(expected))
	}

	fmt.Printf("%d %d\n", actual[0], actual[1])
	fmt.Printf("%d %d\n", expected[0], expected[1])

	for index, value := range actual {
		if value != expected[index] {
			t.Errorf("Value %d is different than %d expected\n", value, expected[index])
		}
	}
}

func TestCase1A(t *testing.T) {
	numbers := []int{2, 7, 11, 15}
	target := 9

	expected := []int{0, 1}
	actual := twoSum(numbers, target)

	CheckValues(actual, expected, t)
}

func TestCase1B(t *testing.T) {
	numbers := []int{2, 7, 11, 15}
	target := 9

	expected := []int{0, 1}
	actual := twoSumOptimized(numbers, target)

	CheckValues(actual, expected, t)
}

func TestCase2A(t *testing.T) {
	numbers := []int{3, 2, 4}
	target := 6

	expected := []int{1, 2}
	actual := twoSum(numbers, target)

	CheckValues(actual, expected, t)
}

func TestCase2B(t *testing.T) {
	numbers := []int{3, 2, 4}
	target := 6

	expected := []int{1, 2}
	actual := twoSumOptimized(numbers, target)

	CheckValues(actual, expected, t)
}

func TestCase3A(t *testing.T) {
	numbers := []int{3, 3, 3}
	target := 6

	expected := []int{0, 1}
	actual := twoSum(numbers, target)

	CheckValues(actual, expected, t)
}

func TestCase3B(t *testing.T) {
	numbers := []int{3, 3, 3}
	target := 6

	expected := []int{0, 1}
	actual := twoSumOptimized(numbers, target)

	CheckValues(actual, expected, t)
}
