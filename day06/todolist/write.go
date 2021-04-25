package main

import (
	"fmt"
	"os"
	"time"
)

type Task struct {
	Id        int
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

func GenId() int {
	var id int
	for i, v := Tasks range {
			
	}
	return id	
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
	fmt.Scan(&Endtime)
	
	for i:= 0; i < 3; i++ {
		strtime, err := time.Parse(Starttime)
		if err != nil {
			fmt.Println("请输入正确的时间格式， 示例： 2021-04-01 10:10:10")
			continue
		} else {
			break
		}

	
	}
	strtime := time.Parse(Starttime)
	endtime := time.Parse(Endtime)
	add := Task{
		Id: GenId(),
		User: User,
		Name: Name,
		Starttim: &strtime,
		Endtime: endtime,
	}
	return &add

}



func main() {

	addtask := AddTask()
	fmt.Printf("%#v\n", addtask)

	start := time.Now()
	//starttime := &start
	end := time.Now().Add(10 * time.Hour)
	//endtime := &end

	Tasks := make([]*Task, 0, 20)
	

	files, err := os.Create("task.log")
	if err != nil {
		return
	}
	defer files.Close()
	for _, task := range *Tasks {
		task.WriteTo(files)
	}

}
