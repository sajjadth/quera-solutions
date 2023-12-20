package main

import "testing"

func TestCheckMatchingNecklace(t *testing.T) {
	type arg struct {
		x1 int
		x2 int
		x3 int
		x4 int
		y1 int
		y2 int
		y3 int
		y4 int
	}
	tests := []struct {
		name string
		args arg
		want string
	}{
		{
			name: "should return happiness status",
			args: arg{
				x1: 1,
				y1: 1,
				x2: 1,
				y2: 2,
				x3: 1,
				y3: 3,
				x4: 8,
				y4: 2,
			},
			want: "unhappy",
		},
		{
			name: "should return happiness status",
			args: arg{
				x1: 1,
				y1: 1,
				x2: 2,
				y2: 2,
				x3: 3,
				y3: 3,
				x4: 4,
				y4: 4,
			},
			want: "unhappy",
		},
		{
			name: "should return happiness status",
			args: arg{
				x1: 2,
				y1: 1,
				x2: 4,
				y2: 3,
				x3: 2,
				y3: 7,
				x4: 8,
				y4: 6,
			},
			want: "happy",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := printHappinessStatus(tt.args.x1, tt.args.x2, tt.args.x3, tt.args.x4, tt.args.y1, tt.args.y2, tt.args.y3, tt.args.y4)
			want := tt.want

			if want != got {
				t.Errorf("want: %v but got: %v args: %v", want, got, tt.args)
			}
		})
	}
}
