package main

import "testing"

func TestPrintNumber(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "should return printed numbers",
			args: "50943",
			want: "5: 55555\n0: \n9: 999999999\n4: 4444\n3: 333",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := printNumber(tt.args)
			want := tt.want

			if want != got {
				t.Errorf("want: %v but got: %v args: %v", want, got, tt.args)
			}
		})
	}
}
