package main

import (
	"fmt"
	"strings"
)

func main() {
	// 字符串操作

	// 比较
	a := "13"
	b := "12"
	fmt.Println(strings.Compare(a, b))   // 比较a与B的值
	fmt.Println(strings.Compare("d", "e"))  // 字节英文字符也可以比较 a -- z 的顺序大小
	// a == b  返回0，   a < b 返回 -1  ,  a > b 返回1 



	// 包含
	fmt.Println(strings.Contains("my name is chengz", "chen"))   // 左侧的值内是否包含右侧给定的内容，包含则为true,否则为false
	fmt.Println(strings.Contains("my name is chengz", "xyz"))  // 必须全部包含结果才为true

	
	// 计算次数
	fmt.Println(strings.Count("aabbcc", "cc"))  // 返回值为1， cc字符串出现1次
	fmt.Println(strings.Count("aabbcc", "c"))  //  返回值为2 ，c 出现2次, 给定的值未出现时，结果为0


	// 比较
	
	chart() {
		
	}

}