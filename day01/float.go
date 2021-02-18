package main

import (
	"fmt"
)

func main() {
	var height float32 = 1.70
	var heightDefault = 1.70
	fmt.Printf("%T, %#v, %f\n", height, height, height) // %f float类型的占位符  %#v 以优雅的方式打印变量的值
	fmt.Printf("%T, %#v, %f\n", heightDefault, heightDefault, heightDefault)

	//科学计数法  1e3  =  1* 10^3   = 1000

	// 浮点数的操作： 算数运算 + - * / ++ --
	f1 := 1.5
	f2 := 2.7
	fmt.Println(f1 + f2)
	fmt.Println(f1 - f2)
	fmt.Println(f1 * f2)
	fmt.Println(f1 / f2)
	f1++ // 自加1
	f2-- // 自减1

	fmt.Println(f1, f2) //  因保存结果为非精度，此处f2打印结果为 1.700000000000002

	// 关系运算 > >= < <=    浮点数为非精度，一般不使用等于和不等于判断，如果想要判断，== !=  ，需判断差值在一定范围内
	// 浮点数没有位运算

	// 赋值运算 = += -= *= /=
	f1 += f2        // 等同于  f1 = f1+f2
	fmt.Println(f1) // 此处打印结果为 f1 + f2 的值

}
