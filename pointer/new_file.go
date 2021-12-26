package main 

import (
	"fmt"
)

func main() {

	data := 100
	fmt.Printf("the Data Value before change Function call :%v\n", data)
	
	change(&data)
	fmt.Printf("the Data Value before change Function call :%v\n", data)
	
}

func change( p *int){
	*p =150
}

