package sorting

import "errors"

func Bubble(list []int) error {

	if list == nil || len(list) == 0 {
		return errors.New("nil/empty input")
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
	return nil
}
