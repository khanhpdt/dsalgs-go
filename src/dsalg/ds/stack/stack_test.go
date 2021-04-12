package stack

import "testing"

func TestPush(t *testing.T) {
	var s = new(Stack)
	s.Push("a")
	s.Push("b")
	s.Push("c")

	var stackAsString = s.String()
	if stackAsString != "c -> b -> a" {
		t.Errorf(`stack = %q, want: "c -> b -> a"`, stackAsString)
	}
}

func TestPop(t *testing.T) {
	var s = new(Stack)
	s.Push("a")
	s.Push("b")
	s.Push("c")

	key, _ := s.Pop()
	if key != "c" {
		t.Errorf(`got: %q, want: "c"`, key)
	}

	key, _ = s.Pop()
	if key != "b" {
		t.Errorf(`got: %q, want: "b"`, key)
	}

	key, _ = s.Pop()
	if key != "a" {
		t.Errorf(`got: %q, want: "a"`, key)
	}

	_, ok := s.Pop()
	if ok {
		t.Error("Pop from empty stack return OK")
	}
}
