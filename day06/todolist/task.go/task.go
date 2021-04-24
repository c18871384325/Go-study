package main

import (
	"time"
)

type Task struct {
	Id        int64
	User      string
	Name      string
	Starttime *time.Time
	Endtime   *time.Time
}


func AddTask() *Task {
	var Name string
	var User string
	var Starttime string
	var Endtime string
	fmt.Print("请输入用户： ")
	fmt.Scan(&User)
	fmt.Print("请输入名称： ")
	fmt.Scan(&Name)
	fmt.Print("请输入开始时间： ")
	fmt.Scan(&Starttime)
	fmt.Print("请输入结束时间： ")
	fmt.Scan(&User)
}