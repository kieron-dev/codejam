package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	debugLog, err := os.Create("debug.log")
	if err != nil {
		panic(err)
	}
	defer debugLog.Close()

	log.SetOutput(debugLog)
	log.Printf("start")

	var numCases, lenBits int

	_, err = fmt.Scanf("%d %d", &numCases, &lenBits)
	if err != nil {
		panic(err)
	}
	log.Printf("numCases = %+v\n", numCases)
	log.Printf("lenBits = %+v\n", lenBits)

	for i := 0; i < numCases; i++ {
		log.Printf("\nCase #%d\n", i+1)
		count = 0
		doIt(lenBits)
	}
}

var count int

func clear(arr []int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = -1
	}
}

func doIt(l int) {
	arr := make([]int, l)
	clear(arr)

	var sameLow, sameHigh int
	var diffLow, diffHigh int

	haveSame := false
	haveDiff := false

	sames := map[int]bool{}

	var pos int
	for pos = 0; pos < l/2; pos++ {
		if !(haveSame && haveDiff) {
			if count%10 == 0 {
				clear(arr)
			}
			l1 := send(pos)
			arr[pos] = l1
			h1 := send(l - pos - 1)
			arr[l-pos-1] = h1
			if l1 == h1 {
				sames[pos] = true
				haveSame = true
				sameLow = pos
				sameHigh = l - pos - 1
			}
			if l1 != h1 {
				sames[pos] = false
				haveDiff = true
				diffLow = pos
				diffHigh = l - pos - 1
			}
		}
	}

	log.Printf("haveSame = %+v\n", haveSame)
	log.Printf("haveDiff = %+v\n", haveDiff)

	for count%10 > 0 {
		getNext(arr, haveSame, haveDiff, sames)
	}

	for count < 150 {
		curSame := -1
		if haveSame {
			curSame = send(sameLow)
		}
		curDiff := -1
		if haveDiff {
			curDiff = send(diffLow)
		}

		mutateArray(arr, arr[sameLow], arr[diffLow], curSame, curDiff)
		if haveSame {
			arr[sameLow] = curSame
			arr[sameHigh] = curSame
		}
		if haveDiff {
			arr[diffLow] = curDiff
			arr[diffHigh] = 1 - curDiff
		}

		for count%10 != 0 {
			getNext(arr, haveSame, haveDiff, sames)
			if isDone(arr) {
				res := ""
				for i := 0; i < l; i++ {
					res += strconv.Itoa(arr[i])
				}
				ok := sendRes(res)
				if !ok {
					panic("wrong answer")
				}
				return
			}
		}
	}
	panic("out of guesses")
}

func isDone(arr []int) bool {
	for _, v := range arr {
		if v == -1 {
			return false
		}
	}
	return true
}

func getNext(arr []int, haveSame, haveDiff bool, sames map[int]bool) {
	for i, v := range arr {
		if v == -1 {
			log.Printf("getNext: %d\n", i)
			arr[i] = send(i)
			if same, ok := sames[i]; ok {
				if same {
					arr[len(arr)-i-1] = arr[i]
				} else {
					arr[len(arr)-i-1] = 1 - arr[i]
				}
			}
			return
		}
	}
}

func mutateArray(arr []int, prevSame, prevDiff, curSame, curDiff int) {
	if curSame == -1 && prevDiff != curDiff {
		reverse(arr)
		return
	}
	if curDiff == -1 && prevSame != curSame {
		complement(arr)
		return
	}
	if prevSame == curSame {
		if prevDiff == curDiff {
			return
		}
		reverse(arr)
		return
	}
	if prevDiff == curDiff {
		reverse(arr)
	}
	complement(arr)
}

func reverse(arr []int) {
	l := len(arr)
	for i := 0; i < l/2; i++ {
		arr[i], arr[l-i-1] = arr[l-i-1], arr[i]
	}
}

func complement(arr []int) {
	for i, v := range arr {
		if v == -1 {
			continue
		}
		arr[i] = 1 - v
	}
}

func send(n int) int {
	count++
	fmt.Printf("%d\n", n+1)
	var res int
	_, err := fmt.Scanf("%d", &res)
	if err != nil {
		panic(err)
	}
	log.Printf("#%d: arr[%d] = %d\n", count, n, res)
	return res
}

func sendRes(res string) bool {
	log.Printf("sending %s\n", res)
	fmt.Println(res)
	var ok string
	_, err := fmt.Scanf("%s", &ok)
	if err != nil {
		panic(err)
	}
	return ok == "Y"
}
