package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var task = []map[string]string{
	{"id": "1", "taskname": "添加节点", "tasktime": "2020-03-19", "taskdone": "2020-03-20"},
	{"id": "2", "taskname": "删除节点", "tasktime": "2020-03-19", "taskdone": "2020-03-20"},
}

func init() {
	rand.Seed(time.Now().Unix())
}

func randString(n int) string {
	rt := make([]byte, 0, n)
	salt := []byte{'1', '2', '3', 'a', 'b', 'c', 'd'}
	for i := 0; i < n; i++ {
		rt = append(rt, salt[rand.Intn(len(salt))])
	}
	return string(salt)
}

func genid() int {
	var rt int
	for _, v := range task {
		taskid, _ := strconv.Atoi(v["id"])
		if rt < taskid {
			rt = taskid
		}
	}
	return rt + 1
}

func addtask() map[string]string {
	var taskadd map[string]string
	var taskname string
	var tasktime string
	var taskdone string
	fmt.Print("请输入任务：")
	fmt.Scan(&taskname)
	fmt.Print("请输入任务开始时间：")
	fmt.Scan(&tasktime)
	fmt.Print("请输入任务结束时间：")
	fmt.Scan(&taskdone)
	taskadd = map[string]string{"taskname": taskname, "tasktime": tasktime, "taskdone": taskdone}
	taskadd["id"] = strconv.Itoa(genid())
	return taskadd
}

func printtask(task map[string]string) {
	fmt.Println(strings.Repeat("-", 20))
	fmt.Println("ID", task["id"])
	fmt.Println("任务名", task["taskname"])
	fmt.Println("tasktime", task["tasktime"])
	fmt.Println("taskdone", task["taskdone"])
}

func search() {
	var staskname string
	fmt.Print("请输入查询的任务名：")
	fmt.Scan(&staskname)
	for _, v := range task {
		if staskname == "all" || strings.Contains(v["taskname"], staskname) {
			printtask(v)
		}
	}
}

func delete() {
	var deletetask string
	fmt.Print("请输入删除的任务：")
	fmt.Scan(&deletetask)
	for i, v := range task {
		if v["taskname"] == deletetask {
			copy(task[i:], task[i+1:])
			task = task[:len(task)-1]
		}
	}
}

func edit() {
	var edittask string
	fmt.Print("请输入修改的任务名: ")
	fmt.Scan(&edittask)
	for i, v := range task {
		if v["taskname"] == edittask {
			var edittaskname string
			var edittasktime string
			var edittaskdone string
			fmt.Print("请输入任务：")
			fmt.Scan(&edittaskname)
			fmt.Print("请输入任务开始时间：")
			fmt.Scan(&edittasktime)
			fmt.Print("请输入任务结束时间：")
			fmt.Scan(&edittaskdone)
			task[i]["taskname"] = edittaskname
			task[i]["tasktime"] = edittasktime
			task[i]["taskdone"] = edittaskdone
		}
	}
}

func passString() int {
	passsalt := randString(5)
	inputpass := "cheng2018!"
	passsha256 := sha256.New()
	passsha256.Write([]byte(inputpass))
	passsha256.Write([]byte(passsalt))
	endpassword := hex.EncodeToString(passsha256.Sum(nil))

	var inpass string
	fmt.Print("请输入密码：")
	fmt.Scan(&inpass)
	inpass256 := sha256.New()
	inpass256.Write([]byte(inpass))
	inpass256.Write([]byte(passsalt))
	inverity := hex.EncodeToString(inpass256.Sum(nil))

	if inverity == endpassword {
		return 0
	} else {
		return 1
	}
}

func main() {
	
BREAK:
	for i := 0; i < 3; i++ {
		passvalue := passString()
		if passvalue == 0 {
			for true {
				var input string
				fmt.Print("请输入操作：")
				fmt.Scan(&input)
				switch input {
				case "add":
					task = append(task, addtask())
				case "search":
					search()
				case "delete":
					delete()
				case "edit":
					edit()
				case "exit":
					break BREAK
				default:
					fmt.Println("imput error, please imput add/search/exit")
				}
			}
		} else {
			fmt.Println("密码错误，请重新输入！")
		}
	}
}
