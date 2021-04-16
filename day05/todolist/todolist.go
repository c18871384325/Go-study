package main

import (
	"fmt"
	//"time"
)

type Task struct {
	id        int
	taskname  string
	starttime string
	endtime   string
}

func taskid(task []Task) int {
	var id int
	for _, v := range task {
		if v.id == 0 {
			id = 0
		} else {
			id = v.id + 1
		}
	}
	return id
}

func taskadd(task []Task) {
	var taskname string
	var starttime string
	var endtime string
	fmt.Print("请输入任务名称： ")
	fmt.Scan(&taskname)
	fmt.Print("请输入任务开始时间： ")
	fmt.Scan(&starttime)
	fmt.Print("请输入任务结束时间： ")
	fmt.Scan(&endtime)
	tadd:= Task{
		id:        taskid(task),
		taskname:  taskname,
		starttime: starttime,
		endtime:   endtime,
	}

	task = append(task, tadd)
	fmt.Println(task)
}

func taskdelete(task []Task) {
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
	fmt.Println(task)
}

func taskedit(task []Task) {
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
					id: v.id,
					taskname: newname,
					starttime: newstarttime,
					endtime: newendtime,
				}
				task[i] = newtask
			}
		}
		fmt.Println(task)
	}

func main() {

	task := make([]Task, 0, 10)
	for i := 0; i < 3; i++ {
		taskadd(task)
	}

	//taskdelete(task)
	// var inputpasswd string
	// fmt.Println("请输入密码： ")
	// fmt.Scan(&inputpasswd)

	// var initpasswd string
	// fmt.Println("请初始化密码： ")
	// fmt.Scan(&initpasswd)

	// if inputpasswd(inputpasswd) == password() {
	// 	var operator string
	// 	fmt.Println("请输入操作： ")
	// 	fmt.Scan(&operator)
	// 	switch operator {
	// 	case "add":
	// 			taskadd()
	// 	case "delete":
	// 			taskdelete()
	// 	case "edit":
	// 			taskedit()
	// 	case "search":
	// 			tasksearch()
	// 	case "exit":
	// 			return
	// 	default:
	// 		fmt.Println("input error, please input add/delete/edit/search/exit")
	// 	}
	// }

}
