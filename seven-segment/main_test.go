package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "should return seven segment",
			args: "2.3e10",
			want: 64,
		},
		{
			name: "should return seven segment",
			args: "8e10000000",
			want: 60000007,
		},
		{
			name: "should return seven segment",
			args: "9e0",
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sevenSegment(tt.args)
			want := tt.want

			if want != got {
				t.Errorf("want: %v but got: %v args: %v", want, got, tt.args)
			}
		})
	}
}
