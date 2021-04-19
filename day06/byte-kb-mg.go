package main

import (
	"fmt"
	"strconv"
)

func conver(n int64) string {

	gb := n / 1000 / 1000 / 1000
	mb := n / 1000 / 1000
	kb := n / 1000

	if gb > 1 {
		return strconv.Itoa(int(gb))
	} else if mb > 1 {
		return strconv.Itoa(int(mb))
	} else if kb > 1 {
		return strconv.Itoa(int(kb))
	}
	return strconv.Itoa(int(n))
}

func main() {
	fmt.Println(conver(17778))
}
