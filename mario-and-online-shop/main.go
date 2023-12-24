package main

import "fmt"

func canBuySpecialSword(cm, gm, cs, gs, rate int) string {

	if cm > cs {
		gm += (cm - cs) / rate
		cm -= cm - cs
	} else if gm > gs {
		cm += rate * (gm - gs)
		gm -= gm - gs
	}

	if cm >= cs && gm >= gs {
		return "Yes"
	}

	return "No"
}
func main() {
	var cm, gm, cs, gs, rate int

	fmt.Scanln(&cm, &gm)
	fmt.Scanln(&cs, &gs)
	fmt.Scanln(&rate)

	fmt.Println(canBuySpecialSword(cm, gm, cs, gs, rate))
}
