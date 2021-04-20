package main

import (
	"fmt"
	"os"
	//"time"
)

func main() {
	file, err := os.Create("./test/fmtF.txt")
	if err != nil {
		return
	}
	defer file.Close()
	fmt.Fprintf(file, "%s", "are you beeter full") // fmt.Fprintf，支持将内容写入到io.write, file 为io输入流，%s 为格式，与printf 类似， 最后为写入的内容
}
