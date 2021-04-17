package main

func ReverseString(input string) string {
	if len(input) == 0 {
		return input
	}

	output := make([]rune, len(input))

	index2 := len(input) - 1
	for _, character := range input {
		output[index2] = character
		index2--
	}

	return string(output)
}
