package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	lim := 20
	rand.Seed(time.Now().UnixNano())
	fmt.Println(1)
	fmt.Println(lim)
	for i := 0; i < lim; i++ {
		fmt.Printf("%d ", rand.Intn(10000000))
	}
	fmt.Println()
}
