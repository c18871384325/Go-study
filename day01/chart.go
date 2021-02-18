package main

import "fmt"

func main() {
	letters := "abcdehuawei"
	var i = len(letters)

	for j := 0; j < i; j++ {
		fmt.Printf("%c\n", letters[j]) // %c 表示打印字节的占位符
	}

	letters = "华为容器云平台"
	for k, v := range letters {   // 注：当不需要打印k 时，可以用——下划线替代，在Go语言中——的变量不需要使用 示例： for ——, v := range letters
		fmt.Printf("%T, %#v, %T, %#v\n", k, k, v, v)
		fmt.Printf("%q\n", v)   // %q 表示打印字符串切片后的字符内容 
	}
}
