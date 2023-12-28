package main

import "testing"

func TestPrintNumber(t *testing.T) {
	type arg struct {
		boxes  []int
		prices []int
		k      int
		v      int
		n      int
	}
	tests := []struct {
		name string
		args arg
		want int
	}{
		{
			name: "should return max chocolate box to avoid invitation",
			args: arg{
				boxes:  []int{5, 9, 10, 4, 14},
				prices: []int{1, 10, 2, 1, 3},
				k:      3,
				v:      10,
				n:      5,
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChocolateBoxToAvoidInvitation(tt.args.boxes, tt.args.prices, tt.args.k, tt.args.v, tt.args.n)
			want := tt.want

			if want != got {
				t.Errorf("want: %v but got: %v args: %v", want, got, tt.args)
			}
		})
	}
}
