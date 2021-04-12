package heap

import (
	"dsalg/utils/slice"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	h := new(MaxHeap)

	h.Insert("c")
	if r, _ := h.Root(); r != "c" {
		t.Errorf(`root = %q, want: "c"`, r)
	}

	h.Insert("a")
	if r, _ := h.Root(); r != "c" {
		t.Errorf(`root = %q, want: "c"`, r)
	}

	h.Insert("d")
	if r, _ := h.Root(); r != "d" {
		t.Errorf(`root = %q, want: "d"`, r)
	}
}

func TestRootEmpty(t *testing.T) {
	var h *MaxHeap = nil
	if _, ok := h.Root(); ok {
		t.Errorf("root of nil return OK OK")
	}

	h = new(MaxHeap)
	if _, ok := h.Root(); ok {
		t.Errorf("root of empty returns OK")
	}
}

func TestLen(t *testing.T) {
	h := new(MaxHeap)

	if h.Len() != 0 {
		t.Errorf("len = %d, want: 0", h.Len())
	}

	h.Insert("a")

	if h.Len() != 1 {
		t.Errorf("len = %d, want: 1", h.Len())
	}

	h.Insert("b")

	if h.Len() != 2 {
		t.Errorf("len = %d, want: 2", h.Len())
	}
}

func TestNewMaxHeap(t *testing.T) {
	tests := []struct {
		items []string
		want  []string
	}{
		{items: []string{}, want: []string{}},
		{items: []string{"a"}, want: []string{"a"}},
		{items: []string{"a", "b"}, want: []string{"b", "a"}},
		{items: []string{"a", "b", "c"}, want: []string{"c", "b", "a"}},
		{items: []string{"a", "c", "d", "e", "b"}, want: []string{"e", "c", "d", "a", "b"}},
	}

	for _, test := range tests {
		h := NewMaxHeap(test.items)

		keys := h.Keys()

		if !slice.Equal(keys, test.want) {
			t.Errorf("got = %v, want: %s", keys, test.want)
		}
	}
}
