package main

import "testing"

func TestComparet(t *testing.T) {
	type arg struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args arg
		want string
	}{
		{
			name: "should return compare",
			args: arg{a: 123, b: 421},
			want: "421 < 123",
		},
		{
			name: "should return compare",
			args: arg{a: 123, b: 123},
			want: "123 = 123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := compare(tt.args.a, tt.args.b)
			want := tt.want

			if want != got {
				t.Errorf("want: %v but got: %v args: %v", want, got, tt.args)
			}
		})
	}
}
