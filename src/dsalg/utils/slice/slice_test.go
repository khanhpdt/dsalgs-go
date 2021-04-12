package slice

import "testing"

func TestEqual(t *testing.T) {
	tests := []struct {
		x, y []string
		want bool
	}{
		{x: nil, y: nil, want: true},
		{x: []string{}, y: []string{}, want: true},
		{x: []string{"a", "b", "c"}, y: []string{"a", "b", "c"}, want: true},
		{x: []string{"a", "b", "c"}, y: []string{"a", "b", "c", "d"}, want: false},
		{x: []string{"a", "b", "c"}, y: []string{"a", "b"}, want: false},
		{x: []string{"a", "b"}, y: []string{"b", "a"}, want: false},
	}

	for _, test := range tests {
		isEqual := Equal(test.x, test.y)
		if isEqual != test.want {
			t.Errorf("got = %t, want: %t", isEqual, test.want)
		}
	}
}
