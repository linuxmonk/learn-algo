package largest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargest(t *testing.T) {

	tests := []struct {
		name   string
		input  []int
		retval int
		err    error
	}{
		{
			name:   "LIST_OF_NUMBERS",
			input:  []int{2, 4, 1, 3, 7},
			retval: 7,
			err:    nil,
		},
		{
			name:   "ALL_NUMS_SAME",
			input:  []int{5, 5, 5, 5, 5},
			retval: 5,
			err:    nil,
		},
		{
			name:   "LIST_OF_TWO_ELEMS",
			input:  []int{2, 1},
			retval: 2,
			err:    nil,
		},
		{
			name:   "1_ELEM_LIST",
			input:  []int{1},
			retval: 1,
			err:    nil,
		},
		{
			name:   "NIL_LIST",
			input:  nil,
			retval: -1,
			err:    ErrEmptyList,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			retval, err := Largest(test.input)
			assert.Equal(t, test.retval, retval)
			assert.Equal(t, test.err, err)
		})
	}
}

func testLargestK(t *testing.T, K int) {
	unsorted := make([]int, K*(1<<10))
	for i := range unsorted {
		unsorted[i] = i ^ 0x2cc
	}
	data := make([]int, len(unsorted))
	_, err := Largest(data)
	assert.Equal(t, err, nil)
}

func TestLargest1K(t *testing.T) {
	testLargestK(t, 1)
}

func TestLargest10K(t *testing.T) {
	testLargestK(t, 10)
}

func BenchmarkLargest(b *testing.B) {
	b.StopTimer()
	unsorted := make([]int, 1<<10)
	for i := range unsorted {
		unsorted[i] = i ^ 0x2cc
	}
	data := make([]int, len(unsorted))
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		copy(data, unsorted)
		b.StartTimer()
		Largest(data)
		b.StopTimer()
	}
	b.StopTimer()
}
