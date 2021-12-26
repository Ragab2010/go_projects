package main 

import (
	"fmt"
)

type customInt int

func (p customInt) sumWithReceiver( num customInt) customInt{
	return p + num
}


type mapping []map[int]string

func main() {
	var  value customInt = 5
	sumValue := value.sumWithReceiver(5)
	fmt.Println(sumValue)
	mapString :=mapping{
		{1:"ramy"},
		{5:"sayed"},
	}
	fmt.Println(mapString)
}

