package hasdup

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testHasDup(t *testing.T, alg func([]int) bool) {

	tests := []struct {
		name   string
		input  []int
		retval bool
	}{
		{
			name:   "HAS_NO_DUP",
			input:  []int{2, 4, 1, 3, 7},
			retval: false,
		},
		{
			name:   "HAS_DUP",
			input:  []int{5, 1, 3, 2, 1},
			retval: true,
		},
		{
			name:   "HAS_DUP_2_ELEM",
			input:  []int{2, 2},
			retval: true,
		},
		{
			name:   "1_ELEM_LIST",
			input:  []int{1},
			retval: false,
		},
		{
			name:   "NIL_LIST",
			input:  nil,
			retval: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			retval := alg(test.input)
			assert.Equal(t, test.retval, retval)
		})
	}
}

func TestHasDupV1(t *testing.T) {
	testHasDup(t, HasDupV1)
}

func TestHasDupV2(t *testing.T) {
	testHasDup(t, HasDupV2)
}

func benchHasDup(b *testing.B, alg func([]int) bool) {
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
		alg(data)
		b.StopTimer()
	}
	b.StopTimer()
}

func BenchmarkHasDup(b *testing.B) {
	v := os.Getenv("V")
	if v == "" || strings.TrimSpace(v) == "1" {
		benchHasDup(b, HasDupV1)
	} else {
		benchHasDup(b, HasDupV2)
	}
}
