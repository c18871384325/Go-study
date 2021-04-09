package main

import (
	"fmt"
	"flag"
)


// flag  go语言命令行处理库，将用户传入的参数值写入到指定变量，并执行
func main() {
	var (
		host string
		port int
		h bool
		help bool
	)

	// &host表示将用户参数传入到指定的变量， H 表示用户使用的参数 -H or --H  ,  127.0.0.1 为该参数的默认值（用户未输入时），连接地址表示该参数的帮助信息
	flag.StringVar(&host, "H", "127.0.0.1", "连接地址")
	flag.IntVar(&port, "P", 22, "连接端口")
	flag.BoolVar(&h, "h", false, "帮助信息")
	flag.BoolVar(&help, "help", false, "帮助信息")

	flag.Usage = func () {   // 使用指南
		fmt.Println("Usage: testflag [-H 127.0.0.1] [-P 22]")
		flag.PrintDefaults()
	}

	flag.Parse()

	if h || help {
		flag.Usage()
		return
	}

	fmt.Println(host,port,h,help)
	fmt.Println(flag.NArg())   // 未处理的参数的数量

	fmt.Println("%#v\n", flag.Args())
}