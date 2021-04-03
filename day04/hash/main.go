package hash

import (
    "crypto/md5"
    "encoding/hex"
    "fmt"
)
func main() {
    // hash 一般用在签名的时候，hash算法运算后是不可逆的（不可通过密文恢复原文）
    // md5,sha1,sha256,sha512 都是hash算法
    // des, aes 等算法是可逆的，使用私钥可以恢复原文
    md5demo := md5.Sum([]byte("我是md5测试")) // md5加密字符串切片内容
    fmt.Printf("%x\n", md5demo)           // 57f967ce561f7c1c3d37dd1f9f15884a
    hasher := md5.New()
    hasher.Write([]byte("我是"))
    hasher.Write([]byte("md5测试"))
    fmt.Println(hex.EncodeToString(hasher.Sum(nil))) // 结果与md5加密相同
    // 加盐 hash 或 md5算法

    "crypto/sha1"
    "crypto/sha256"
    "crypto/sha512"
    "encoding/hex"
    "fmt"
    "math/rand"
    "time"
)
func init() {
    rand.Seed(time.Now().Unix()) // 在程序运行前初始化, 随机数种子初始化
}
func randString(n int) (string, string) { //函数获取随机数
    chars := "abcdefghijklmnopqrstuvwxyzABCDE"
    rt1 := make([]byte, 0, n)
    rt2 := make([]byte, n, n)
    for i := 0; i < n; i++ {
        rt1 = append(rt1, chars[rand.Intn(len(chars))])
        rt2[i] = chars[rand.Intn(len(chars))]
    }
    return string(rt1), string(rt2)
}
func md5String(text string, salt string) string { // 通过函数调用返回加盐后的md5值
    bytes := []byte(salt)
    bytes = append(bytes, ':')
    bytes = append(bytes, []byte(text)...)
    return fmt.Sprintf("%x", md5.Sum(bytes))
}
func main() {
    salt1, salt2 := randString(5)
    fmt.Println(salt1, salt2)
    fmt.Println(md5String("this is czhecheng", salt1)) // 加盐后提取字符串的md5值
    // sha1 sha256, sha512 加密 , 需导入加密包
    fmt.Printf("%x\n", sha1.Sum([]byte("THis is czhecheng")))      // 使用sha1 对字符切片进行加密
    fmt.Printf("%x\n", sha256.Sum256([]byte("This is czhecheng"))) // 使用sha256 对字符切片进行加密
    fmt.Printf("%x\n", sha512.Sum512([]byte("This is czhecheng"))) // 使用sha512 对字符切片进行加密
    sha256hasher := sha256.New()
    sha256hasher.Write([]byte("This is"))
    sha256hasher.Write([]byte("czhecheng"))
    fmt.Println(hex.EncodeToString(sha256hasher.Sum(nil))) // md5加密
}
