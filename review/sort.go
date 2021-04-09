package main

import (
	"fmt"
	"sort"
)

func main() {

	// 对int类型切片进行排序

	demonums := []int{1,3,7,5,9,3,4}

	sort.Sort(sort.Reverse(sort.IntSlice(demonums)))  // 对切片 从高到低的排序
	fmt.Println(demonums)
	// [9 7 5 4 3 3 1]


	sort.Ints(demonums)   // 对切片进行从低到高的数字排序
	fmt.Println(demonums)
	// [1 3 3 4 5 7 9]


	// 对string类型切片进行排序

	demostring := []string{"a", "b", "d", "c", "xx", "bb", "xx", "ww"}
	sort.Sort(sort.Reverse(sort.StringSlice(demostring)))   // 对切片内的字符串进行从高到低的排序，  a < z   A < z    
	fmt.Println(demostring)

	sort.Strings(demostring)  // 从低到高的顺序排序
	fmt.Println(demostring)


	// 二分查找   
	/* searchInts() 使用的条件为切片已经排序，示例：
	s := []int{5, 2, 6, 3, 1, 4} // 未排序的切片数据
    sort.Ints(s) //排序后的s为[1 2 3 4 5 6]
    fmt.Println(sort.SearchInts(s, 3)) //将会输出2 */    
	// 搜到到值为3的索引，并将索引返回，此处 fmt.Println返回值为2


	demosearch := []int{23,12,4,11,22,3,1,7,8,9}
	sort.Ints(demosearch)
	fmt.Println(demosearch)
	fmt.Println(demosearch[sort.SearchInts(demosearch, 8)] == 8)   //true
	fmt.Println(demosearch[sort.SearchInts(demosearch, 11)] == 11)   //true 

}