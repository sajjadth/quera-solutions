package main

import "testing"

func TestMinTotalShoutsForDaraAndSara(t *testing.T) {
	type arg struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args arg
		want int
	}{
		{
			name: "should return min uncoverable length",
			args: arg{a: 5, b: 2},
			want: 1,
		},
		{
			name: "should return min uncoverable length",
			args: arg{a: 10, b: 8},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateMinUncoverableLength(tt.args.a, tt.args.b)
			want := tt.want

			if want != got {
				t.Errorf("want: %v but got: %v args: %v", want, got, tt.args)
			}
		})
	}
}
