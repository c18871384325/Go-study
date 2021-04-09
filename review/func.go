package main

import (
	"fmt"
)

// 函数，一组代码的集合，接受参数，并通过代码对参数处理，返回结果

// 示例： 
func add(n1 int, n2 int) int {
	nums := n1 + n2
	return nums
}


// 函数返回函数
func demoreturn(base int) func (n int) int {
	return func (n int ) int {
		return n + base
	} 
}

// 多参数的函数, args多参数实际传入的是一个切片，此处为int类型的切片
func manyargs(n1 int, n2 int, args ...int) int {
	total := n1 + n2
	for _, v := range args {
		total += v
	}
	return total
}

// 函数的 defer 延时处理使用, 在return 调用之后，函数退出之前执行，函数遇到panic时，退出前也会执行defer
func demodefer(redish string) string {
	defer func () {
		fmt.Println("This is defer exec result")
	}()

	with := []byte(redish)
	for i := 0; i < len(with); i++ {
		if with[i] == 'h' {
			with[i] = 'm'
		}
	}
	return string(with)
}


// point , 在函数内修改外部变量的值
func changepoint(cpoint *int) {
	*cpoint = *cpoint + 1
}


// panic 错误处理
func demopanci() (err error) {
	defer func () {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
		}
	}()
	
	fmt.Println("befor")
	panic("This is single panic")
	fmt.Println("after")   // panic 只有的代码并不会执行
	return 
}


func main() {
	
	// 调用第一个add函数
	fmt.Println(add(3,4))   // 返回值为函数add处理后的结果，为7

	// 调用函数的返回值为函数
	returnfunc := demoreturn(2)
	fmt.Println(returnfunc(5))   // 结果为 7
	

	// 调用地三个多参数的函数
	fmt.Println(manyargs(1,2))
	fmt.Println(manyargs(1,2,3))
	fmt.Println(manyargs(1,2,3,4))
	fmt.Println(manyargs(1,2,3,4,5))

	
	// 调用第四个带defer的函数
	fmt.Println(demodefer("huawei"))
	// This is defer exec result
	// muawei
	
	// 调用第五个point修改外部变量的函数
	value := 3
	changepoint(&value)
	fmt.Println(value)
	// 4

	// 调用panic错误处理
	fmt.Println("befor")
	err := demopanci()
	fmt.Println("after", err)
	// befor
	// befor
	// after This is single panic

}	
