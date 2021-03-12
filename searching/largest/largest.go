package largest

import "errors"

var (
	ErrEmptyList = errors.New("nil/empty list")
)

func Largest(list []int) (int, error) {
	if list == nil || len(list) == 0 {
		return -1, ErrEmptyList
	}
	n := len(list)
	if n == 1 {
		return list[0], nil
	}
	largest := list[0]
	for i := 1; i < n; i++ {
		if list[i] > largest {
			largest = list[i]
		}
	}
	return largest, nil
}
