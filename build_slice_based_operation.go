package main
import "fmt"
func add(arr []int) int{
	var s=0
	for i:=0;i<len(arr);i++{
		s=s+arr[i]
	}
	return s
}
func min(arr []int)int{
	var m=arr[0]
	for i:=0;i<len(arr);i++{
        if arr[i]<m{
	      m=arr[i]
		}
	}
	return m
}
func max(arr []int) int{
	var m = arr[0]
	for i:=1;i<len(arr);i++{
		if arr[i]>m{
			m=arr[i]
		}
	}
	return m
}
func main(){
	arr:=[]int{9,2,8,4,5}
	sum:=add(arr)
	minimum:=min(arr)
	maximum:=max(arr)
	fmt.Println("sum is",sum)
	fmt.Println("minimum element in an array is",minimum)
	fmt.Println("maximum element in an array is",maximum)
}