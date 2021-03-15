package selection

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "UNSORTED_LIST",
			input:    []int{2, 4, 1, 3, 7},
			expected: []int{1, 2, 3, 4, 7},
		},
		{
			name:     "LIST_IN_DESCENDING_ORDER",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "LIST_ALREADY_IN_ORDER",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "TWO_ELEMENT_LIST_UNORDERED",
			input:    []int{2, 1},
			expected: []int{1, 2},
		},
		{
			name:     "SINGLE_ELEMENT_LIST",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "NIL_LIST",
			input:    nil,
			expected: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			SelectionSort(test.input)
			equal := reflect.DeepEqual(test.input, test.expected)
			assert.Equal(t, equal, true)
		})
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	b.StopTimer()
	unsorted := make([]int, 1<<10)
	for i := range unsorted {
		unsorted[i] = i ^ 0x2cc
	}
	data := make([]int, len(unsorted))
	for i := 0; i < b.N; i++ {
		copy(data, unsorted)
		b.StartTimer()
		SelectionSort(data)
		b.StopTimer()
	}
	b.StopTimer()
}
