package main

import "fmt"

func main() {
	// int8,16,32,64   (int 类型支持负数如 -127~127)
	// uint8,16,32,64  （带 u 的为无符号整数，只能表示0和0以上的正数）
	// byte rune uintptr

	// var zs = 31
	// var cs int = 32
	// fmt.Printf("%d, %d\n", zs, cs)   // %d 占位符表示打印十进制 , \n 表示换行
	// fmt.Printf("%T", zs)
	// 八进制 0xx (后两位xx的值小于8)
	// 十六进制 0xxx （后两位xx 的值为 0-F）
	// 十进制 0-9
	// 算数运算， + - * / % ++ --    （ ++ 表示自增长1，--表示自减少1）
	a, b := 2, 4
	fmt.Println(a + b)
	fmt.Println(a - b) // 0
	fmt.Println(a * b) // 2
	fmt.Println(a / b) // 0  （除运算值为0时，程序编译后运行时会报错 runtime error）
	fmt.Println(a % b)
	a++
	b--
	fmt.Println(a, b) // 3, 3

	// 关系运算 >  <  >=  <= != ==
	fmt.Println(a > b)  // false
	fmt.Println(a < b)  // false
	fmt.Println(a >= b) //true
	fmt.Println(a <= b) // true
	fmt.Println(a != b) // false
	fmt.Println(a == b) //true

	// 位运算
	/* & | ^

	注意：int ,int8 ,16,32,64  不同类型整形之间不能进行算数运算
	*/

	var (
		i   int   = 31
		i32 int32 = 31
		//i64 int64 = 31
	)
	// fmt.Println(i + i32)   错误示例，i 为 int 类型，i32为 int32类型，运行运算会报错
	fmt.Println(i + int(i32)) // 将i32转换为int类型后进行运算正常

	// 进制占位符 %d %b %o %x %c %U 十进制  二进制  八进制  十六进制  byte类型  码点类型
	var (
		achar byte = 'A'  // 在Ascill编码中，A经过编码转换后为 65 ， 此处Print打印 achar 变量时，打印的值为 65
		aint byte = 65   // byte类型为Ascill编码， 65 
		unicodepoint rune = '中'  // Print打印时，打印的是该中文在unicode编码中转换后的值，中 在编码中时20013
	)
	fmt.Println(achar,aint,unicodepoint)
}
