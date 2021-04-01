package hashing

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testSubset(t *testing.T, alg func([]int, []int) bool) {

	tests := []struct {
		name     string
		testfn   func([]int, []int) bool
		arr1     []int
		arr2     []int
		expected bool
	}{
		{
			name:     "SECOND_SUBSET_OF_FIRST",
			testfn:   alg,
			arr1:     []int{2, 4, 1, 3, 7},
			arr2:     []int{3, 4, 7},
			expected: true,
		},
		{
			name:     "FIRST_SUBSET_OF_SECOND",
			testfn:   alg,
			arr1:     []int{3, 4, 7},
			arr2:     []int{2, 4, 1, 3, 7},
			expected: true,
		},
		{
			name:     "NOT_SUBSET",
			testfn:   alg,
			arr1:     []int{3, 4, 7},
			arr2:     []int{8, 9, 10, 11, 12},
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			isSubset := alg(test.arr1, test.arr2)
			assert.Equal(t, test.expected, isSubset)
		})
	}
}

func TestSubsetV1(t *testing.T) {
	testSubset(t, SubsetV1)
}

func TestSubsetV2(t *testing.T) {
	testSubset(t, SubsetV2)
}

func benchmarkSubset(b *testing.B, alg func([]int, []int) bool) {
	b.StopTimer()
	arr1 := make([]int, 1<<10)
	for i := range arr1 {
		arr1[i] = i ^ 0x2cc
	}
	arr2 := make([]int, 1<<8)
	for i := range arr2 {
		arr2[i] = i ^ 0x2cb
	}
	for i := 0; i < b.N; i++ {
		b.StartTimer()
		alg(arr1, arr2)
		b.StopTimer()
	}
}

func BenchmarkSubset(b *testing.B) {
	v := os.Getenv("V")
	if v == "" || strings.TrimSpace(v) == "1" {
		benchmarkSubset(b, SubsetV1)
	}
	if v == "" || strings.TrimSpace(v) == "2" {
		benchmarkSubset(b, SubsetV2)
	}
}
