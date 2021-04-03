package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {

	// hash 一般用在签名的时候，hash算法运算后是不可逆的（不可通过密文恢复原文）
	// md5,sha1,sha256,sha512 都是hash算法

	// des, aes 等算法是可逆的，使用私钥可以恢复原文

	md5demo := md5.Sum([]byte("我是md5测试")) // md5加密字符串切片内容
	fmt.Printf("%x\n", md5demo)           // 57f967ce561f7c1c3d37dd1f9f15884a

	hasher := md5.New()
	hasher.Write([]byte("我是"))
	hasher.Write([]byte("md5测试"))
	fmt.Println(hex.EncodeToString(hasher.Sum(nil))) // 结果与md5加密相同

	// 加盐 hash 或 md5算法
	

}
