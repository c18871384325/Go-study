package main

import (
	"encoding/hex"
	"fmt"
	//"math/rand"
)

func main() {

	// hex 实际是将字符切片转换为16进制表示方法

	fmt.Printf("%x\n", []byte("my name is czhecheng"))           // 将字符切片编码
	fmt.Println(hex.EncodeToString([]byte("This is 徐青")))        // 将字符切片转换为hex编码
	txt, _ := (hex.DecodeString("5468697320697320e5be90e99d92")) // 将hex解码
	fmt.Println(string(txt))
}
