package main

import (
	"fmt"
	"os"
	"time"
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

func (Tasks *Task) WriteTo(file *os.File) error {
	_, err := fmt.Fprintf(file, "%d,%s,%s,%s,%s\n",
		Tasks.Id,
		Tasks.Name,
		Tasks.User,
		timestr(Tasks.Starttime),
		timestr(Tasks.Endtime),
	)
	return err
}

func main() {

	start := time.Now()
	//starttime := &start
	end := time.Now().Add(10 * time.Hour)
	//endtime := &end

	Tasks := &[]Task{}
	Tasks = &[]Task{
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

	files, err := os.Create("task.log")
	if err != nil {
		return
	}
	defer files.Close()
	for _, task := range *Tasks {
		task.WriteTo(files)
	}

}
