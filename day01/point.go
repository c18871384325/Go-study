package main

import "fmt"

func main() {
	//零值 nil
	var (
		pointint    *int
		pointstring *string
	)
	fmt.Printf("%T,%#v\n", pointint, pointint)
	fmt.Printf("%T,%#v\n", pointstring, pointstring)

	// 赋值
	// 使用现有变量 取地址 &name
	age := 26
	fmt.Printf("%T, %#v\n", &age, &age)
	pointage := &age

	fmt.Println(pointage)

	*pointage = 30 // 解引用， 加上* 号后，表示用指针地址取变量值

	age2 := age

	fmt.Println(&age2)  // 声明age2的变量时，会重新想内存申请一个地址，将age的值复制到age2指针指向的值

	age = 28  // 此处 修改age的值后，age2的值不会同步修改
	fmt.Println(age2)

	pointstring = new(string)   //  new 函数会初始化一个string类型零值 （指针默认零值为 nil ，string默认零值为"" 空字符串）
	fmt.Printf("%#v, %#v", pointstring, *pointstring)  // 此处 *pointstring 解引用后得到的值为变量实际值，空字符串
}
