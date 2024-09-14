package util

func CalculateOffset(page int, pageSize int) int {
	offset := 0
	if page != 0 {
		offset += (page - 1) * pageSize
	}

	return offset
}
