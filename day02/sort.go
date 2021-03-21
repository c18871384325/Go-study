package main

import "fmt"
import "sort"

func main() {

	// sort  对切片不同类型进行排序 ， 需要导入sort包

	nums := []int{1, 7, 4, 2, 6}

	sort.Ints(nums)   // 通过sort ，将切片nums内的元素排序
	fmt.Println(nums)  // [1 2 4 6 7]

	names := []string{"ab", "xy", "zk", "bd"}
	sort.Strings(names)
	fmt.Println(names)   // [ab bd xy zk]

	
	// 倒序
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println(nums)  // [7 6 4 2 1] 


	// 查找
	// sort 二分查找方法
	nums = []int{1, 3, 5, 7, 9}
	fmt.Println(nums[sort.SearchInts(nums, 8)] == 8)  // false , 切片内不存在值为8 的元素
	fmt.Println(nums[sort.SearchInts(nums, 7)] == 7)  // true ， 切片内存在值为7的元素

}