package hashing

// Array subset problem -
// Given two arrays determine if one is a subset of another.
//

func SubsetV1(arr1, arr2 []int) bool {

	len1 := len(arr1)
	len2 := len(arr2)

	if len1 == 0 || len2 == 0 {
		return false
	}

	smaller := arr1
	larger := arr2
	if len2 < len1 {
		smaller = arr2
		larger = arr1
	}

	for _, s := range smaller {
		found := false
		for _, l := range larger {
			if s == l {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func SubsetV2(arr1, arr2 []int) bool {

	len1 := len(arr1)
	len2 := len(arr2)
	m1 := make(map[int]bool, len1)
	m2 := make(map[int]bool, len2)
	for _, v := range arr1 {
		m1[v] = true
	}
	for _, v := range arr2 {
		m2[v] = true
	}
	smaller := m1
	larger := m2
	if len2 < len1 {
		smaller = m2
		larger = m1
	}
	isSubset := true
	for k := range smaller {
		_, ok := larger[k]
		if !ok {
			isSubset = false
			break
		}
	}
	return isSubset
}
