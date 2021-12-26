package main 

import (
	"fmt"
)


type Product struct {
	id          string
	title       string
	description string
	price       string
}

func main() {
	productMap := map[string]Product{
		"CAR" : {"A_car " , "a car" , "great Car" , "15.6"},
	}
	fmt.Println(productMap)
	fmt.Println(productMap["CAR"])
	
	productMap["PEN"] =Product{"A_pen" , "a_pen ", "new Pen " , "15.9"}
	fmt.Println(productMap)
	
	delete(productMap,"CAR")
	fmt.Println(productMap)

	

}

