package main

import (
	"fmt"
	"time"
)

type task struct {
	id        int
	name      string
	starttime time.Time
	endtime   time.Time
	status    int
	user      string
}

func main() {
	// type task int    // 定义一个结构体task ，结构体类型为 int
	// var cnt task
	// fmt.Printf("%T, %#v", cnt, cnt)
	var demo task

	demo = task{1, "tencent", time.Now(), time.Now().Add(24 * time.Hour), 1, "huawei"}
	fmt.Printf("%T, %#v", demo, demo)

}
