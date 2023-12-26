package main

import "testing"

func TestMinTotalShoutsForDaraAndSara(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want string
	}{
		{
			name: "should return readable result",
			args: []int{5, 2},
			want: "5 2 -1",
		},
		{
			name: "should return readable result",
			args: []int{5, 5},
			want: "5 2 0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := displayReadableResult(tt.args)
			want := tt.want

			if want != got {
				t.Errorf("want: %v but got: %v args: %v", want, got, tt.args)
			}
		})
	}
}
