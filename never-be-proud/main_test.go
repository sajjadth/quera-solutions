package main

import "testing"

func TestPrintNumber(t *testing.T) {
	tests := []struct {
		name string
		args [][2]int
		want int
	}{
		{
			name: "should return chromatic number",
			args: [][2]int{{1, 2}, {2, 3}},
			want: 2,
		},
		{
			name: "should return chromatic number",
			args: [][2]int{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := chromaticNumberOfGraph(tt.args)
			want := tt.want

			if want != got {
				t.Errorf("want: %v but got: %v args: %v", want, got, tt.args)
			}
		})
	}
}
