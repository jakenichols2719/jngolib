package arrutil_test

import (
	"jngolib/pkg/arrutil"
	"testing"
)

var intSlice []int = []int{0, 1, 2, 3, 4, 5}
var strSlice []string = []string{"hello", "world", "my", "world", "name", "is", "jakenichols2719"}

func TestContains(t *testing.T) {
	// Basic Contains
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
	// ContainsAll
	// Valid; strSlice contains all of "hello" and "world"
	if !arrutil.ContainsAll(strSlice, []string{"hello", "world"}) {
		t.Errorf("ContainsAll(strSlice, ['hello, 'world']) == false; should be true")
	}
	// Invalid; strSlice does not contain all of "hello", "world", and "foo"
	if arrutil.ContainsAll(strSlice, []string{"hello", "world", "foo"}) {
		t.Errorf("ContainsAll(strSlice), ['hello', 'world', 'foo'] == true; should be false. strSlice does not contain foo.")
	}
	// ContainsExactly
	// Valid; intSlice contains exactly 0->5
	if !arrutil.ContainsExactly(intSlice, []int{1, 0, 3, 2, 5, 4}) {
		t.Errorf("ContainsExactly(intSlice, [1, 0, 3, 2, 5, 4]) == false; should be true")
	}
	// Invalid; intSlice does not contain 6
	if arrutil.ContainsExactly(intSlice, []int{0, 1, 2, 3, 4, 5, 6}) {
		t.Errorf("ContainsExactly(intSlice, [0...6]) == true; should be false. vals contains a value that intSlice does not.")
	}
	// Invalid; intSlice also contains 5
	if arrutil.ContainsExactly(intSlice, []int{0, 1, 2, 3, 4}) {
		t.Errorf("ContainsExactly(intSlice, [0...4] == true; should be false. intSlice contains a value that vals does not.")
	}
}

func TestFind(t *testing.T) {
	if idx := arrutil.FindFirst(intSlice, 2); idx != 2 {
		t.Errorf("FindFirst(intSlice, 2) == %d; should be 2", idx)
	}
	if idx := arrutil.FindFirst(intSlice, 7); idx != len(intSlice) {
		t.Errorf("FindFirst(intSlice, 7) == %d; should be %d", idx, len(intSlice))
	}
	if idxs := arrutil.FindAll(strSlice, "world"); !(arrutil.Contains(idxs, 1) && arrutil.Contains(idxs, 3)) {
		t.Errorf("FindAll(strSlice, 'world') == %v; should be [1, 3]", idxs)
	}
	if idxs := arrutil.FindAll(strSlice, "foo"); len(idxs) != 0 {
		t.Errorf("FindAll(strSlice, 'foo') == %v; should be []", idxs)
	}
}

func TestCompare(t *testing.T) {
	// And with test slice, compare using ContainsExactly
	testSlice := []int{0, 0, 3, 4, 5, 7, 2}
	// AND
	resAnd, err := arrutil.Compare(intSlice, testSlice, arrutil.OpAnd)
	if err != nil {
		t.Errorf("intSlice AND testSlice returned %v", err)
	}
	if !arrutil.ContainsExactly(resAnd, []int{0, 2, 3, 4, 5}) {
		t.Errorf("resAnd == %v; should have exactly [0, 2, 3, 4, 5]", resAnd)
	}
	// OR
	resOr, err := arrutil.Compare(intSlice, testSlice, arrutil.OpOr)
	if err != nil {
		t.Errorf("intSlice OR testSlice returned %v", err)
	}
	if !arrutil.ContainsExactly(resOr, []int{0, 1, 2, 3, 4, 5, 7}) {
		t.Errorf("resOr == %v; should have exactly [0, 1, 2, 3, 4, 5, 7]", resAnd)
	}
	// BOTH (this should fail)
	resWrong, err := arrutil.Compare(intSlice, testSlice, arrutil.OpOr|arrutil.OpAnd)
	if err == nil {
		t.Errorf("intSlice or|and testSlice returned nil error; should fail")
	}
	if len(resWrong) != 0 {
		t.Errorf("intSlice or|and testSlice returned %v, <err>, but the result should have len(0)", resWrong)
	}
}
