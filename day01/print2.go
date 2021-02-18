package main

import (
	"fmt"
)

func main() {
	var a = 10
	fmt.Printf("%3d\n", a) // %3d 表示打印时占用3个光标的位置， 如此处a的值为10 ，仅占用2个光标，则打印结果为  " 10" ,空出一个
	fmt.Printf("%04d", a)  // 此处表示空出的光标位置，使用0 替代，不使用默认的空格
}
