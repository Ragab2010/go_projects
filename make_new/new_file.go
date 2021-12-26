package main

import "fmt"

func main() {
	str := []string{"ragab", "elsayed"}
	fmt.Println(str)
	makeStr := make([]string, 3, 5)
	fmt.Println(makeStr)
	newStr := new([]string)
	fmt.Println(*newStr)
	

}
