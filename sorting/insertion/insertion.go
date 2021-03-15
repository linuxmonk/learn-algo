package insertion

func InsertionSort(list []int) {
	if list == nil || len(list) <= 1 {
		return
	}

	n := len(list)
	pos := 0
	for pivot := 1; pivot < n; pivot++ {
		temp := list[pivot]
		shifted := false
		pos = 0
		for i := pivot - 1; i >= 0; i-- {
			if list[i] > temp {
				list[i+1] = list[i]
				pos++
				shifted = true
			}
		}
		if shifted {
			list[abs(pos-pivot)] = temp
		}
	}
}

//go:inline
func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}
