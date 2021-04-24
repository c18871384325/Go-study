package main 

import (
	"encoding/json"
	"bytes"
	"os"
	"time"
	"fmt"
)

type Task struct {
	Id        int64
	User      string
	Name      string
	Starttime *time.Time
	Endtime   *time.Time
}

const TimeLayte string = "2006-01-02 15:04:05"

func timestr(t *time.Time) string {
	return t.Format(TimeLayte)
}

func main() {

	start := time.Now()
	end := start.Add(24 * time.Hour)

	Tasks := []*Task{
		{
			Id:        1,
			User:      "qing",
			Name:      "整理笔记1",
			Starttime: &start,
			Endtime:   &end, //time.Now().Add(24 * time.Hour),
		},
		{
			Id:        2,
			User:      "cheng",
			Name:      "整理笔记2",
			Starttime: &start,
			Endtime:   &end,
		},
	}

	ctx, _ := json.Marshal(Tasks)   // 将Tasks 内容以字节切片写入到ctx

	var buffer bytes.Buffer

	json.Indent(&buffer, ctx, "", "\t")   // 将 ctx 的内容写入到 buffer， 以json格式，   ""  为前缀，  \t 是格式

	file, err := os.Create("task.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	buffer.WriteTo(file)
}