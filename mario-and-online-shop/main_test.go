package main

import "testing"

func TestMinTotalShoutsForDaraAndSara(t *testing.T) {
	type arg struct {
		cm   int
		gm   int
		cs   int
		gs   int
		rate int
	}
	tests := []struct {
		name string
		args arg
		want string
	}{
		{
			name: "should return can buy a special sword or not",
			args: arg{cm: 5, gm: 2, cs: 9, gs: 0, rate: 3},
			want: "Yes",
		},
		{
			name: "should return can buy a special sword or not",
			args: arg{cm: 10, gm: 2, cs: 1, gs: 4, rate: 5},
			want: "No",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canBuySpecialSword(tt.args.cm, tt.args.gm, tt.args.cs, tt.args.gs, tt.args.rate)
			want := tt.want

			if want != got {
				t.Errorf("want: %v but got: %v args: %v", want, got, tt.args)
			}
		})
	}
}
