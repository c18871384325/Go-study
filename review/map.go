package main

import (
	"fmt"
)

func main() {
	// map 是一组同类型的 key : value 的键值对存储,元素长度不固定 , 其他特性与数组相同
	var demo map[string]string      // 定义map 名为 demo ，key ,value 的类型均为 string

	demo = map[string]string{"huawei": "95", "cisco": "100", "h3c": "15"}

	fmt.Println(demo)     //map[cisco:100 h3c:15 huawei:95]

	demo["tencent"] = "230"   // 直接为map的key赋值
	fmt.Println(demo)
}