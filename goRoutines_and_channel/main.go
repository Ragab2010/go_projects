package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*channel is channel  to send int value */
	channel := make(chan int)
	/* limit is  channel with 3 element in buffer to avoid more than 3 goRoutine work concurrent*/
	limit := make(chan int, 3)
	go generateValue(channel, limit)
	go generateValue(channel, limit)
	go generateValue(channel, limit)
	go generateValue(channel, limit)
	var sum int
	var i int = 0
	for num := range channel {
		sum += num
		i++
		if i == 4 {
			close(channel)
			close(limit)
		}
	}
	fmt.Println(sum)

}

func generateValue(c chan int, limit chan int) {
	limit <- 1 // for make this go routine one of the 3 go routine will work buffer (limit) 
	fmt.Println("generate Value...")
	time.Sleep(time.Duration(1) * time.Second)
	result := rand.Intn(3)
	c <- result
	<-limit  // remove this rotine from the buffers (limit)

}
