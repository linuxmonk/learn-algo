package intersection

// Return a slice of intersecting values between the two sets.
// Uses brute force technique (N-Squared) to compare every element
// in 'a' with every element in 'b'
//
func Intersect(a []int, b []int) []int {
	if a == nil || len(a) == 0 || b == nil || len(b) == 0 {
		return nil
	}
	aLen := len(a)
	bLen := len(b)
	n := aLen
	if aLen > bLen {
		n = bLen
	}
	result := make([]int, 0, n)
	for i := 0; i < aLen; i++ {
		for j := 0; j < bLen; j++ {
			if a[i] == b[j] {
				result = append(result, a[i])
			}
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}

// To improvise the intersection a bit. There is no need to have the inner
// loop run through all the elements of 'b' if an element from 'a' is found.
// Breaking there would speed this up a bit.
func IntersectV2(a []int, b []int) []int {
	if a == nil || len(a) == 0 || b == nil || len(b) == 0 {
		return nil
	}
	aLen := len(a)
	bLen := len(b)
	n := aLen
	if aLen > bLen {
		n = bLen
	}
	result := make([]int, 0, n)
	for i := 0; i < aLen; i++ {
		for j := 0; j < bLen; j++ {
			if a[i] == b[j] {
				result = append(result, a[i])
				break
			}
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}
