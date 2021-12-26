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
	arr := []int{1,2,3,4,5,6}
	slice := arr[:4]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice[:2]))
	arr = append(arr, 7)
	fmt.Println(arr)
	arr = arr[1:]
	fmt.Println(arr)
	
	prod := []Product{
		{"Abook " , "abook" , "great book" , "19.9"}, 
		{"A_car " , "a car" , "great Car" , "15.6"}, 

	}
	fmt.Println(prod)
	prod2 :=Product{"Apen " , "a pen " , "New pen" , "8.3"}

	prod  = append(prod, prod2)
	fmt.Println(prod)


}

