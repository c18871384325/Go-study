package main

import "fmt"

func main() {
	isGra := true // 布尔类型的变量，值只能为 true,false, 一般表示真假
	fmt.Printf("%T %#v", isGra, isGra)
	fmt.Println("")
	/*
		布尔逻辑运算：与，或，非
		与：左操作符与又操作符都为true，结果为true &&
		或：左操作符与右操作符其中一个为true，结果为true ||
		非：取反， true => false  , false => true !
	*/
	a, b, c, d := true, true, false, false

	fmt.Println("a, b:", a && b) // a,b 为true， 结果为true  &&
	fmt.Println("a, c:", a && c) // a为ture,c为false，结果为false
	fmt.Println("c, b:", c && b) // c为false,b为true,结果为false
	fmt.Println("c, d:", c && d) // c 为false,d为false,结果为false

	fmt.Println("a, b:", a || b) // a,b 为true， 结果为true ||
	fmt.Println("a, c:", a || c) // a为ture,c为false，结果为true
	fmt.Println("c, b:", c || b) // c为false,b为true,结果为true
	fmt.Println("c, d:", c || d) // c 为false,d为false,结果为false

	fmt.Println("a:", !a) // !取反，a的值为true打印结果为false

	// 关系   等于 ==      不等于!=

	// 布尔类型打印使用 %t 做占位符
	fmt.Println(a == b) // ture == true   true
	fmt.Println(a == c) // true == false  false
	fmt.Println(a != b) // true != true false
	fmt.Println(a != c) // true != false true

	fmt.Printf("%t, %t", a, c) // 打印布尔类型变量的值使用 %t作为占位符，此处打印 a , c 变量的布尔值
}

