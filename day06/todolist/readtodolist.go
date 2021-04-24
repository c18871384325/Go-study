package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const TimeLayteOut string = "2006-01-02 15:04:05"

type Task struct {
	Id        int64
	User      string
	Name      string
	Starttime *time.Time
	Endtime   *time.Time
}

func Parsestr(line string) (*Task, error) {
	node := strings.Split(line, ",")
	if len(node) < 5 {
		return nil, errors.New("字符串数量不正确")
	}

	Id, err := strconv.Atoi(node[0])
	if err != nil {
		return nil, err
	}

	User := node[1]

	Name := node[2]

	var StartTime, EndTime *time.Time

	if node[3] != "" {
		if t, err := time.Parse(TimeLayteOut, node[3]); err != nil {
			return nil, err
		} else {
			StartTime = &t
		}
	}

	if node[4] != "" {
		if t, err := time.Parse(TimeLayteOut, node[3]); err != nil {
			return nil, err
		} else {
			EndTime = &t
		}
	}

	task := &Task{
		Id:        int64(Id),
		User:      User,
		Name:      Name,
		Starttime: StartTime,
		Endtime:   EndTime,
	}

	return task, nil

}






func main() {

	file, err := os.Open("task.log")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	tasks := make([]*Task, 0, 10)

	for scanner.Scan() {
		task, err := Parsestr(scanner.Text())
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}
		tasks = append(tasks, task)
	}
	for _, v := range tasks {
		fmt.Printf("%#v\n", v)
	}

}
