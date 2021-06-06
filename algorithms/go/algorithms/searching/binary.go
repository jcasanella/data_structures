package searching

func Binary(values []int, target int) bool {
	if len(values) == 0 {
		return false
	} else if len(values) == 1 {
		return values[0] == target
	} else {
		var (
			start = 0
			end   = len(values) - 1
		)

		for start < end {
			mid := (start + end) / 2

			if values[mid] < target {
				start = mid + 1
			} else if values[mid] > target {
				end = mid - 1
			} else {
				return true
			}
		}

		return false
	}
}
