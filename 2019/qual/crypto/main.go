package main

import (
	"fmt"
	"log"
	"math/big"
	"sort"
)

func main() {

	var numCases int

	_, err := fmt.Scanf("%d\n", &numCases)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < numCases; i++ {
		var primeLim int
		var listLen int

		_, err := fmt.Scanf("%d %d\n", &primeLim, &listLen)
		if err != nil {
			panic(err)
		}

		nums := []*big.Int{}
		for j := 0; j < listLen; j++ {
			var s string
			_, err := fmt.Scanf("%s", &s)
			if err != nil {
				log.Panic(err)
			}
			var num big.Int
			num.SetString(s, 10)
			nums = append(nums, &num)
		}

		out := solve(primeLim, nums)
		fmt.Printf("Case #%d: %s\n", i+1, out)
	}
}

type Pair struct {
	pos int
	val *big.Int
}

type ByVal []Pair

func (a ByVal) Len() int {
	return len(a)
}
func (a ByVal) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByVal) Less(i, j int) bool {
	return a[i].val.Cmp(a[j].val) < 0
}
func solve(n int, nums []*big.Int) string {
	var primes []Pair

	i := 0
	for nums[i] == nums[i+1] {
		i++
	}
	p := gcd(nums[i], nums[i+1])

	primes = append(primes, Pair{pos: i + 1, val: p})

	prevPrime := big.NewInt(0)
	prevPrime.Set(p)
	for j := i; j >= 0; j-- {
		nextPrime := big.NewInt(0)
		nextPrime.Div(nums[j], prevPrime)
		primes = append(primes, Pair{pos: j, val: nextPrime})
		fmt.Printf("nextPrime = %+v\n", nextPrime)
		prevPrime.Set(nextPrime)
	}
	for _, p := range primes {
		fmt.Printf("p.val = %+v\n", p.val)
	}

	prevPrime = p
	for j := i + 2; j < len(nums); j++ {
		nextPrime := big.NewInt(0)
		nextPrime.Div(nums[j], prevPrime)
		primes = append(primes, Pair{pos: j, val: nextPrime})
		prevPrime.Set(nextPrime)
	}
	for _, p := range primes {
		fmt.Printf("p.val = %+v\n", p.val)
	}

	sort.Sort(ByVal(primes))

	out := make([]byte, len(primes))
	prev := big.NewInt(0)
	prevVal := byte('A') - 1
	for _, p := range primes {
		if prev.Cmp(p.val) != 0 {
			prev = p.val
			prevVal = prevVal + 1
		}
		out[p.pos] = prevVal
	}

	return string(out)
}

func gcd(a, b *big.Int) *big.Int {
	z := big.NewInt(0)
	if a.Cmp(b) < 1 {
		return gcd(b, a)
	}
	m := z.Mod(a, b)
	if m.Cmp(big.NewInt(0)) == 0 {
		return b
	}
	return gcd(b, m)
}
