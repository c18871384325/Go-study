package main

import "fmt"

func main() {

	// 切片，长度不固定的，类型相同的一组序列

	// 声明：
	var names = []string{} // 声明切片，并初始化值， 切片零值为 nil

	names = []string{"huawei", "cisco", "h3c"} // 直接赋值

	fmt.Printf("%T, %#v\n", names, names)

	// 第二种赋值方式：
	names = []string{1: "中国", 2: "美国"} // 根据索引赋值
	fmt.Printf("%T, %#v\n", names, names)

	// 第三种赋值方式
	// make 函数
	// make()  2个参数
	// make()  3个参数

	names = make([]string, 5) //申请5个字符串元素的切片

	fmt.Printf("%T, %#v\n", names, names)

	names = make([]string, 0, 10) // 申请0个字符串元素的切片，但切片容量是10
	fmt.Printf("%#v\n", names)

	// 获取切片的长度和容量  len函数获取长度，  cap函数获取容量
	fmt.Println(len(names), cap(names))

	// 为切片追加元素，   注：使用make定义元素为3个长度，当需要定义第4个元素值时，需要使用 append函数
	// 注：当切片定义的元素长度使用完成后，直接使用索引添加赋值会报错，只能通过 append追加
	names = append(names, "jzt")
	fmt.Printf("%#v", names)

	// 切片的遍历
	names = []string{"a", "b", "c"}
	for i := 0; i < len(names); i++ {
		fmt.Println(i, names[i])
	}

	for index, value := range names {
		fmt.Println(index, value)
	}

	// copy 切片之间的赋值
	// copy 目的 源
	aslice := []string{"h", "j", "k"}
	bslice := []string{"a", "b", "c"}

	copy(aslice, bslice) // 将切片 bslice 的值赋值给 切片 aslice， 打印aslice时，值为bslice的值，注意如果两个切片元素长度不同，则只复制两个切片都存在的元素索引的值

	fmt.Println(aslice, bslice)

	// 切片操作  针对对象有2个， 1：数组， 2：切片
	// start <= end <= cap
	// nums[1:3]   [start:end]
	nums := []int{0, 1, 2, 3, 4, 5}
	numchild := nums[1:3] // 相当于 nums切片的 1，2索引  ，不包含3
	// new_cap = cap - start
	fmt.Printf("%T, %#v\n", numchild, numchild)

	// 切片更改时要注意，切片操作是共享底层的数组
	// 示例：  该种方式一般不建议使用，容易出现bug
	numchild = append(numchild, 100)       // 也会将 nums的元素值修改，因为numchild产生于 nums，与nums共用底层数组，当numchild元素值修改，nums也随着修改
	fmt.Printf("%#v, %#v", nums, numchild) // []int{0, 1, 2, 100, 4, 5}, []int{1, 2, 100}

	// 切片的操作一般不使用，可以直接使用新的切片编写程序
}
