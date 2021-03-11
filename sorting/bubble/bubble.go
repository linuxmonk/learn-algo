package bubble

func BubbleV1(list []int) {

	if list == nil || len(list) <= 1 {
		return
	}

	n := len(list)
	swapped := true
	for {
		i := 0
		if !swapped {
			break
		}
		swapped = false
		for j := 1; j < n; j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
				swapped = true
			}
			i++
		}
	}
}

func BubbleV2(list []int) {

	if list == nil || len(list) <= 1 {
		return
	}

	n := len(list) - 1
	sorted := false
	for !sorted {
		sorted = true
		for i := 0; i < n; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				sorted = false
			}
		}
		n--
	}
}
