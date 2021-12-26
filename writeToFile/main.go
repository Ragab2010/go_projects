package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Product struct {
	id          string
	title       string
	description string
	price       string
}

func main() {
	
	prod := getProtuct()
	prod.store()

}

func getProtuct() *Product{
	reader :=bufio.NewReader(os.Stdin)
	pro := Product{
	id:readInput(reader, "protuct ID: ") , 
	title:readInput(reader, "protuct Title: "),
	description:readInput(reader, "protuct Description: "),
	price:readInput(reader, "protuct Price: ")}
	return &pro
	
}

func readInput(reader *bufio.Reader , promptText string )  string {
	fmt.Print(promptText)
	input , _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	return input
}

func (pro *Product)store(){
	file ,_:= os.Create("src/writeToFile/"+pro.id + ".txt")
	
	stringFile := fmt.Sprintf(
	"ID:%v\nTitle:%v\nDescription:%v\nrice:%v\n" , 
	pro.id, 
	pro.title,
	pro.description,
	pro.price,
	)
	
	file.WriteString(stringFile)
	file.Close()
	
	
}

