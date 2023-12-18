package main

import "testing"

func TestCountTileWorkMethods(t *testing.T) {
	tests := []struct {
		name string
		args int
		want int
	}{
		{
			name: "should return count tile work methods",
			args: 6,
			want: 13,
		},
		{
			name: "should return count tile work methods",
			args: 7,
			want: 21,
		},
		{
			name: "should return count tile work methods",
			args: 8,
			want: 34,
		},
		{
			name: "should return count tile work methods",
			args: 9,
			want: 55,
		},
		{
			name: "should return count tile work methods",
			args: 10,
			want: 89,
		},
		{
			name: "should return count tile work methods",
			args: 11,
			want: 144,
		},
		{
			name: "should return count tile work methods",
			args: 12,
			want: 233,
		},
		{
			name: "should return count tile work methods",
			args: 13,
			want: 377,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countTileWorkMethods(tt.args)
			want := tt.want

			if want != got {
				t.Errorf("want: %v but got: %v args: %v", want, got, tt.args)
			}
		})
	}
}
