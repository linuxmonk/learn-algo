package selection

func SelectionSort(list []int) {

	if list == nil || len(list) <= 1 {
		return
	}

	unsortedIndex := 0
	n := len(list)

	for i := unsortedIndex; i < n; i++ {
		lowest, lowestIndex := list[i], i
		foundLower := false
		for j := i; j < n; j++ {
			if list[j] < lowest {
				lowest, lowestIndex = list[j], j
				foundLower = true
			}
		}
		if foundLower {
			list[unsortedIndex], list[lowestIndex] = list[lowestIndex], list[unsortedIndex]
		}
		unsortedIndex++
	}
}
