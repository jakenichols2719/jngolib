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

// ContainsAll checks if ALL values exist in a slice.
//
// This returns a plain true/false value; it does not tell you which values do not exist
// if false is returned.
func ContainsAll[K comparable](arr []K, vals []K) bool {
	uniquemap := make(map[K]K)
	for _, arrItem := range arr {
		uniquemap[arrItem] = arrItem
	}
	for _, valsItem := range vals {
		if _, ok := uniquemap[valsItem]; !ok {
			return false
		}
	}
	return true
}

// ContainsAll checks if ALL and ONLY a values exist in a slice.
//
// This function does account for duplicated in arr and vals; if arr is [2] and vals is [2, 2],
// ContainsExactly(arr, vals) == false.
// This returns a plain true/false value; it does not tell you which values do/do not exist
// if false is returned.
//
// ContainsExactly is O(n^2).
func ContainsExactly[K comparable](arr []K, vals []K) bool {
	cpy := make([]K, len(arr))
	copy(cpy, arr)
	for _, valsItem := range vals {
		// Try to find each value in vals
		if idx := FindFirst(cpy, valsItem); idx == len(cpy) {
			// If we didn't find it, return false. arr must contain every item in vals.
			return false
		} else {
			// If we found the value, remove it by replacing with last value then truncating cpy
			cpy[idx] = cpy[len(cpy)-1]
			cpy = cpy[:len(cpy)-1]
		}
	}
	// If every value in vals was found, len(cpy) should be 0.
	return len(cpy) == 0
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
