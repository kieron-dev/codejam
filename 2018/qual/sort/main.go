package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	numCasesLine, _ := reader.ReadString('\n')
	numCases, _ := strconv.Atoi(strings.TrimSpace(numCasesLine))

	for i := 1; i <= numCases; i++ {
		listLenLine, _ := reader.ReadString('\n')
		listLen, _ := strconv.Atoi(strings.TrimSpace(listLenLine))

		listStr, _ := reader.ReadString('\n')
		listStrs := strings.Split(strings.TrimSpace(listStr), " ")
		var listEven, listOdd []int
		for j := 0; j < listLen; j++ {
			n, _ := strconv.Atoi(strings.TrimSpace(listStrs[j]))
			if j%2 == 0 {
				listEven = append(listEven, n)
			} else {
				listOdd = append(listOdd, n)
			}
		}
		ok, pos := QuickSolve(listEven, listOdd)
		var answer string
		if ok {
			answer = "OK"
		} else {
			answer = fmt.Sprintf("%d", pos)
		}
		fmt.Printf("Case #%d: %s\n", i, answer)
	}
}

func QuickSolve(listEven, listOdd []int) (ok bool, pos int) {
	sort.Ints(listEven)
	sort.Ints(listOdd)
	for i := 0; i < len(listEven)+len(listOdd)-1; i++ {
		var cur, next int
		if i%2 == 0 {
			cur = listEven[i/2]
			next = listOdd[i/2]
		} else {
			cur = listOdd[i/2]
			next = listEven[i/2+1]
		}
		if cur > next {
			return false, i
		}
	}
	return true, 0
}
