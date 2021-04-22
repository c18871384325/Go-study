package main

import (
	"encoding/gob"
	"fmt"
	"time"
	"os"
)

type Task struct{
	Id int
	User string
	Name string
	Starttime *time.Time
	Endtime *time.Time
}



func main() {

	now := time.Now()
	endnow := time.Now().Add(24 * time.Hour)

	tasks := []*Task{
		{
			Id: 1,
			User: "cisco",
			Name: "3660",
			Starttime: &now,
			Endtime: &endnow,
		},
		{
			Id: 2,
			User: "h3c",
			Name: "5700",
			Starttime: &now,
			Endtime: &endnow,
		},
	}

	file, err := os.Create("task.gob")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	encoding := gob.NewEncoder(file)
	encoding.Encode(tasks)

}