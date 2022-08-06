package stringsplit

import (
	"reflect"
	"testing"
)
type splitvar struct {
	str string
	sep string
	want []string
}
// TestSplit tests the Split function.
func TestSplit(t *testing.T) {
	testvar := map[string]splitvar{
		"one": {str: "abcdbcd", sep: "bc", want: []string{"a","d","d"}},
		"two": {str: "abcdbcd", sep: "b",  want: []string{"a","cd","cd"}},
		"three": {str: "上海自来水来自海上", sep: "来", want: []string{"上海自","水","自海上"}},
	}
	for k, v := range testvar {
		t.Run(k, func(t *testing.T) {
			got := Split(v.str, v.sep)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("Split(%v, %v) = %v, want %v", v.str, v.sep, got, v.want)
			}
		})	
	}
}
