package bubble

import (
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testBubbleSort(t *testing.T, alg func([]int)) {

	tests := []struct {
		name     string
		input    []int
		expected []int
		fn       func([]int)
		err      error
	}{
		{
			name:     "RANDOM_UNSORTED",
			input:    []int{2, 4, 1, 3, 7},
			expected: []int{1, 2, 3, 4, 7},
			fn:       alg,
			err:      nil,
		},
		{
			name:     "REVERSE_ORDERED",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
			fn:       alg,
			err:      nil,
		},
		{
			name:     "TWO_ELEMENT_UNSORTED",
			input:    []int{2, 1},
			expected: []int{1, 2},
			fn:       alg,
			err:      nil,
		},
		{
			name:     "SINGLE_ELEMENT",
			input:    []int{1},
			expected: []int{1},
			fn:       alg,
			err:      nil,
		},
		{
			name:     "NIL_LIST",
			input:    nil,
			expected: nil,
			fn:       alg,
			err:      nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.fn(test.input)
			equal := reflect.DeepEqual(test.input, test.expected)
			assert.Equal(t, equal, true)
		})
	}
}

func TestBubbleSort(t *testing.T) {
	testBubbleSort(t, BubbleV1)
}

func TestBubbleSortV2(t *testing.T) {
	testBubbleSort(t, BubbleV2)
}

func benchBubbleSort(b *testing.B, alg func([]int)) {
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

func BenchmarkBubbleSort(b *testing.B) {
	v := os.Getenv("V")
	if v == "" || strings.TrimSpace(v) == "1" {
		benchBubbleSort(b, BubbleV1)
	} else {
		benchBubbleSort(b, BubbleV2)
	}
}
