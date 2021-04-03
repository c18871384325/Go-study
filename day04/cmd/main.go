package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	ping := exec.Command("ping", "-c", "3", "www.baidu.com")   // 定义调用系统命令,仅仅定义,并不执行

	// value, _ := ping.Output()  // output 为执行系统命令调用,并返回执行结果

	// fmt.Println(string(value))   // 打印 系统命令  ping 的值

	output, err := ping.StdoutPipe()

	ping.Start()
	fmt.Println(err)
	io.Copy(os.Stdout, output)   // 将output获取的执行结果输出流拷贝到os标准输出,通过该种方式实现返回值的输出
	ping.Wait()   // 命令执行必须wait

}