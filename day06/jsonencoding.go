package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	Id        int
	User      string
	Name      string
	Starttime *time.Time
	Endtime   *time.Time
}

func main() {

	now := time.Now()
	endnow := time.Now().Add(24 * time.Hour)

	tasks := []*Task{
		{
			Id:        1,
			User:      "cisco",
			Name:      "3660",
			Starttime: &now,
			Endtime:   &endnow,
		},
		{
			Id:        2,
			User:      "h3c",
			Name:      "5700",
			Starttime: &now,
			Endtime:   &endnow,
		},
	}

	file, err := os.Create("json.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	encode := json.NewEncoder(file) // 定义新的json编码

	encode.Encode(tasks) // 将tasks内容以json编码的方式写入到文件，该种方式占用文件空间比 marshal要小，marshal结构化易读，但是占用空间大些

}
