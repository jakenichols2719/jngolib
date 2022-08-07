package arrutil_test

import (
	"jngolib/pkg/arrutil"
	"testing"
)

var intSlice []int = []int{0, 1, 2, 3, 4, 5}
var strSlice []string = []string{"hello", "world", "my", "world", "name", "is", "jakenichols2719"}

func TestContains(t *testing.T) {
	if !arrutil.Contains(intSlice, 2) {
		t.Errorf("Contains(intSlice, 2) == false; should be true")
	}
	if arrutil.Contains(intSlice, 7) {
		t.Errorf("Contains(intSlice, 7) == true; should be false")
	}
	if !arrutil.Contains(strSlice, "world") {
		t.Errorf("Contains(strSlice, 'world') == false; should be true")
	}
	if arrutil.Contains(strSlice, "foo") {
		t.Errorf("Contains(strSlice, 'foo') == true; should be false")
	}
}

func TestFindFirst(t *testing.T) {
	if idx := arrutil.FindFirst(intSlice, 2); idx != 2 {
		t.Errorf("FindFirst(intSlice, 2) == %d; should be 2", idx)
	}
	if idx := arrutil.FindFirst(intSlice, 7); idx != len(intSlice) {
		t.Errorf("FindFirst(intSlice, 7) == %d; should be %d", idx, len(intSlice))
	}
}

func TestFindAll(t *testing.T) {
	if idxs := arrutil.FindAll(strSlice, "world"); !(arrutil.Contains(idxs, 1) && arrutil.Contains(idxs, 3)) {
		t.Errorf("FindAll(strSlice, 'world') == %v; should be [1, 3]", idxs)
	}
	if idxs := arrutil.FindAll(strSlice, "foo"); len(idxs) != 0 {
		t.Errorf("FindAll(strSlice, 'foo') == %v; should be []", idxs)
	}
}
