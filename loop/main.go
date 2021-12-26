package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	//"error"
)

	
func main() {
	selectInput := getInput()
	switch selectInput {
	case 1:fmt.Printf("The Sum Reslut: %v" , calculateUpToNumber())
	case 2:fmt.Println("Under Construction!")
	case 3:fmt.Println("Under Construction!")

	}

}
func calculateUpToNumber() int {
	reader := bufio.NewReader(os.Stdin)
	checkBool := true
	sum :=0 ; 
	var inputNumber int64 
	var err error
	for checkBool {
		fmt.Print("PleaseNumber up to (X) :")
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		inputNumber, err = strconv.ParseInt(input, 0, 32)

		if err == nil {
			checkBool = false
			continue
		}
		fmt.Println("INVALID INPUT")

	}
	for i :=0 ; i <=int(inputNumber); i++{
		sum +=i
	}
	return int(sum)
}
func getInput() int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Please Make your choice .")
		fmt.Println("1) Add up the numbers of to number X ")
		fmt.Println("2) Build the factorial up to number X")
		fmt.Println("3) Sum up alist of enterd numbers .")
		fmt.Print(":")
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)
		inputNumber, err := strconv.ParseInt(input, 0, 64)

		if err == nil && inputNumber >= 0 && inputNumber <= 3 {
			return int(inputNumber)
		}
		fmt.Println("INVALID INPUT")

	}

}
