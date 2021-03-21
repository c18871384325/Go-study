// 数组：长度固定的，相同的数据类型的一组值的序列

// 数组一旦定义，长度就不可变了 (长度可以理解为数组内元素数量)
// 数组内的数据类型必须相同，如数据类型都为 int
// 类型： [lenght]type

package main

import "fmt"

func main() {
	// 数组的声明定义

	var names [3]string // 定义一个数组，指定元素为string类型，元素长度为3个，相当可以定义3个数组值
	// 数组的零值为对应类型的零值，此处类型为string，零值为空字符串
	fmt.Printf("%T, %#v\n", names, names) // [3]string, [3]string{"", "", ""}
	// 为数组赋值 （字面量）
	names = [3]string{"huawei", "cisco", "h3c"} // [3]string, [3]string{"huawei", "cisco", "h3c"}
	fmt.Printf("%T, %#v\n", names, names)

	// 定义数组时使用3个 ... 替代数组元素长度，效果类似于类型推导，会根据后面的值推导出元素长度
	testnames := [...]int{1, 2, 3, 4, 5} // [5]int, [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%T, %#v\n", testnames, testnames)

	// 更新数组的值，只更新索引为2的元素的值， 另外2个元素未定义时为默认值
	names = [3]string{2: "jzt"}
	fmt.Printf("%#v\n", names) // [3]string{"", "", "jzt"}

	// 数组的运算  ==   !=    注意：只有元素类型相同的数组才可以进行运算， 如类型都为string
	number1 := [3]int{1, 2, 3}
	number2 := [3]int{2, 3, 4}

	fmt.Println(number1 == number2) // number与number2的值不同，结果为false

	// 数组的遍历， 与字符串遍历方式相同，支持 for ， for rang 方式
	// 第一种方式
	for i := 0; i < len(number1); i++ {
		fmt.Println(number1[i]) // 逐个打印number1 数组内的元素值
	}

	// 第二种方式
	for index, value := range number2 {
		fmt.Println(index, value) // 使用for range 的方式遍历数组内的各元素的值， index 为元素索引， value为索引内的值
	}

	// 多维数组，将数组内的元素定义为数组类型

	// 将数组 many 定义为3个元素长度的数组，同时每个元素定义为 2个元素长度的数组类型， 并为每个元素的数组赋值
	many := [3][2]string{0: [2]string{0: "h1", 1: "h2"}, 1: [2]string{0: "c1", 1: "c2"}, 2: [2]string{0: "b1", 1: "b2"}}

	fmt.Printf("%T, %#v\n", many, many) // [3][2]string, [3][2]string{[2]string{"h1", "h2"}, [2]string{"c1", "c2"}, [2]string{"b1", "b2"}}

	// 更新数组内单个元素索引的值
	many[0] = [2]string{1: "t1"}
	fmt.Println(many[0])

}
