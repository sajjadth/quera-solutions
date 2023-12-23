package main

import "testing"

func TestMinDistanceFromAminsImaginaryCastle(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "should return printed numbers",
			args: []int{1, 2, 1, 0, 2},
			want: 0,
		},
		{
			name: "should return printed numbers",
			args: []int{1, 1, 1},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minDistanceFromAminsImaginaryCastle(tt.args)
			want := tt.want

			if want != got {
				t.Errorf("want: %v but got: %v args: %v", want, got, tt.args)
			}
		})
	}
}
