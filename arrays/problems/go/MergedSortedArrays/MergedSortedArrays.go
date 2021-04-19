package main

func MergedSortedArrays(array1 []int, array2 []int) []int {
	if array1 == nil && array2 != nil {
		return array2
	} else if array1 != nil && array2 == nil {
		return array1
	} else if array1 == nil && array2 == nil {
		return nil
	}

	var (
		output []int
		idx2   int = 0
	)

	for index, value := range array1 {

		for idx2 < len(array2) && array2[idx2] <= array1[index] {
			output = append(output, array2[idx2])
			idx2++
		}

		output = append(output, value)
	}

	for idx2 < len(array2) {
		output = append(output, array2[idx2])
		idx2++
	}

	return output
}
