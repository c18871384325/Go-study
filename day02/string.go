package main

import "fmt"
import "unicode/utf8"
import "strconv"

func main() {

	// 字节切片和字符串的转换
	ascii := "abc我爱中华人民共和国"

	fmt.Println([]byte(ascii))   // 字节切片方式打印string类型
	fmt.Println([]rune(ascii))   // 以 unicode 字节切片形式打印string类型

	// 求字符串类型变量的字符个数，如果包含中文需要用一下方式，因为英文一个占1字节，但是中文会占用3个字节左右，len(string-name) 获取的是变量字节的长度，如果包含中文并不准确
	fmt.Println(len(ascii))  //   30   ， 实际ascii 变量字符个数为12，此处因为有中文，显示为30
	fmt.Println(len([]rune(ascii)))   // 12  , 通过将string类型变量转为rune类型的切片后打印，通过len获取到的长度就是实际变量字符的个数

	// 另外一种获取字符个数的实现方式
	fmt.Println(utf8.RuneCountInString(ascii))    // 该方式需导入包  import "unicode/utf8"


	// 打印字符串
	fmt.Println(string([]byte(ascii)))   // abc我爱中华人民共和国
	fmt.Println(string([]rune(ascii)))   // abc我爱中华人民共和国


	// 类型转换  int, float, bool
	// int to string
	fmt.Println(strconv.Itoa('a'))   // Int类型转换到string类型
	ch, err := strconv.Atoi("97")  // string 转换到int类型
	fmt.Println(ch, err)


	// float to string
	fmt.Println(strconv.FormatFloat(3.1415926, 'f', 10, 64) ) //   3.xxx 表示要转换的值，f表示float类型，10表示转换的字节长度，64表示float64类型
	// 3.1415926000    ， 不足10位的字节用float64的零值替代

	pi, err := (strconv.ParseFloat("3.1415926", 64))  // 字符串转float， 3.xxx表示字符串的值，64表示转换为float64
	fmt.Println(pi, err)


	// bool to string
	fmt.Println(strconv.FormatBool(true))  // 将bool类型的true转换为字符串
	fmt.Println(strconv.ParseBool("true"))  // 将字符串类型的true转换为bool类型的值


	// string to int
	b, err := strconv.ParseInt("57", 10, 10)  // 5表示字符串的值，10表示转换10个字节的字符串
	fmt.Printf("%T, %#v\n", b, b)

	
	fmt.Println(strconv.FormatInt(156, 10))   // int转换为字符串，156表示值，10表示10进制
}