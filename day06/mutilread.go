package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file1, err := os.Open("./test/w1.go")
	file2, err := os.Open("./test/w2.go")
	if err != nil {
		return
	}

	reader := io.MultiReader(file1, file2)   // 同时将多个文件传递给 multiRead ， 按顺序读取文件内容

	ctx := make([]byte, 5)
	context := make([]byte, 0, 1024)
	for {
		n, err := reader.Read(ctx)   // 将多个文件的内容按顺序读取到ctx, 赋值给 context
		if err != nil {
			break
		}
		context = append(context, ctx[:n]...)
	}
	fmt.Println(string(context))
}
