package main

import "fmt"

func main() {
	var msg = "双照科技"   // 转义符 如果要将字符串分行打印可以 \n  如 "双照\n科技"， 打印时会换行打印
    // var msg = "双照\\n科技"  \ 表示转义符

	fmt.Printf("%T, %s\n", msg, msg)  // %s 表示字符串占位符

	//操作
	// 字符串连接 + +=
	msg += "————huawei"
	fmt.Printf("%T, %#v", msg, msg)  // 双照科技————huawei

	// 关系运算  > < >= <= != ==
		fmt.Println("abc" > "acb") // false  比较从第一个字符a开始，第一个字符相同，比较第二个字符，此处 c 大于b ，因此该运算结果为false
		fmt.Println("abc" == "abc") // true 

	//索引，切片
	msg = "abcd"
	fmt.Printf("%T, %#v, %c\n", msg[0], msg[0], msg[0])  // 取字符串的第一个元素 a
	fmt.Println(msg[1:3])  // 取字符串内的 1~3 之间的元素  bc  

	// len 占用字节大小
	msgRaw := "查看"
	fmt.Println(len(msg))  // 4 
	fmt.Println(len(msgRaw))  // 6

}