package main

import "fmt"

func main(){
	arr := []int{1,4,2,15,11,3,26,30,22,21,50}
	fmt.Println(BubbleSorts(arr))
}



func BubbleSort(arr []int) []int {
	var tmp int = 0
	for i:=0;i<len(arr);i++ {
		for j:=0;j<len(arr);j++ {
			if arr[i]<arr[j] {
				tmp = arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
			}
			fmt.Println(arr)
		}
	}
	return arr
}
//第一次循环找第一小的，第二次循环找第二小的
func BubbleSorts (arr []int)[]int{
	for i:=0;i<len(arr);i++{
		for j:=i+1;j<len(arr);j++{
			if(arr[i]>arr[j]){
				arr[i],arr[j] = arr[j],arr[i]
			}
			fmt.Println(arr)
		}
	}
	return  arr
}


