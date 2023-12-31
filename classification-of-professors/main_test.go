package main

import "testing"

func TestCalculateHInex(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "should return h-index",
			args: []int{1, 5, 3, 4, 2},
			want: 3,
		},
		{
			name: "should return h-index",
			args: []int{1, 13, 1, 3, 5, 2, 21, 8},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateHIndex(tt.args)
			want := tt.want

			if want != got {
				t.Errorf("want: %v but got: %v args: %v", want, got, tt.args)
			}
		})
	}
}
