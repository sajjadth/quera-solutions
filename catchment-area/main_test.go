package main

import "testing"

func TestFindFirstOverflowOperation(t *testing.T) {
	type arg struct {
		d        []int
		stock    []int
		capacity []int
		q        int
	}
	tests := []struct {
		name string
		args arg
		want int
	}{
		{
			name: "should return find first overflow operation",
			args: arg{
				capacity: []int{5, 2, 1, 3, 4},
				d:        []int{2, 2, 4, 4, 4, 4, 3, 2},
				stock:    make([]int, 5),
				q:        8,
			},
			want: 7,
		},
		{
			name: "should return find first overflow operation",
			args: arg{
				capacity: []int{3, 4, 2, 3},
				d:        []int{3, 3, 2},
				stock:    make([]int, 4),
				q:        3,
			},
			want: 0,
		},
		{
			name: "should return find first overflow operation",
			args: arg{
				capacity: []int{1, 4, 2, 3},
				d:        []int{3, 3, 2, 3},
				stock:    make([]int, 4),
				q:        4,
			},
			want: 0,
		},
		{
			name: "should return find first overflow operation",
			args: arg{
				capacity: []int{3, 4, 2, 1},
				d:        []int{3, 3, 2, 2, 2, 2, 2, 2},
				stock:    make([]int, 4),
				q:        8,
			},
			want: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findFirstOverflowOperation(tt.args.d, tt.args.stock, tt.args.capacity, tt.args.q)
			want := tt.want

			if want != got {
				t.Errorf("want: %v but got: %v args: %v", want, got, tt.args)
			}
		})
	}
}
