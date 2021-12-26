package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

type player struct {
	speed int
	team  string
	person
}


func printdata()

func (p *player) printData(){
	fmt.Printf("the player is\nname:%v\nage:%v\nteam:%v\nspeed:%v" , p.name , p.age , p.team , p.speed)
}
func initPlayer(name_ string , age_ int , team_ string ,speed_ int) *player{
		newPlayer := player{speed:speed_ , team:team_, person:person{name:name_ , age:age_}}
		return &newPlayer
}

func main() {
//	Mido := player{speed:8 , team:"AC milan" , person:person{name:"ahmed hossam mideo" , age:22}}
//	Mido.printData()
	
	
	mido :=initPlayer("ahmed hossam mideo", 22, "AC milan", 8)
	mido.printData() 
}





