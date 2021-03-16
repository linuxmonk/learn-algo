package intersection

import (
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testIntersect(t *testing.T, alg func([]int, []int) []int) {

	tests := []struct {
		name     string
		testfn   func([]int, []int) []int
		setA     []int
		setB     []int
		expected []int
	}{
		{
			name:     "SAME_SIZE_SOME_COMMON",
			testfn:   alg,
			setA:     []int{2, 4, 1, 3, 7},
			setB:     []int{9, 8, 7, 6, 5},
			expected: []int{7},
		},
		{
			name:     "SAME_SIZE_SAME_ALL_COMMON",
			testfn:   alg,
			setA:     []int{5, 4, 3, 2, 1},
			setB:     []int{5, 4, 3, 2, 1},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "DIFFERENT_SIZE_SOME_COMMON",
			testfn:   alg,
			setA:     []int{1, 2, 3, 4, 5},
			setB:     []int{1, 3, 6},
			expected: []int{1, 3},
		},
		{
			name:     "NOTHING_COMMON",
			testfn:   alg,
			setA:     []int{1, 2, 3, 4, 5},
			setB:     []int{6, 7, 8, 9, 0},
			expected: nil,
		},
		{
			name:     "ONE_SET_EMPTY_OR_NIL",
			testfn:   alg,
			setA:     []int{},
			setB:     []int{1, 2, 3, 4, 5},
			expected: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ret := alg(test.setA, test.setB)
			equal := reflect.DeepEqual(ret, test.expected)
			assert.Equal(t, equal, true)
		})
	}
}

func TestIntersect(t *testing.T) {
	testIntersect(t, Intersect)
}

func TestIntersectV2(t *testing.T) {
	testIntersect(t, IntersectV2)
}

func benchIntersect(b *testing.B, alg func([]int, []int) []int) {
	b.StopTimer()
	setA := make([]int, 0)
	for i := 0; i < 1000; i++ {
		setA = append(setA, i^0x2cc)
	}
	setB := make([]int, 0)
	for i := 500; i < 1500; i++ {
		setB = append(setB, i^0x2cc)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_ = alg(setA, setB)
	}
	b.StopTimer()
}

func BenchmarkIntersect(b *testing.B) {
	v := os.Getenv("V")
	if v == "" || strings.TrimSpace(v) == "1" {
		benchIntersect(b, Intersect)
	} else {
		benchIntersect(b, IntersectV2)
	}

}
