package main

import (
	"fmt"
	"os"
	"time"
)

type Task struct {
	id        int64
	user      string
	name      string
	starttime *time.Time
	endtime   *time.Time
}

const TimeLayte string = "2006-01-02 15:04:05"

func timestr(t *time.Time) string {
	return t.Format(TimeLayte)
}

func (task *Task) WriteTo(file *os.File) error {
	_, err := fmt.Fprintf(file, "%d,%s,%s,%s,%s\n",
		task.id,
		task.name,
		task.user,
		timestr(task.starttime),
		timestr(task.endtime),
	)
	return err
}

func main() {

	start := time.Now()
	starttime := &start
	end := time.Now().Add(10 * time.Hour)
	endtime := &end

	task := &[]Task{{
		id:        1,
		user:      "qing",
		name:      "整理笔记1",
		starttime: starttime,
		endtime:   endtime,      //time.Now().Add(24 * time.Hour),
	},
	{
		id: 2,
		user: "cheng",
		name: "整理笔记2",
		starttime: starttime,
		endtime: endtime,
	},
}

	files, err := os.Create("task.log")
	if err != nil {
		return
	}
	defer files.Close()
	task.WriteTo(files)
}
