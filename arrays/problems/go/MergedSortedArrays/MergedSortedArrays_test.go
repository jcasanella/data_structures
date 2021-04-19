package main

import "testing"

func CompareArrays(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func TestBothArraysNull(t *testing.T) {
	actual := MergedSortedArrays(nil, nil)
	if actual != nil {
		t.Errorf("Actual is different from nil")
	}
}

func TestFirstArrayNull(t *testing.T) {
	array2 := [...]int{1, 2, 3}
	actual := MergedSortedArrays(nil, array2[:])
	if !CompareArrays(actual, array2[:]) {
		t.Errorf("Actual %v is different from expected %v\n", actual, array2)
	}
}

func TestSecondArrayNull(t *testing.T) {
	array1 := [...]int{1, 2, 3}
	actual := MergedSortedArrays(array1[:], nil)
	if !CompareArrays(actual, array1[:]) {
		t.Errorf("Actual %v is different from expected %v\n", actual, array1)
	}
}

func TestArraysOneElement(t *testing.T) {
	array1 := [...]int{200}
	array2 := [...]int{50}

	expected := [2]int{50, 200}
	actual := MergedSortedArrays(array1[:], array2[:])
	if !CompareArrays(actual, expected[:]) {
		t.Errorf("Actual %v is different from expected %v\n", actual, array1)
	}
}

func TestFirstArrayMinor(t *testing.T) {
	array1 := []int{1, 2, 3, 4, 5}
	array2 := []int{6, 7, 8}

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8}
	actual := MergedSortedArrays(array1[:], array2[:])
	if !CompareArrays(actual, expected[:]) {
		t.Errorf("Actual %v is different from expected %v\n", actual, array1)
	}
}

func TestFirstArrayBigger(t *testing.T) {
	array1 := []int{6, 7, 8}
	array2 := []int{1, 2, 3, 4, 5}

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8}
	actual := MergedSortedArrays(array1[:], array2[:])
	if !CompareArrays(actual, expected[:]) {
		t.Errorf("Actual %v is different from expected %v\n", actual, array1)
	}
}

func TestMergeIntoArray1(t *testing.T) {
	array1 := []int{1, 2, 4, 5, 6}
	array2 := []int{3}

	expected := []int{1, 2, 3, 4, 5, 6}
	actual := MergedSortedArrays(array1[:], array2[:])
	if !CompareArrays(actual, expected[:]) {
		t.Errorf("Actual %v is different from expected %v\n", actual, array1)
	}
}

func TestMergeIntoArray1B(t *testing.T) {
	array1 := []int{1, 2, 4, 5, 6}
	array2 := []int{3, 7}

	expected := []int{1, 2, 3, 4, 5, 6, 7}
	actual := MergedSortedArrays(array1[:], array2[:])
	if !CompareArrays(actual, expected[:]) {
		t.Errorf("Actual %v is different from expected %v\n", actual, array1)
	}
}

func TestMergeIntoArray2(t *testing.T) {
	array1 := []int{3}
	array2 := []int{1, 2, 4, 5, 6}

	expected := []int{1, 2, 3, 4, 5, 6}
	actual := MergedSortedArrays(array1[:], array2[:])
	if !CompareArrays(actual, expected[:]) {
		t.Errorf("Actual %v is different from expected %v\n", actual, array1)
	}
}

func TestMergeIntoArray2B(t *testing.T) {
	array1 := []int{3, 7}
	array2 := []int{1, 2, 4, 5, 6}

	expected := []int{1, 2, 3, 4, 5, 6, 7}
	actual := MergedSortedArrays(array1[:], array2[:])
	if !CompareArrays(actual, expected[:]) {
		t.Errorf("Actual %v is different from expected %v\n", actual, array1)
	}
}

func TestMerge(t *testing.T) {
	array1 := []int{3, 7, 21, 100, 150}
	array2 := []int{4, 23, 44, 130, 600}

	expected := []int{3, 4, 7, 21, 23, 44, 100, 130, 150, 600}
	actual := MergedSortedArrays(array1[:], array2[:])
	if !CompareArrays(actual, expected[:]) {
		t.Errorf("Actual %v is different from expected %v\n", actual, array1)
	}
}

func TestMergeWithDuplicates(t *testing.T) {
	array1 := []int{3, 7, 21, 23, 100, 150, 600}
	array2 := []int{3, 4, 23, 44, 130, 600}

	expected := []int{3, 3, 4, 7, 21, 23, 23, 44, 100, 130, 150, 600, 600}
	actual := MergedSortedArrays(array1[:], array2[:])
	if !CompareArrays(actual, expected[:]) {
		t.Errorf("Actual %v is different from expected %v\n", actual, array1)
	}
}
