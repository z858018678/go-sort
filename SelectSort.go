package main

import "fmt"

func main()  {
	arr := []int{1,4,2,15,11,3,26,30,22,21,50}
	fmt.Println(SelectSort(arr))
}

//选择排序
//每循环一次只找最小的值
func SelectSort(arr []int)[]int{
	for i:=0;i<len(arr);i++{
		min:=i
		for j:=i+1;j<len(arr);j++ {
			if( arr[min] > arr[j] ){
				min = j
			}
		}
		arr[i],arr[min] = arr[min],arr[i]
		fmt.Println(arr)
	}
	return arr
}
//排序过程
/**
[1 4 2 15 11 3 26 30 22 21 50]
[1 2 4 15 11 3 26 30 22 21 50]
[1 2 3 15 11 4 26 30 22 21 50]
[1 2 3 4 11 15 26 30 22 21 50]
[1 2 3 4 11 15 26 30 22 21 50]
[1 2 3 4 11 15 26 30 22 21 50]
[1 2 3 4 11 15 21 30 22 26 50]
[1 2 3 4 11 15 21 22 30 26 50]
[1 2 3 4 11 15 21 22 26 30 50]
[1 2 3 4 11 15 21 22 26 30 50]
[1 2 3 4 11 15 21 22 26 30 50]
[1 2 3 4 11 15 21 22 26 30 50]
*/
