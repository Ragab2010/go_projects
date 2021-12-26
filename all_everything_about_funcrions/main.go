package main

import (
	"fmt"
)

//type TransformFn func(int)int


func main() {
	intArr := []int{1,2,3,4,5}
	
	doubleyArr :=transformArr(&intArr, double)
	fmt.Println(intArr)
	fmt.Println("---------------------")
	fmt.Println(doubleyArr)
	
	
	/////////////////////////////////////////
	
	intArr1 := []int{1,2,3,4,5}
	intArrMore := []int{2,6,5,6,7}
	anoFn1 :=getTransformFn(&intArr)
	anoFn2 :=getTransformFn(&intArrMore)
	
	newIntArr :=transformArr(&intArr1, anoFn1)
	newIntArrMore :=transformArr(&intArrMore, anoFn2)
	fmt.Println("----------choose Ano Function-----------")
	fmt.Println(newIntArr)
	fmt.Println(newIntArrMore)
	
	fmt.Println("----------anonymous function (lamba)-----------")
	newIntArrByLamba :=transformArr(&intArr, func( val int) int{
		return 6*val
	})
	fmt.Println(newIntArrByLamba)
	
	fmt.Println("---------- closures function-----------")
	generatedFifthFun :=generateTransformFn(5)
	newIntArrFifthFun :=transformArr(&intArr,generatedFifthFun) 
	fmt.Println(newIntArrFifthFun)
	
	fmt.Println("---------- recursion function-----------")
	factorialValue  :=factorial(3)
	fmt.Println(factorialValue)
	
	fmt.Println("---------- variadic function-----------")
	variadicSum :=variadicFun_sumup(1,2,3,6,5,4,8,5,4)
	fmt.Println(variadicSum)
	splitNumberSum :=variadicFun_sumup(3, intArr...)// as variadicFun_sumup(3,  1,2,3,4,5 ) 	//intArr := []int{1,2,3,4,5}
	fmt.Println(splitNumberSum)
	
	fmt.Println("---------- defer keyword for execute inst at end of {} -----------")
	defeTestFunctionLooping(5)
	
	
	fmt.Println("---------- panic function it called to stop executing program immediately before call defer instraction  -----------")
	panicTestFunctionLooping(9)
	panicTestFunctionLooping(16) //program will stop here
	panicTestFunctionLooping(3)
	
}

func transformArr( intArr *[]int ,  transformFn func(int)int    /*transformFn TransformFn*/)  []int {
	newIntArr :=[]int{}
	for _,value := range *intArr {
		newIntArr = append(newIntArr, transformFn(value))
	}
	
	return newIntArr
	
}

func double(val int )int {
	return val *2
}

func triple(val int )int {
	return val *3
}

///////////////////////////////////////////

func getTransformFn(intArr *[]int ) func(val int )int {
	if (*intArr)[0] == 1{
		return double
	} else{
		return triple
	}
}




////////////////////////////////

func generateTransformFn(protectedValue int )func(int) int {
	return func(value int )int{
		return value *protectedValue
	}
}
/////////////////////////////////
func factorial(val int ) int{
	if val == 0{
		return 1
	}
	return val *factorial(val-1)
}


/////////////////
func variadicFun_sumup(firstNumber int , sliceList ...int) int {
	sum :=0
	for _ , number := range sliceList{
		sum = sum +number
	}
	
	return sum
}


///////////
func  defeTestFunctionLooping (number int ){
	defer func(){
		if number ==0{
		fmt.Println("loop executed well!")			
		}else{
		fmt.Println("loop dosn't  execute well!")						
		}
	}()
	
	for ;number>0 ; number--{
		fmt.Println(number)
	}
	
}


////////////////////

func panicTestFunctionLooping(number int){
		defer func(){
		if number >10{
		fmt.Println("the program will stop Now bye bye!")			
		}else{
		fmt.Println("loop execute well! without call panic function ")						
		}
	}()
	
	if number >10{
		panic("program will stop Now")
	}
	
	for ;number>0 ; number--{
		fmt.Println(number)
	}
}







