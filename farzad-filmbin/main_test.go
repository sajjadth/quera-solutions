package main

import "testing"

func TestPrintNumber(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "should organize movie names with mixed cases",
			args: []string{
				"rEd riDing HOoD",
				"DraCUla",
				"Bad LiEutenAnt",
			},
			want: "Red Riding Hood\nDracula\nBad Lieutenant",
		},
		{
			name: "should organize movie names with mixed cases",
			args: []string{
				"21 jUMp Street",
				"Mr. SMith GoEs To WashinGTon",
			},
			want: "21 Jump Street\nMr. Smith Goes To Washington",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := organizeMovieNames(tt.args)
			want := tt.want

			if want != got {
				t.Errorf("want: %v but got: %v args: %v", want, got, tt.args)
			}
		})
	}
}
