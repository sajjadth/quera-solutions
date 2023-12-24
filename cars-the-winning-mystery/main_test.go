package main

import "testing"

func TestGetWinnerCarNumber(t *testing.T) {
	type arg struct {
		loop     float64
		length   float64
		carsInfo [][2]float64
	}
	tests := []struct {
		name string
		args arg
		want int
	}{
		{
			name: "should organize movie names with mixed cases",
			args: arg{loop: 20, length: 10, carsInfo: [][2]float64{{100, 100}, {50, 4}, {1, 2}}},
			want: 2,
		},
		{
			name: "should organize movie names with mixed cases",
			args: arg{loop: 200, length: 200, carsInfo: [][2]float64{{100, 100}, {50, 90}, {0, 80}}},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getWinnerCarNumber(tt.args.loop, tt.args.length, tt.args.carsInfo)
			want := tt.want

			if want != got {
				t.Errorf("want: %v but got: %v args: %v", want, got, tt.args)
			}
		})
	}
}
