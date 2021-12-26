package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("welcome in egypt")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please Enter you Age :")
	age, _ := reader.ReadString('\n')
	age = strings.Replace(age, "\n", "", -1)
	ageValue, err := strconv.ParseInt(age, 0, 64)
	//
	fmt.Println(ageValue)
	fmt.Println(err)
}
