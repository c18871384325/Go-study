package main

import (
	"fmt"
	"strconv"
	"errors"
)

func div(n1, n2 int) (int, error) {
	if n2 == 0 {
		return -1, errors.New("除数为0")
	}
	return n1 / n2, nil
}

func main() {
	value, err := strconv.Atoi("huawei")
	fmt.Printf("%T\n", err)
	fmt.Println(value)

	e1 := errors.New("demo-error")
	fmt.Printf("%T %#v\n", e1, e1)

	
	if v, e2 := div(2, 2); e2 == nil {
		fmt.Println(v)
	}else{
		fmt.Println(e2)
	}
}