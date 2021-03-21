package main

import (
	"fmt"
)


func demo(name string, grade int) string {
	if grade > 80 {
		return "优秀"
	}else {
		return "不及格"
	}
}

func main () {
	gate := demo("huawei", 75)
	fmt.Println(gate)
}
