package sorting

func Bubble(values []int, size int) []int {
	if size == 0 {
		return nil
	} else if size == 1 {
		return values
	} else {
		var (
			limit  = size - 1
			sorted = false
		)

		for sorted == false {
			for idx := 0; idx < limit; idx++ {
				if values[idx] > values[idx+1] {
					temp := values[idx]
					values[idx] = values[idx+1]
					values[idx+1] = temp
				}
			}

			limit--
			if limit <= 1 {
				sorted = true
			}
		}

		return values
	}
}
