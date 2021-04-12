package queue

import "testing"

func TestEnqueue(t *testing.T) {
	q := new(Queue)
	q.Enqueue("a")
	q.Enqueue("b")
	q.Enqueue("c")

	qs := q.String()
	if qs != "a -> b -> c" {
		t.Errorf(`got: %q, want: "a -> b -> c"`, qs)
	}
}

func TestDequeue(t *testing.T) {
	q := new(Queue)
	q.Enqueue("a")
	q.Enqueue("b")
	q.Enqueue("c")

	k, _ := q.Dequeue()
	if k != "a" {
		t.Errorf(`got: %q, want: "a"`, k)
	}

	k, _ = q.Dequeue()
	if k != "b" {
		t.Errorf(`got: %q, want: "b"`, k)
	}

	k, _ = q.Dequeue()
	if k != "c" {
		t.Errorf(`got: %q, want: "c"`, k)
	}

	_, ok := q.Dequeue()
	if ok {
		t.Errorf("Dequeue empty queue returns OK")
	}
}

func TestEqual(t *testing.T) {
	tests := []struct {
		q1   *Queue
		q2   *Queue
		want bool
	}{
		{nil, nil, true},
		{nil, NewQueue("a"), false},
		{NewQueue("a"), nil, false},
		{NewQueue("a", "b"), NewQueue("a", "b"), true},
		{NewQueue("a", "b"), NewQueue("b", "a"), false},
		{NewQueue("a", "b"), NewQueue("a", "b", "a"), false},
	}
	for _, test := range tests {
		if isEqual := test.q1.Equal(test.q2); isEqual != test.want {
			t.Errorf("%s == %s? got %t, want: %t", test.q1, test.q2, isEqual, test.want)
		}
	}
}
