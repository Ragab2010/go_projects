package main 

import (
	"fmt"
	"BMI_Project_enhance1_function/info"
	"bufio"
	"os"
	 "strings"
	 "strconv"
)

var reader = bufio.NewReader(os.Stdin)

func getUserMetrics()(weight float64 , height float64){
	weight =getUserInput(info.WeightPrompt)
	height =getUserInput(info.HeightPrompt)
	return 
}

func getUserInput(prompt string ) (value float64){
	fmt.Print(prompt)
	input , _ := reader.ReadString('\n')
	
	//remove the last character '\n'
	input = strings.Replace(input, "\n","" , -1)
	
	//convert string to float
	value , _ = strconv.ParseFloat(input, 64);
	return 

}
func claculateBMI(weight float64 , height float64) (bmi float64){
	bmi = weight / (height * height);
	return 
}	