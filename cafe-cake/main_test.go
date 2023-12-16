package main

import "testing"

func TestCalculateMrFirouzPaymentAmount(t *testing.T) {
	type arg struct {
		cakes []int
		k     int
	}
	tests := []struct {
		name string
		args arg
		want int
	}{
		{
			name: "should return how much Mr.firoozi should pay",
			args: arg{[]int{3, 2, 3}, 2},
			want: 3,
		},
		{
			name: "should return how much Mr.firoozi should pay",
			args: arg{[]int{5, 4, 3, 2, 2}, 3},
			want: 2,
		},
		{
			name: "should return how much Mr.firoozi should pay",
			args: arg{[]int{1, 3, 4, 2}, 1},
			want: 4,
		},
		{
			name: "should return how much Mr.firoozi should pay",
			args: arg{[]int{2}, 5},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateMrFirouzPaymentAmount(tt.args.cakes, tt.args.k)
			want := tt.want

			if want != got {
				t.Errorf("want: %v but got: %v", want, got)
			}
		})
	}
}
