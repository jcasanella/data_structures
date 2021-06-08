package sorting

func Bubble(values []int) []int {
	if len(values) == 0 {
		return nil
	} else if len(values) == 1 {
		return values
	} else {
		var (
			limit  = len(values) - 1
			sorted = false
		)

		for sorted == false {
			for idx := 0; idx < limit; idx++ {
				if values[idx] > values[idx+1] {
					values[idx], values[idx+1] = values[idx+1], values[idx]
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
