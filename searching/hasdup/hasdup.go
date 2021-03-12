package hasdup

func HasDupV1(list []int) bool {

	if list == nil || len(list) <= 1 {
		return false
	}

	n := len(list)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && list[i] == list[j] {
				return true
			}
		}
	}
	return false
}

func HasDupV2(list []int) bool {

	if list == nil || len(list) <= 1 {
		return false
	}
	n := len(list)
	countMap := make(map[int]int, n)
	for i := 0; i < n; i++ {
		countMap[list[i]]++
	}
	for _, v := range countMap {
		if v > 1 {
			return true
		}
	}
	return false
}
