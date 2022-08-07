package arrutil

import "fmt"

type ArrayCompareFlag uint8

// A AND B: any element in both, preserve order of A
const OpAnd ArrayCompareFlag = 1

// A OR B: any element in either, preserve order of A then B
const OpOr ArrayCompareFlag = 2

// Create a slice where each item is in A *and* B.
func compareAnd[K comparable](A []K, B []K) []K {
	alkp, blkp := newLookup(A), newLookup(B)
	out := make([]K, len(A))
	ct := 0
	blkp.inorder(func(item K) {
		if alkp.contains(item) {
			out[ct] = item
			ct++
		}
	})
	if ct > 0 {
		out = out[:ct]
	}
	return out
}

// Create a slice where each item is in A *or* B.
func compareOr[K comparable](A []K, B []K) []K {
	alkp, blkp := newLookup(A), newLookup(B)
	out := make([]K, len(A)+len(B))
	ct := 0
	alkp.inorder(func(item K) {
		out[ct] = item
		ct++
	})
	blkp.inorder(func(item K) {
		out[ct] = item
		ct++
	})
	if ct > 0 {
		out = out[:ct]
	}
	return unique(out)
}

// Compare 2 slices given an operator.
//
// This function preserves the order of
func Compare[K comparable](A []K, B []K, op ArrayCompareFlag) ([]K, error) {
	var out []K = nil
	switch op & 0b111 {
	case OpAnd:
		{
			out = compareAnd(A, B)
			break
		}
	case OpOr:
		{
			out = compareOr(A, B)
			break
		}
	default:
		{
			return nil, fmt.Errorf("ArrayCompareFlag %b is invalid", op)
		}
	}
	return out, nil
}
