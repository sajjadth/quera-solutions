package main

import "testing"

func TestOperationRequired(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want int
	}{
		{
			name: "should return operation required",
			args: []string{"0"},
			want: 0,
		},
		{
			name: "should return operation required",
			args: []string{"01", "10"},
			want: 1,
		},
		{
			name: "should return operation required",
			args: []string{"01010", "00101", "11000", "00110"},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := operationRequired(tt.args)
			want := tt.want

			if want != got {
				t.Errorf("want: %v but got: %v", want, got)
			}
		})
	}
}
