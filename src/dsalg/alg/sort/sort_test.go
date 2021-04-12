package sort

import (
	"dsalg/utils/slice"
	"testing"
)

var tests = []struct {
	input []string
	want  []string
}{
	{input: nil, want: nil},
	{input: []string{}, want: []string{}},
	{input: []string{"c"}, want: []string{"c"}},
	{input: []string{"c", "b"}, want: []string{"b", "c"}},
	{input: []string{"b", "c", "a"}, want: []string{"a", "b", "c"}},
	{input: []string{"b", "c", "a", "d"}, want: []string{"a", "b", "c", "d"}},
	{input: []string{"b", "c", "e", "a", "d"}, want: []string{"a", "b", "c", "d", "e"}},
	{input: []string{"b", "e", "c", "a", "f", "d"}, want: []string{"a", "b", "c", "d", "e", "f"}},
}

func runTest(t *testing.T, sortFunc func([]string) []string) {
	for _, test := range tests {
		sorted := sortFunc(test.input)
		if !slice.Equal(sorted, test.want) {
			t.Errorf("got = %v, want = %v", sorted, test.want)
		}
	}
}

func TestHeapSort(t *testing.T) {
	runTest(t, HeapSort)
}

func TestMergeSort(t *testing.T) {
	runTest(t, MergeSort)
}
