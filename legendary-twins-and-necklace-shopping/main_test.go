package main

import "testing"

func TestCheckMatchingNecklace(t *testing.T) {
	tests := []struct {
		name string
		args [2]string
		want string
	}{
		{
			name: "should check necklaces for match",
			args: [2]string{"ab", "abab"},
			want: "NO",
		},
		{
			name: "should check necklaces for match",
			args: [2]string{"abc", "cba"},
			want: "YES",
		},
		{
			name: "should check necklaces for match",
			args: [2]string{"gcd", "lcm"},
			want: "NO",
		},
		{
			name: "should check necklaces for match",
			args: [2]string{"abc", "bca"},
			want: "YES",
		},
		{
			name: "should check necklaces for match",
			args: [2]string{"lca", "lcs"},
			want: "NO",
		},
		{
			name: "should check necklaces for match",
			args: [2]string{"abacb", "abcab"},
			want: "YES",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkMatchingNecklace(tt.args)
			want := tt.want

			if want != got {
				t.Errorf("want: %v but got: %v args: %v", want, got, tt.args)
			}
		})
	}
}
