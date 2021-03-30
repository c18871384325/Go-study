package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		host string
		port int
		h    bool
		help bool
	)

	// -H host  -P port
	// -h --help

	// 解析变量的指针，命令行中指定的参数名，默认值，帮助
	//  用户输入 -H value ， 将用户的参数值传递给变量 host ，未输入则默认值为127.0.0.1， 帮助信息显示 连接地址

	flag.StringVar(&host, "H", "127.0.0.1", "连接地址")
	flag.IntVar(&port, "P", 22, "连接端口")
	flag.BoolVar(&h, "h", false, "帮助")
	flag.BoolVar(&help, "help", false, "帮助")

	flag.Usage = func() {
		fmt.Println("usage: testflag [-H 127.0.0.1] [-P 22]")
		flag.PrintDefaults() // print 默认的帮助信息，会将flag.xxVar定义的信息反馈给用户

		/* D:\src\Go-study\day04\flags>flags.exe -h
		usage: testflag [-H 127.0.0.1] [-P 22]
		-H string
				连接地址 (default "127.0.0.1")
		-P int
				连接端口 (default 22)
		-h    帮助
		-help
				帮助 */
	}

	flag.Parse() // 解析

	if h || help {
		flag.Usage()
		return
	}

	fmt.Println(host, port, h, help)
	fmt.Println(flag.NArg()) // 返回未处理的参数数量，一般为flag未定义的参数

	fmt.Printf("%#v\n", flag.Args()) // 返回未处理的参数值

	// 验证
	// flags.exe kk huawei
	// 127.0.0.1 22 false false
	// 2
	// []string{"kk", "huawei"}

}
