package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix()) // 在程序运行前初始化, 随机数种子初始化
}

func randString(n int) (string, string) { //函数获取随机数
	chars := "abcdefghijklmnopqrstuvwxyzABCDE"

	rt1 := make([]byte, 0, n)
	rt2 := make([]byte, n, n)

	for i := 0; i < n; i++ {
		rt1 = append(rt1, chars[rand.Intn(len(chars))])
		rt2[i] = chars[rand.Intn(len(chars))]

	}

	return string(rt1), string(rt2)

}

func md5String(text string, salt string) string { // 通过函数调用返回加盐后的md5值
	bytes := []byte(salt)
	bytes = append(bytes, ':')
	bytes = append(bytes, []byte(text)...)

	return fmt.Sprintf("%x", md5.Sum(bytes))
}

func main() {

	salt1, salt2 := randString(5)

	fmt.Println(salt1, salt2)

	fmt.Println(md5String("this is czhecheng", salt1)) // 加盐后提取字符串的md5值

	// sha1 sha256, sha512 加密 , 需导入加密包
	fmt.Printf("%x\n", sha1.Sum([]byte("THis is czhecheng")))      // 使用sha1 对字符切片进行加密
	fmt.Printf("%x\n", sha256.Sum256([]byte("This is czhecheng"))) // 使用sha256 对字符切片进行加密
	fmt.Printf("%x\n", sha512.Sum512([]byte("This is czhecheng"))) // 使用sha512 对字符切片进行加密

	sha256hasher := sha256.New()
	sha256hasher.Write([]byte("This is"))
	sha256hasher.Write([]byte("czhecheng"))
	fmt.Println(hex.EncodeToString(sha256hasher.Sum(nil))) // md5加密
}
