package stringsplit

import (
	"reflect"
	"testing"
)

// TestSplit tests the Split function.
func Test1Split(t *testing.T) {
	str := "abcdbcd"
	sep := "bc"
	got := Split(str, sep)
	want := []string{"a","d","d"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Split(%v, %v) = %v, want %v", str, sep, got, want)
	}
}

func Test2Split(t *testing.T){
	str := "abcdbcd"
	sep := "b"
	got := Split(str, sep)
	want := []string{"a","cd","cd"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Split(%v, %v) = %v, want %v", str, sep, got, want)
	}
}