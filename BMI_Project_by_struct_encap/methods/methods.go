package methods 

import (
	"fmt"
	"BMI_Project_by_struct_encap/info"
	"bufio"
	"os"
	"strings"
	"strconv"
)

var Reader = bufio.NewReader(os.Stdin)

type PersonalInfo struct {
	weight float64 
	height float64
	bmi	   float64
} 
 

func ( p *PersonalInfo)InputData(){
	fmt.Print(info.WeightPrompt)
	 weightInput , _ := Reader.ReadString('\n')
	
	fmt.Print(info.HeightPrompt)
	 heightInput , _ := Reader.ReadString('\n')
	 
	 //remove the last character '\n'
	weightInput = strings.Replace(weightInput, "\n","" , -1)
	heightInput = strings.Replace(heightInput, "\n","" , -1)
	 p.ConvertDataToFloat(weightInput, heightInput)

	 
} 


func ( p *PersonalInfo)ConvertDataToFloat(weightInput string , heightInput string){
	//convert string to float
	p.weight , _ = strconv.ParseFloat(weightInput, 64);
	p.height , _ = strconv.ParseFloat(heightInput, 64);
	
	

} 

func ( p *PersonalInfo) CalculateBMI(){
	//calculate  BMI
	p.bmi = p.weight / (p.height * p.height);

	
}


func ( p *PersonalInfo)PrintData(){
	
		fmt.Printf("the BMI= %.2f ",p.bmi)
} 


func ( p *PersonalInfo)OutputInfromation(){
	//output infromation 
	fmt.Println(info.MainTitle)
	fmt.Println(info.Separator)

}