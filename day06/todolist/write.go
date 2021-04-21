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

func (Tasks *Task) WriteTo(file *os.File) error {
	_, err := fmt.Fprintf(file, "%d,%s,%s,%s,%s\n",
		Tasks.id,
		Tasks.name,
		Tasks.user,
		timestr(Tasks.starttime),
		timestr(Tasks.endtime),
	)
	return err
}

func main() {

	start := time.Now()
	starttime := &start
	end := time.Now().Add(10 * time.Hour)
	endtime := &end

	Tasks := &[]Task{}
	Tasks = &[]Task{
		{
			id:        1,
			user:      "qing",
			name:      "整理笔记1",
			starttime: starttime,
			endtime:   endtime, //time.Now().Add(24 * time.Hour),
		},
		{
			id:        2,
			user:      "cheng",
			name:      "整理笔记2",
			starttime: starttime,
			endtime:   endtime,
		},
	}

	files, err := os.Create("task.log")
	if err != nil {
		return
	}
	defer files.Close()
	for _, task := range *Tasks {
		task.WriteTo(files)
	}

}
