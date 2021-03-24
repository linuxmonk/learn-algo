package merge

func Merge(arr1, arr2 []int) []int {

	if arr1 == nil {
		return arr2
	}
	if arr2 == nil {
		return arr1
	}

	longer := arr1
	shorter := arr2
	if len(arr1) < len(arr2) {
		longer = arr2
		shorter = arr1
	}
	mergedList := make([]int, 0, len(arr1)+len(arr2))
	i, j := 0, 0
	len1 := len(shorter)
	len2 := len(longer)
	for i < len1 && j < len2 {
		v := shorter[i]
		if v < longer[j] {
			mergedList = append(mergedList, v)
			i++
		} else {
			mergedList = append(mergedList, longer[j])
			j++
		}
	}
	mergedList = append(mergedList, shorter[i:]...)
	mergedList = append(mergedList, longer[j:]...)
	return mergedList
}
