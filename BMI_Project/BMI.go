package main 

import (
	"fmt"
	"bufio"
	"os"
	 "BMI_Project/info"
	 "strings"
	 "strconv"
)

var reader = bufio.NewReader(os.Stdin)
func main() {
//output infromation 
	fmt.Println(info.MainTitle)
	fmt.Println(info.Separator)

//input infromation 
	fmt.Print(info.WeightPrompt)
	 weightInput , _ := reader.ReadString('\n')
	
	fmt.Print(info.HeightPrompt)
	 heightInput , _ := reader.ReadString('\n')
	
	//remove the last character '\n'
	weightInput = strings.Replace(weightInput, "\n","" , -1)
	heightInput = strings.Replace(heightInput, "\n","" , -1)
	
	//convert string to float
	weight , _ := strconv.ParseFloat(weightInput, 64);
	height , _ := strconv.ParseFloat(heightInput, 64);

	fmt.Println(weight)
	fmt.Println(height)
	

//calculate  BMI
	bmi := weight / (height * height);

//output BMI
		fmt.Printf("the BMI= %.2f ",bmi)
	

}

