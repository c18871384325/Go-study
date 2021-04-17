package main

import (
	"fmt"
	"todolist/password"
)

type Task struct {
	id        int
	taskname  string
	starttime string
	endtime   string
}

var task []Task

func main() {

	task = make([]Task, 0, 10)
	for i := 0; i < 3; i++ {
		if password.Password() == true {
			for {
				var operator string
				fmt.Print("请输入操作： ")
				fmt.Scan(&operator)
				switch operator {
				case "add":
					Taskadd()
				case "delete":
					Taskdelete()
				case "edit":
					Taskedit()
				case "search":
					Tasksearch()
				case "exit":
					break
				default:
					fmt.Println("input error, please input add/delete/edit/search/exit")
				}
			}
		} else {
			fmt.Println("密码输入错误，请重新输入： ")
			continue
		}
	}
}
func taskid() int {
	var id int
	for _, v := range task {
		if v.id == 0 {
			id = 1
		} else {
			id = v.id + 1
		}
	}
	return id
}

func TaskPrint(print Task) {
	fmt.Println("任务ID： ", print.id)
	fmt.Println("任务名称： ", print.taskname)
	fmt.Println("任务开始时间： ", print.starttime)
	fmt.Println("任务结束时间： ", print.endtime)
}

func Taskadd() {
	var taskname string
	var starttime string
	var endtime string
	fmt.Print("请输入任务名称： ")
	fmt.Scan(&taskname)
	fmt.Print("请输入任务开始时间： ")
	fmt.Scan(&starttime)
	fmt.Print("请输入任务结束时间： ")
	fmt.Scan(&endtime)
	tadd := Task{
		id:        taskid(),
		taskname:  taskname,
		starttime: starttime,
		endtime:   endtime,
	}

	task = append(task, tadd)
}

func Taskedit() {
	var edit string
	var newname string
	var newstarttime string
	var newendtime string
	fmt.Print("请输入修改的任务： ")
	fmt.Scan(&edit)
	fmt.Print("请输入任务名： ")
	fmt.Scan(&newname)
	fmt.Print("请输入任务开始时间： ")
	fmt.Scan(&newstarttime)
	fmt.Print("请输入任务结束时间： ")
	fmt.Scan(&newendtime)
	for i, v := range task {
		if v.taskname == edit {
			newtask := Task{
				id:        v.id,
				taskname:  newname,
				starttime: newstarttime,
				endtime:   newendtime,
			}
			task[i] = newtask
			TaskPrint(task[i])
		}
	}
}

func Taskdelete() {
	var delete string
	fmt.Print("请输入删除的任务： ")
	fmt.Scan(&delete)
	for i, v := range task {
		if v.taskname == delete {
			if i == 0 {
				task = task[1:]
			} else if i == len(task)-1 {
				task = task[:i]
			} else {
				dtask := task[i+1:]
				copy(task[i:], dtask)
				task = task[:len(task)-1]
			}
		}
	}
}

func Tasksearch() {
	for i := 0; i < 3; i++ {
		var search string
		fmt.Print("请输入查询的任务： ")
		fmt.Scan(&search)
		for i, v := range task {
			if v.taskname == search {
				TaskPrint(task[i])
				return
			}
		}
		fmt.Println("task is not exists, please again input")
	}
	fmt.Println("input error beyond three cont , please again executed operator")
	return
}
