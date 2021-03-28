package main

import "fmt"

func main() {
	var testCases int
	fmt.Scanf("%d", &testCases)

	for i := 1; i <= testCases; i++ {
		var reqSquares int
		fmt.Scanf("%d", &reqSquares)
		Solve(reqSquares)
	}
}

func Solve(r int) {
	var w int
	if r%3 == 0 {
		w = r / 3
	} else {
		w = r/3 + 1
	}
	squares := [1000][1000]bool{}

	for x := 2; x < w; x++ {
		for !(squares[1][x-1] && squares[2][x-1] && squares[3][x-1]) {
			fmt.Printf("2 %d\n", x)
			var ax, ay int
			fmt.Scanf("%d %d", &ax, &ay)
			if ax == 0 && ay == 0 {
				break
			}
			squares[ax][ay] = true
		}
	}
	for !(squares[1][w-1] && squares[2][w-1] && squares[3][w-1] && squares[1][w] && squares[2][w] && squares[3][w]) {
		fmt.Printf("2 %d\n", w-1)
		var ax, ay int
		fmt.Scanf("%d %d", &ax, &ay)
		if ax == 0 && ay == 0 {
			break
		}
		squares[ax][ay] = true
	}
}
