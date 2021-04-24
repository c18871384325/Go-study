package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"time"
	"strconv"
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

	file, err := os.Create("csv.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	

	write := bufio.NewWriter(file)

	defer write.Flush()    // 注意，使用 bufio 写入时，需要延时关闭使用 flush ，将写入到bufio的内容 flush 到文件内，否则写入为空

	CsvWrite := csv.NewWriter(write)

	CsvWrite.Write([]string{"ID", "用户", "名称", "开始时间", "结束时间"})

	for _, task := range Tasks {
		CsvWrite.Write(
			[]string{
				strconv.Itoa(int(task.Id)),
				task.User,
				task.Name,
				timestr(task.Starttime),
				timestr(task.Endtime),

			},
		)
	}
}
