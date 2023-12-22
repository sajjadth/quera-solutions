package main

import "testing"

func TestPrintNumber(t *testing.T) {
	type arg struct {
		n int
		a int
		b int
		c int
		d int
	}
	tests := []struct {
		name string
		args arg
		want int
	}{
		{
			name: "should return printed numbers",
			args: arg{
				n: 24,
				a: 2,
				b: 3,
				c: 4,
				d: 5,
			},
			want: 17,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countDivisibleNumbers(tt.args.n, tt.args.a, tt.args.b, tt.args.c, tt.args.d)
			want := tt.want

			if want != got {
				t.Errorf("want: %v but got: %v args: %v", want, got, tt.args)
			}
		})
	}
}
