package main

import "testing"

func TestMinTotalShoutsForDaraAndSara(t *testing.T) {
	type arg struct {
		m int
		n int
		a int
		b int
	}
	tests := []struct {
		name string
		args arg
		want int
	}{
		{
			name: "should return minimum stages",
			args: arg{m: 3, n: 3, a: 1, b: 3},
			want: 1,
		},
		{
			name: "should return minimum stages",
			args: arg{m: 2, n: 2, a: 1, b: 1},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMinimumStages(tt.args.m, tt.args.n, tt.args.a, tt.args.b)
			want := tt.want

			if want != got {
				t.Errorf("want: %v but got: %v args: %v", want, got, tt.args)
			}
		})
	}
}
