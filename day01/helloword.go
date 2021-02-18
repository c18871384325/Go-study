package main

/* 第一个代码：

package main

import "fmt"

func main() {
	fmt.Println("hello word")
}


package main  （声明一个包， main 为包名，如果需要编译为二进制程序，必须为 main 包）
import "fmt" （导入fmt包）

func main() {
	fmt.Println("hello word")
}
使用main函数作为入口，执行函数内的代码，fmt.Println 调用fmt包内的Println函数打印内容到终端

查看编译的详细过程：
go build -x helloword.go

指定编译出来的二进制文件名称为test.exe  -o 选项
go build -x -o test.exe helloword.go

直接编译后运行（支持 -x 查看编译执行过程）
go run helloword.go

go run helloword.go -n   （打印过程，但编译出的程序不会运行 ，-n 选项）

手动格式化代码
go fmt helloword.go  (将代码格式统一为go语言格式，编辑器保存时会自动格式化）



编程语言的注释：
	1.行注释
		//  （//开头的行后内容被注释）
	2.快注释
		/*  
			*/   /* */  //注释包含的内容被注释，支持多行）

import "fmt"  // 导入fmt包

func main() {  // main 函数作为程序执行的入口，执行函数内的代码
	fmt.Println("hello word")  // 表示调用fmt包内的println函数打印内容到终端

}
