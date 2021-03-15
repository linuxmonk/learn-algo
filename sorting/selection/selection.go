package selection

func SelectionSortV1(list []int) {

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

// Remove unwanted variables from V1 and more compact code.
// !!This seems slower due to extra operations of indexing into array etc.!!
func SelectionSortV2(list []int) {

	if list == nil || len(list) <= 1 {
		return
	}

	n := len(list)
	for i := 0; i < n; i++ {
		lowestIndex := i
		found := false
		for j := i; j < n; j++ {
			if list[j] < list[lowestIndex] {
				lowestIndex = j
				found = true
			}
		}
		if found {
			list[i], list[lowestIndex] = list[lowestIndex], list[i]
		}
	}
}
