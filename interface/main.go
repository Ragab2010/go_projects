package main

import (
	"fmt"
)

type printer interface {
	print()
}

type person struct {
	id         int
	personName string
	age        int
}

type car struct {
	id      int
	carName string
	manYear int
}


func (p person ) print(){
	fmt.Printf("id:%v\npersonName:%v\nage:%v\n" , p.id,p.personName , p.age)
	fmt.Println("-------------------------------------")
	
}
func (c car ) print(){
	fmt.Printf("id:%v\ncarName:%v\nmanYear:%v\n" , c.id,c.carName , c.manYear)
	fmt.Println("-------------------------------------")
	
}

func main() {
	ahmed :=person{4 , "Ahemd" , 26}
	//ahmed.print()
	
	toyota :=car{6 , "Toyota" , 2013}
	//toyota.print()
	print(ahmed)
	print(toyota)

}

func print(p printer){
	p.print()
}
