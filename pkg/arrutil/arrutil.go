// package arrutil helps perform common but verbose tasks on arrays and slices.
//
// Operations are O(n) unless otherwise noted.
package arrutil

// file arrutil.go contains the types and convenience functions used by components of package arrutil

// type arrlookup serves to assist with repeated searches of arrays/slices.
//
// This is not intended to be a general-use solution for searching; it is specialized
// for inorder traversal, and repetitive searching, of a slice of values, where repeated generic Contains()
// would not perform well enough.
//
// Correct uses:
//   - functions that require repeated searching of a slice
//   - functions that want to iterate unique array values in-order
//
// Incorrect uses:
//   - functions that only need to search a slice < 3 times
//   - functions that need to operate on duplicate slice values
type arrLookup[K comparable] struct {
	lkp map[K]bool
	arr []K
}

// Create a new arrLookup from an array of comparable type K.
func newLookup[K comparable](from []K) *arrLookup[K] {
	lkp := &arrLookup[K]{
		make(map[K]bool),
		make([]K, len(from)),
	}
	ct := 0
	for _, item := range from {
		if _, exists := lkp.lkp[item]; !exists {
			lkp.lkp[item] = true
			lkp.arr[ct] = item
			ct++
		}
	}
	lkp.arr = lkp.arr[:ct]
	return lkp
}

func (l *arrLookup[K]) inorder(f func(K)) {
	for _, item := range l.arr {
		f(item)
	}
}

func (l *arrLookup[K]) contains(item K) bool {
	_, ok := l.lkp[item]
	return ok
}

// Get a unique array by creating a lookup then returning its underlying array
func unique[K comparable](from []K) []K {
	lkp := newLookup(from)
	return lkp.arr
}
