package main

import (
	"fmt"
	"encoding/gob"
	"os"
	"time"
)

type Task struct{
	Id int
	User string
	Name string
	Starttime *time.Time
	Endtime *time.Time
}

func init() {
	// 注册对象到gob管理器中
	gob.Register(&Task{})
}


func main() {

	file, err := os.Open("task.gob")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	decoding := gob.NewDecoder(file)

	tasks := []*Task{}

	decoding.Decode(&tasks)    // 注意，如果解析的gob文件内的结构体字段非包外可见，则print不到数据的   
		// Decode 解析后写入tasks 需要是一个指针类型，不为指针类型，会解析为空

	for _, task := range tasks {
		fmt.Printf("%T, %#v\n", task, task)
	}





}


