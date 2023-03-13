package main

import (
	"fmt"
	"github.com/liondadev/fibn"
	"time"
)

func main() {
	var i int64 = 0
	for {
		fmt.Println(fibn.Calc(i))
		i++
		time.Sleep(100 * time.Millisecond)
	}
}
