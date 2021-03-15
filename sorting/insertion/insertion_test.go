package insertion

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testInsertionSort(t *testing.T, alg func([]int)) {

	tests := []struct {
		name     string
		testfn   func([]int)
		input    []int
		expected []int
	}{
		{
			name:     "UNSORTED_LIST",
			testfn:   alg,
			input:    []int{2, 4, 1, 3, 7},
			expected: []int{1, 2, 3, 4, 7},
		},
		{
			name:     "LIST_IN_DESCENDING_ORDER",
			testfn:   alg,
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "LIST_ALREADY_IN_ORDER",
			testfn:   alg,
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "TWO_ELEMENT_LIST_UNORDERED",
			testfn:   alg,
			input:    []int{2, 1},
			expected: []int{1, 2},
		},
		{
			name:     "SINGLE_ELEMENT_LIST",
			testfn:   alg,
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "NIL_LIST",
			testfn:   alg,
			input:    nil,
			expected: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			alg(test.input)
			equal := reflect.DeepEqual(test.input, test.expected)
			assert.Equal(t, equal, true)
		})
	}
}

func TestInsertionSort(t *testing.T) {
	testInsertionSort(t, InsertionSort)
}

func benchInsertionSort(b *testing.B, alg func([]int)) {
	b.StopTimer()
	unsorted := make([]int, 1<<10)
	for i := range unsorted {
		unsorted[i] = i ^ 0x2cc
	}
	data := make([]int, len(unsorted))
	for i := 0; i < b.N; i++ {
		copy(data, unsorted)
		b.StartTimer()
		alg(data)
		b.StopTimer()
	}
	b.StopTimer()
}

func BenchmarkInsertionSort(b *testing.B) {
	benchInsertionSort(b, InsertionSort)
}
