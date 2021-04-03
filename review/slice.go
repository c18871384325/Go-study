package main

import (
	"fmt"
)

func main() {

	// 数组：元素长度不固定的，同类型的变量的集合

	// 定义,数组未初始化时，零值为 nil
	var demo []int
	fmt.Printf("%#v\n", demo)

	// 数组初始化之后零值为数组变量类型的零值
	demo = []int{} // 初始化数组
	fmt.Printf("%#v\n", demo)

	// 赋值、
	demo = []int{1, 2, 3, 4, 5, 100: 100} // 数组内的元素长度不固定，可以直接赋值100的元素
	fmt.Printf("%T, %#v\n", demo, demo)

	// make 函数  slice := make([]string, 3 10)//  3 表示数组元素的长度，10表示数组的容量 , 未指容量时，数组容量默认为元素长度
	demomake := make([]string, 3)

	demomake[0] = "demo0"
	demomake[1] = "demo1"
	demomake[2] = "demo2"
	//  demomake[3] = "demo3"  元素长度不足会报错
	fmt.Printf("%#v\n", demomake[0])
	fmt.Printf("%#v\n", demomake[1])
	fmt.Printf("%#v\n", demomake[2])
	fmt.Printf("%T, %#v\n", demomake, demomake)

	// 切片的切片

	// start < end < capacity
	demoslice := []int{1, 2, 3, 4, 5, 6}

	demosliceone := demoslice[3:4] // 3：4 表示取 demoslice 切片元素3~4以下的值，不包含4因此demosliceone 仅有一个值    注意：元素从0开始
	fmt.Printf("%#v\n", demosliceone)

	demosliceone = demoslice[3:5:5] //   ：5 表切片的容量结束大小，不能小于end,此处切片容量为2， 因为仅有3，4 两个元素赋值， 若后面不加5， 则切片容量为6 (0,1,2,3,4,5 5个元素，因为容量为6)
	fmt.Printf("%#v\n", demosliceone)

	// 遍历
	for i := 0; i < len(demosliceone); i++ {
		fmt.Println(i, demosliceone[i])
	}

	for i, v := range demosliceone {
		fmt.Println(i, v)
	}

	// 切片的操作， copy 切片的值
	// copy  目的 源
	aslice := []string{"x", "y", "z"}
	bslice := []string{"a", "b"}
	fmt.Printf("%#v, %#v\n", aslice, bslice)

	copy(aslice, bslice) // bslice 仅有2个元素，因此会将元素0，1 拷贝到aslice 的元素 0，1  ，aslice 的元素2的值不变
	fmt.Printf("%#v, %#v\n", aslice, bslice)

	// 获取切片的容量和长度
	fmt.Println(len(aslice), cap(aslice))

	//
	numslice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	arrayslice := numslice[3:4:4]
	fmt.Println(cap(arrayslice))
	// 1
	arrayslice = append(arrayslice, 100)
	fmt.Println(cap(arrayslice))
	//2

	// 当start 为0 时可省略， 当end 为 len(slice)-1 时可省略
	startslice := []int{1, 2, 3, 4, 5, 6}
	qslice := startslice[:3] // []int, []int{1, 2, 3}
	fmt.Printf("%T, %#v\n", qslice, qslice)
	fmt.Println(cap(qslice)) // 6

	endslice := []int{1, 2, 3, 4, 5, 6}
	mslice := endslice[3:]
	fmt.Printf("%T, %#v\n", mslice, mslice)
	fmt.Println(cap(mslice))

	// 删除切片的某个值

	// 下面示例删除切片内元素4
	middleslice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	wslice := middleslice[5:]
	copy(middleslice[4:], wslice)
	middleslice = middleslice[:len(middleslice)-1]
	fmt.Printf("%#v\n", middleslice)

}
