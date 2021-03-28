package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	From  int
	To    int
	Order int
	Who   string
}

type ByTimes []*Interval
type ByOrder []*Interval

func (a ByTimes) Len() int      { return len(a) }
func (a ByTimes) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByTimes) Less(i, j int) bool {
	if a[i].From == a[j].From {
		return a[i].To < a[j].To
	}
	return a[i].From < a[j].From
}

func (a ByOrder) Len() int           { return len(a) }
func (a ByOrder) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByOrder) Less(i, j int) bool { return a[i].Order < a[j].Order }

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

		var intervals []*Interval
		for i := 0; i < jobs; i++ {
			var interval Interval
			interval.Order = i
			_, err := fmt.Scanf("%d %d", &interval.From, &interval.To)
			if err != nil {
				panic(err)
			}
			intervals = append(intervals, &interval)
		}
		soln := schedule(intervals)
		fmt.Printf("Case #%d: %s\n", ncase+1, soln)
	}
}

func schedule(intervals []*Interval) string {
	sort.Sort(ByTimes(intervals))
	CFree := 0
	JFree := 0
	for _, interval := range intervals {
		if interval.From >= CFree {
			interval.Who = "C"
			CFree = interval.To
			continue
		}
		if interval.From >= JFree {
			interval.Who = "J"
			JFree = interval.To
			continue
		}
		return "IMPOSSIBLE"
	}
	sort.Sort(ByOrder(intervals))
	out := ""
	for _, interval := range intervals {
		out += interval.Who
	}
	return out
}
