package main

import (
	"fmt"
	"os"
)

func main() {
	os.Mkdir("dirdemo", os.ModePerm)                     // 创建目录 dirdemo , 权限为 os.ModePerm 为777  ，可自定义 如： 0644
	fmt.Println(os.MkdirAll("dirdemo/a/b", os.ModePerm)) // 创建多个目录，当子目录不存在时自动创建
	os.Remove("dirdemo/a/b")                             // 移除目录，注意，当目录下不为空时无法移除
	os.RemoveAll("dirdemo")                              //   removeall  , 移除目录，不管目录下是否有文件都移除
}
