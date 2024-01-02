package main

import "testing"

func TestMinPassengersDiscomfort(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "should return min passenger discomfort",
			args: []int{1, 1, 1, 1},
			want: 0,
		},
		{
			name: "should return min passenger discomfort",
			args: []int{2, 1, 1, 2},
			want: 0,
		},
		{
			name: "should return min passenger discomfort",
			args: []int{3, 1, 4, 2},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minPassengersDiscomfort(tt.args)
			want := tt.want

			if want != got {
				t.Errorf("want: %v but got: %v args: %v", want, got, tt.args)
			}
		})
	}
}
