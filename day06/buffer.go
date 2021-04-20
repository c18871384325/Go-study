package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	buffer1 := bytes.NewBuffer([]byte("abc,"))
	buffer2 := bytes.NewBufferString("123,")

	buffer1.Write([]byte("def,")) // 写入字节切片
	buffer1.Write([]byte("123,"))
	buffer2.WriteString("456,") // 写入字符串

	ctx1 := make([]byte, 3)
	n, _ := buffer1.Read(ctx1)   // 此处因为ctx大小为5, 因此一次只读取5个字节
	fmt.Println(n, string(ctx1)) // 5 abcde

	n, _ = buffer1.Read(ctx1)    // 第二次读取时是从上次读取后光标所在的位置开始读取
	fmt.Println(n, string(ctx1)) // 5 fghij

	text, _ := buffer2.ReadBytes(',') // 当碰到指定的字符后停止读取， 此处碰到 . 后停止读取
	fmt.Println(string(text))         //  123,

	text, _ = buffer2.ReadBytes(',')
	fmt.Println(string(text)) // 456,

	text2, _ := buffer1.ReadString(',')
	fmt.Println(text2)

	// 将buffer流中剩余内容转换为字节切片
	fmt.Println(buffer2.String())

	buffer1.Reset() // 清空buffer中的内容
	buffer2.WriteTo(os.Stdout)

}
