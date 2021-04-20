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
	starttime time.Time
	endtime   time.Time
}

const TimeLayte string = "2006-01-02 15:04:05"

func timestr(t time.Time) string {
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
	task := &Task{
		id:        1,
		user:      "qing",
		name:      "demo",
		starttime: time.Now(),
		endtime:   time.Now().Add(24 * time.Hour),
	}

	files, err := os.Create("task.log")
	if err != nil {
		return
	}
	defer files.Close()
	task.WriteTo(files)
}
