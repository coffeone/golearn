package foo

import "testing"

func TestIsOne(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsOne(tt.args.n); got != tt.want {
				t.Errorf("IsOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
