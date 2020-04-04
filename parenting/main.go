package main

import (
	"fmt"
)

type Interval struct {
	From int
	To   int
}

var People []string = []string{"C", "J"}

func main() {
	var numCases int
	_, err := fmt.Scanf("%d", &numCases)
	if err != nil {
		panic(err)
	}

	for ncase := 0; ncase < numCases; ncase++ {
		var jobs int
		_, err := fmt.Scanf("%d", &jobs)
		if err != nil {
			panic(err)
		}

		var intervals []Interval
		for i := 0; i < jobs; i++ {
			var interval Interval
			_, err := fmt.Scanf("%d %d", &interval.From, &interval.To)
			if err != nil {
				panic(err)
			}
			intervals = append(intervals, interval)
		}
		soln := schedule(intervals, map[string][]Interval{})
		fmt.Printf("Case #%d: %s\n", ncase+1, soln)
	}
}

func schedule(intervals []Interval, assigned map[string][]Interval) string {
	if len(intervals) == 0 {
		return ""
	}
	first := intervals[0]

	for _, person := range People {
		if isFree(first, assigned[person]) {
			rest := intervals[1:]
			newAssigned := assign(person, first, assigned)
			soln := schedule(rest, newAssigned)
			if soln != "IMPOSSIBLE" {
				return person + soln
			}
		}
	}

	return "IMPOSSIBLE"
}

func isFree(period Interval, assignments []Interval) bool {
	for _, assignment := range assignments {
		if (period.From >= assignment.From && period.To <= assignment.To) ||
			(period.From <= assignment.From && period.To > assignment.From) ||
			(period.From < assignment.To && period.To >= assignment.To) {
			return false
		}
	}
	return true
}

func assign(person string, period Interval, assigned map[string][]Interval) map[string][]Interval {
	ret := map[string][]Interval{}
	for _, p := range People {
		newSlice := make([]Interval, len(assigned[p]))
		copy(newSlice, assigned[p])
		if p == person {
			newSlice = append(newSlice, period)
		}
		ret[p] = newSlice
	}
	return ret
}
