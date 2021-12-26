package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/* channels  to send int value */
	x := make(chan int)
	y := make(chan int)
	/* limit is  channel with 3 element in buffer to avoid more than 3 goRoutine work concurrent*/
	limit := make(chan int, 3)
	go generateValue(x, limit)
	go generateValue(y, limit)
	
	var a int
	var b int
	 
	select {
		case a =<-x :
			fmt.Printf("X finished first , with value:%v" , a)
		case b =<-y :
			fmt.Printf("Y finished first , with value:%v" , b)
	}
}

func generateValue(c chan int, limit chan int) {
	limit <- 1 // for make this go routine one of the 3 go routine will work buffer (limit) 
	fmt.Println("generate Value...")
	time.Sleep(time.Duration(1) * time.Second)
	result := rand.Intn(3)
	c <- result
	<-limit  // remove this rotine from the buffers (limit)

}
