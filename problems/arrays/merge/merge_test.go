package merge

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testMerge(t *testing.T, alg func([]int, []int) []int) {

	tests := []struct {
		name   string
		testfn func([]int, []int) []int
		arr1   []int
		arr2   []int
		arr3   []int
	}{
		{
			name:   "TWO_SORTED_LISTS_FIRST_LONGER_THAN_SECOND",
			testfn: alg,
			arr1:   []int{1, 5, 9, 14},
			arr2:   []int{2, 3, 7, 10, 17},
			arr3:   []int{1, 2, 3, 5, 7, 9, 10, 14, 17},
		},
		{
			name:   "TWO_SORTED_LISTS_EQ_SIZE",
			testfn: alg,
			arr1:   []int{1, 3, 5, 7},
			arr2:   []int{2, 4, 6, 8},
			arr3:   []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name:   "SORTED_LISTS_REPEATING_ELEMS",
			testfn: alg,
			arr1:   []int{1, 2, 3, 4, 5},
			arr2:   []int{1, 2, 3, 4, 5},
			arr3:   []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5},
		},
		{
			name:   "SINGLE_ELEMENT_LISTS",
			testfn: alg,
			arr1:   []int{1},
			arr2:   []int{2},
			arr3:   []int{1, 2},
		},
		{
			name:   "NIL_LIST",
			testfn: alg,
			arr1:   nil,
			arr2:   nil,
			arr3:   nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := alg(test.arr1, test.arr2)
			equal := reflect.DeepEqual(result, test.arr3)
			assert.Equal(t, equal, true)
		})
	}
}

func TestMerge(t *testing.T) {
	testMerge(t, Merge)
}
