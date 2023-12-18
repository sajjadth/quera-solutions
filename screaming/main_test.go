package main

import "testing"

func TestMinTotalShoutsForDaraAndSara(t *testing.T) {
	tests := []struct {
		name string
		args int
		want int
	}{
		{
			name: "should return minimum total shouts for dara and sara",
			args: 1,
			want: 2,
		},
		{
			name: "should return minimum total shouts for dara and sara",
			args: 2,
			want: 3,
		},
		{
			name: "should return minimum total shouts for dara and sara",
			args: 3,
			want: 3,
		},
		{
			name: "should return minimum total shouts for dara and sara",
			args: 4,
			want: 3,
		},
		{
			name: "should return minimum total shouts for dara and sara",
			args: 40,
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minTotalShoutsForDaraAndSara(tt.args)
			want := tt.want

			if want != got {
				t.Errorf("want: %v but got: %v args: %v", want, got, tt.args)
			}
		})
	}
}
