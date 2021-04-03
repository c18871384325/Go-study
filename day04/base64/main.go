package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// base64 加密解密

	fmt.Println(base64.StdEncoding.EncodeToString([]byte("This is go 语言测试"))) // 对字符串切片进行base64 的加密
	// VGhpcyBpcyBnbyDor63oqIDmtYvor5U=

	txt, _ := base64.StdEncoding.DecodeString("VGhpcyBpcyBnbyDor63oqIDmtYvor5U=") // 对base64加密后的密文进行解码
	fmt.Printf(string(txt))                                                       // 使用字符串显示解密后的内容
	fmt.Println()

	// 对URL进行base64加解密
	fmt.Println(base64.URLEncoding.EncodeToString([]byte("This is URL 加密"))) // 对URL进行base64加密
	// VGhpcyBpcyBVUkwg5Yqg5a-G

	txturl, _ := base64.URLEncoding.DecodeString("VGhpcyBpcyBVUkwg5Yqg5a-G") // 对URL加密的内容进行解密

	fmt.Println(string(txturl))

	// 标准
	fmt.Println(base64.RawStdEncoding.EncodeToString([]byte("我有一个梦想，当一个音乐家")))

	//url
	fmt.Println(base64.RawURLEncoding.EncodeToString([]byte("我有一个URL，自己编写的")))

}
