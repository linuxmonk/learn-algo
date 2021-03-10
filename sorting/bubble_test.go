package sorting

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {

	tests := []struct {
		input    []int
		expected []int
		err      error
	}{
		{
			input:    []int{2, 4, 1, 3, 7},
			expected: []int{1, 2, 3, 4, 7},
			err:      nil,
		},
		{
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
			err:      nil,
		},
	}
	for _, test := range tests {
		err := Bubble(test.input)
		assert.Equal(t, err, nil)
		equal := reflect.DeepEqual(test.input, test.expected)
		assert.Equal(t, equal, true)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	// make an array of up to 1000 random numbers
	n := rand.Intn(1000)
	data := make([]int, n)
	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())
		data = append(data, rand.Intn(10000))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Bubble(data)
	}
}
