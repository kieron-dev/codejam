package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	log.SetOutput(os.Stderr)
	var numCases, listLength, questionLimit int

	if _, err := fmt.Scanf("%d %d %d", &numCases, &listLength, &questionLimit); err != nil {
		log.Panic(err)
	}

	for n := 1; n <= numCases; n++ {
		log.Printf("Solving: case %d, len %d, limit %d", n, listLength, questionLimit)
		if err := solve(listLength, questionLimit); err != nil {
			log.Panic(err)
		}
	}
}

func solve(l, q int) error {
	// 1 < 2 by assumption
	orderedList := []int{1, 2}

	for n := 3; n <= l; n++ {
		orderedList = insert(n, orderedList, 0, n-2)
	}

	s := []string{}
	for _, n := range orderedList {
		s = append(s, strconv.Itoa(n))
	}

	log.Printf("solution %s", s)
	fmt.Println(strings.Join(s, " "))

	var ok int
	fmt.Scanf("%d", &ok)
	if ok != 1 {
		log.Panic("incorrect answer")
	}

	return nil
}

func insert(n int, list []int, left, right int) []int {
	log.Printf("insert(n: %d, list: %v, left: %d, right: %d)\n", n, list, left, right)
	if right-left < 0 {
		log.Panic("right < left??")
	}

	if right == left {
		if left == 0 {
			right = len(list) - 1
			m := ask(list[left], list[right], n)
			switch m {
			case list[left]:
				return []int{n, list[left]}
			case n:
				return []int{list[left], n}
			}
		}

		left = 0
		m := ask(list[left], list[right], n)
		switch m {
		case list[right]:
			return []int{list[right], n}
		case n:
			return []int{n, list[right]}
		}
	}

	if right-left == 1 {
		m := ask(list[left], list[right], n)

		switch m {
		case list[left]:
			return []int{n, list[left], list[right]}

		case list[right]:
			return []int{list[left], list[right], n}
		case n:
			return []int{list[left], n, list[right]}
		default:
			log.Panicf("ask returned %d, not in %v", m, list)
		}
	}

	midpoint := left + (right-left)/2
	m := ask(list[midpoint], list[midpoint+1], n)

	newList := []int{}
	switch m {
	case n:
		log.Println("case n in middle")
		newList = append(newList, list[left:midpoint+1]...)
		newList = append(newList, n)
		newList = append(newList, list[midpoint+1:right+1]...)

	case list[midpoint]:
		log.Println("case n in left")
		if midpoint == left {
			newList = append(newList, n)
			newList = append(newList, list[left:right+1]...)
		} else {
			newList = append(newList, insert(n, list, left, midpoint-1)...)
			newList = append(newList, list[midpoint:right+1]...)
		}
	case list[midpoint+1]:
		log.Println("case n in right")
		if midpoint+1 == right {
			newList = append(newList, list[left:right+1]...)
			newList = append(newList, n)
		} else {
			newList = append(newList, list[left:midpoint+2]...)
			newList = append(newList, insert(n, list, midpoint+2, right)...)
		}
	}

	// 5 3 1 4 2
	// 1 5 3 6 4 2
	// 1 7 5 6 4 2 3

	return newList
}

func ask(a, b, c int) int {
	fmt.Printf("%d %d %d\n", a, b, c)
	log.Printf("ask %d %d %d: ", a, b, c)

	var resp int
	if _, err := fmt.Scanf("%d", &resp); err != nil {
		log.Panic(err)
	}

	if resp == -1 {
		log.Panic("ask got a -1")
	}

	log.Printf("-> %d\n", resp)
	return resp
}
