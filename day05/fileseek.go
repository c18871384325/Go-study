package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("filedemo", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)

	if err != nil {
		return
	}

	defer file.Close()
	file.Seek(3, os.SEEK_SET) // seek 函数作用为移动文件操作光标的位置， 3 为偏移量， os.SEEK_SET 表示从文件初始位置开始向后便宜3个字节，因此后面的 Read 是从第四个字节开始读取，print的结果为456
	ctx := make([]byte, 3)

	n, err := file.Read(ctx)
	fmt.Println(n, err, string(ctx)) // 3 <nil> 456

	fmt.Println(file.Seek(0, os.SEEK_CUR)) // 6 <nil>     os.SEEK_CUR 为查看当前光标的位置，上面向后偏移了3个，又读取了3个，此处为6

	file.Write([]byte("qwer")) // 向后追加内容，基于 os.O_APPEND 函数

	// Go语言中，默认不需要保存文件，文件关闭时会自动进行保存操作

}
