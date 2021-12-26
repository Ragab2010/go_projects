package main 

import (
	"fmt"
)

func main() {
	data :=[]int{1,2,3,4,5,6}
	doubleData :=doubleSlice(data...)
	fmt.Println(doubleData)
	
	
	doubleData2 :=doubleSlice(3,6,5,6,9,8,4)
	fmt.Println(doubleData2)
}

func doubleSlice(data ...int) interface{}{
	newData :=append(data, data...)
	return newData
}