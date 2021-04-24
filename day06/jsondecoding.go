package main

import (
	"encoding/json"
	"time"
	"os"
	"fmt"
)

type Task struct {
	Id        int
	User      string
	Name      string
	Starttime *time.Time
	Endtime   *time.Time
}

func main() {
	Tasks := []*Task{}

	file, err := os.Open("json.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	decode := json.NewDecoder(file)  // 创建新的json 解码，并将 io.Reader 传递到新的解码

	decode.Decode(&Tasks)  // 将解码写入到 &Tasks  , 注意： 此处需要传递指针

	for _, v := range Tasks {
		fmt.Printf("%#v\n", v)
	}

	


}