package main 

import (
	"BMI_Project_by_struct_encap/methods"


)


func main() {

	var bmiForPerson methods.PersonalInfo
	
//output infromation 
	bmiForPerson.OutputInfromation()

//input infromation
	bmiForPerson.InputData() 

//calculate  BMI
	bmiForPerson.CalculateBMI()

//
//output BMI
	bmiForPerson.PrintData()



}

