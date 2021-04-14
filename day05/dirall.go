package main

import (
	"fmt"
	"os"
	"strings"
)

// print 线上目录下的所有目录及目录下的文件，并指明目录，并且显示文件内容

func filterfile(file string) bool { // 对文件名进行过滤，当为.go 结尾时才执行操作，否则不执行
	if strings.HasSuffix(file, ".go") {
		return true
	} else {
		return false
	}

}

func openfile(file string) {
	rfile, err := os.Open(file)
	if err != nil {
		return
	}
	defer rfile.Close()
	txt := make([]byte, 0, 1024*1024)
	ctx := make([]byte, 1024)
	for {
		n, err := rfile.Read(ctx)
		if err != nil {
			break
		}
		txt = append(txt, ctx[:n]...) //  ctx[:n]...  为切片的值
	}
	if string(txt) == "" {
		return
	}
	fmt.Println(string(txt))

}

func ReadDir(path string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	namesall, err := file.Readdirnames(-1)
	if err != nil {
		return
	}
	for _, names := range namesall {
		fpath := path + "/" + names
		if fileinfo, err := os.Stat(fpath); err == nil {
			if fileinfo.IsDir() {
				//fmt.Println("Dir: ", fpath)
				ReadDir(fpath)
			} else {
				if filterfile(fpath) {
					fmt.Println(fpath)
					openfile(fpath)
				}
			}
		}
	}
}

func main() {
	ReadDir("dirdemo")
}
