package main

import (
	"fmt"
)

// 当需要定义的变量比较多时，可以将同类型的所有变量直接存储在相同的数组内，数组内可存储的变量个数为数组元素的长度
// 数组： 一组变量的集合

func main() {
	// 定义一个数组，并指定元素长度为3,类型为string,并初始化
	// var： 声明一个变量      demo： 变量名称    [3]string{} ： 变量类型（此处为3个元素长度的string变量）
	var demo [3]string // {} 表示对该数组初始化 ， 未初始化时，零值也是为空字符串 ""
	demo = [3]string{} // 初始化数组
	fmt.Printf("%T, %#v\n", demo, demo)
	// [3]string, [3]string{"", "", ""}

	// 赋值， 字面量 =》 0，1，2，3，4，5，...
	// 第一种 demo = [3]string{"huawei", "cisco", "h3c"}   都可以
	demo = [3]string{0: "huawei", 1: "cisco", 2: "h3c"}
	fmt.Printf("%T, %#v\n", demo, demo)

	// 第二种：
	demotwo := [...]string{"jzt", "hwjzt", "hnjzt"}
	fmt.Printf("%T, %#v\n", demotwo, demotwo)

	// 第三种
	demo[1] = "TENCENT" // 直接赋值某个元素值
	fmt.Printf("%#v\n", demo)

	// 访问元素，修改索引
	fmt.Printf("%q\n", demo[1])
	demo[1] = "beego"
	fmt.Printf("%#v\n", demo)

	// 关系运算  ==  ！=  , 结果为true,或false
	//false
	fmt.Println(demo == [3]string{})

	//true
	fmt.Println(demo == [3]string{0: "huawei", 1: "beego", 2: "h3c"})

	// 函数len
	fmt.Println(len(demo))

	// 遍历

	for i := 0; i < len(demo); i++ {
		fmt.Println(i, demo[i])
	}

	for i, v := range demo { // i = index,   v = value
		fmt.Println(i, v)
	}

	// 二维数组  ，数组内嵌套数组

	demothree := [3][2]string{0: {"onenil", "oneone"}, 1: {0: "twonil", 1: "twoone"}, 2: {0: "threenil", 1: "threeone"}}

	fmt.Printf("%T, %#v\n", demothree, demothree)

	fmt.Printf("%#v\n", demothree[1])

	fmt.Printf("%#v\n", demothree[1][1])

}
