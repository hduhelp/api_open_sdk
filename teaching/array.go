package teaching

func FirstOfArray(array []int32) int32 {
	m := array[0]
	for _, v := range array {
		if v < m {
			m = v
		}
	}
	return m
}

func LastOfArray(array []int32) int32 {
	m := array[0]
	for _, v := range array {
		if v > m {
			m = v
		}
	}
	return m
}

func InArray(array []int32, item int32) (int32, bool) {
	for _, v := range array {
		if v == item {
			return v, true
		}
	}
	return 0, false
}
