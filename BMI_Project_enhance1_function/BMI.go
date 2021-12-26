package main 

import (

	 "BMI_Project_enhance1_function/info"

)


func main() {
//output infromation 
	info.PrintInfromation()

//input infromation 
	weight , height :=getUserMetrics()

	

//calculate  BMI
	bmi := claculateBMI(weight , height)


//output BMI
	printBMI(bmi)
	

}

