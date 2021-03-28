package main

import (
	"fmt"
	"log"
)

func main() {

	var numCases int

	_, err := fmt.Scanf("%d\n", &numCases)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < numCases; i++ {
		var size int
		var input string

		_, err := fmt.Scanf("%d\n", &size)
		if err != nil {
			log.Fatal(err)
		}

		_, err = fmt.Scanf("%s\n", &input)
		if err != nil {
			log.Fatal(err)
		}

		path := solve(size, input)
		fmt.Printf("Case #%d: %s\n", i+1, path)
	}
}

func solve(n int, prev string) string {
	out := ""
	for _, c := range prev {
		if c == 'S' {
			out += "E"
		} else {
			out += "S"
		}
	}
	return out
}
