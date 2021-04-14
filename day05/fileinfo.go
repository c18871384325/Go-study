package main

import (
	"fmt"
	"os"
)

// 查看文件信息

func main() {
	fileinfo, err := os.Stat("filedemo")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fileinfo.Size()) // 文件大小，默认为字节显示
	fmt.Println(fileinfo.Name())
	fmt.Println(fileinfo.Mode())
	fmt.Println(fileinfo.IsDir())
	fmt.Println(fileinfo.ModTime())

	/*137
		filedemo
		-rw-rw-rw-
		false
		2021-04-13 09:17:35.6384772 +0800 CST

		type FileInfo interface {
	        Name() string       // base name of the file
	        Size() int64        // length in bytes for regular files; system-dependent for others
	        Mode() FileMode     // file mode bits
	        ModTime() time.Time // modification time
	        IsDir() bool        // abbreviation for Mode().IsDir()
	        Sys() interface{}   // underlying data source (can return nil)
	}
	*/

}
