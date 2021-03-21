package main

import "fmt"

func main() {	
	// key values  键值对
	// 映射，由键值对组成的数据结构，通常通过key对value进行操作
	// 无序的数据结构

	// 零值为nil
	var score map[string]float64 

	fmt.Printf("%T, %#v", score, score)    // map[string]float64, map[string]float64(nil)  

	score = map[string]float64{"huawei":78, "cisco":60, "h3c":55}  // key:value 的赋值形式
	fmt.Printf("%T, %#v", score, score)     // map[string]float64, map[string]float64(nil)map[string]float64, map[string]float64{"cisco":60, "h3c":55, "huawei":78}


	// 第二种方式通过make定义map
	// score = make(map[string]float64)   // 等同于 map[string]float64{} 


	//映射的操作，增，删，改，查

	//查找，根据key查找value
	fmt.Println(score["huawei"])
	fmt.Println(score["cisco"])
	fmt.Println(score["h3c"])

	// 判断key是否存在，存在则为true，不存在为false, 打印不存在的映射的key时，如果key不存在，默认打印的是类型的零值
	k, v := score["huawei"]   // 78 true    , 第二个变量v的值为true，表示该key在map里面存在
	fmt.Println(k, v)

	k, v = score["xyz"]    
	fmt.Println(k, v)   // 0 false  ,  v的值 为false，表示xyz的key 在map score内不存在


	// 增加
	score["jzt"] = 101   // 直接为不存在的key赋值，就是在map内增加key:value
	fmt.Println(score)   // map[cisco:60 h3c:55 huawei:78 jzt:101]


	// 删除
	delete(score, "h3c")   // 在映射 score 内删除key 为h3c 的映射
	fmt.Println(score)	  // map[cisco:60 huawei:78 jzt:101]


	// 修改
	score["huawei"] = 42   // 重新为映射内的key赋值，即为修改
	fmt.Println(score)   // map[cisco:60 huawei:42 jzt:101]

}