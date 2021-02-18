package main

import "fmt"

func main() {
	// 在控制台打印 1..10
	for index := 1; index <= 10; index++ {
		fmt.Println(index)
	}

	// 计算1-100 的和
	a := 0
	for i := 1; i <= 100; i++ { // 初始化子语句，条件子语句，后置子语句  for后面跟3种子语句
		a += i
	}
	fmt.Println(a)

	var b = 1
	for b <= 10 { // 仅有条件子语句，当for语句执行完成后，再次判断条件是否满足，直到条件判断结果为false，停止执行
		fmt.Println(b)
		b++
	}
}
