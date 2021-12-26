package main 

import (
	"fmt"
)
func main() {
	fristNumber := 5
	secondNumber := 10.1
	sliceNumber :=[]int {1,2,3,4,5}
	other :=string("welcome")
	
	doubleFristNumber :=double(fristNumber) 
	doubleSecondNumber :=double(secondNumber) 
	doublesliceNumber :=double(sliceNumber) 
	doubleother :=double(other) 
	
	fmt.Println(doubleFristNumber)
	fmt.Println(doubleSecondNumber)
	fmt.Println(doublesliceNumber)
	fmt.Println(doubleother)

	
	

}

func double ( data interface{}) interface{}{
	switch val :=data.(type){
		case int:
			return val *2
		case int64:
			return val *2
		case float64:
			return val *2
		case []int:
			 newSlice :=append(val, val...)
			 return newSlice
		default:
			return ""
	}
}
