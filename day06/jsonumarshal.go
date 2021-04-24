package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type Task struct {
	Id        int64
	User      string
	Name      string
	Starttime *time.Time
	Endtime   *time.Time
}

func main() {

	Tasks := []*Task{}

	read, err := ioutil.ReadFile("task.json") // 通过ioutil 包将文件读取到 read 字节切片内
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(read, &Tasks) // 使用json.unmarshal 函数将 read内的内容，解析到&Task 内

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range Tasks { // print Tasks的内容
		fmt.Printf("%#v\n", v)
	}

}
