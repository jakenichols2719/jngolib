// package arrutil helps perform common but verbose tasks on arrays and slices.
package arrutil

// Contains checks if a value exists in a slice.
//
// This function must be called with a comparable value (can == and !=) and a slice with
// a matching type.
func Contains[K comparable](arr []K, val K) bool {
	for _, item := range arr {
		if item == val {
			return true
		}
	}
	return false
}

// Find the first index where a value occurs in a slice.
//
// This function must be called with a comparable value (can == and !=) and a slice with
// a matching type. Returns an index, or len(arr) if the value is not found.
func FindFirst[K comparable](arr []K, val K) int {
	for idx, item := range arr {
		if item == val {
			return idx
		}
	}
	return len(arr)
}

// Find all indices where a value occurs in a slice.
//
// This function must be called with a comparable value (can == and !=) and a slice with
// a matching type. Returns a slice of int indices, with len 0 if the value is never found.
func FindAll[K comparable](arr []K, val K) []int {
	out := make([]int, 0, len(arr))
	for idx, item := range arr {
		if item == val {
			out = append(out, idx)
		}
	}
	return out
}
