package util

func InArray(ele int, arr []int) bool {
	for _, item := range arr {
		if item == ele {
			return true
		}
	}

	return false
}
