package main

import "testing"

func TestPrintNumber(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "should smallest larger series",
			args: []string{"acs", "lgeuvf", "dqwrmse", "zwsked"},
			want: "asc\nlgevfu\ndqwrsem\nno answer",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findSmallestLargerSeries(tt.args)
			want := tt.want

			if want != got {
				t.Errorf("want: %v but got: %v args: %v", want, got, tt.args)
			}
		})
	}
}
