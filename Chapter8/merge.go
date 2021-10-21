package sample

import (
	"math"
	"runtime"
	"sync"
)

// Example merge sort to demonstrate benchmark pitfalls
// Lots of slicing - Not such a good algorithm for go

func merge(l, r []int) []int {
	ret := make([]int, 0, len(l)+len(r))

	for {
		switch {
		case len(l) == 0:
			return append(ret, r...)

		case len(r) == 0:
			return append(ret, l...)

		case l[0] <= r[0]:
			ret = append(ret, l[0])
			l = l[1:]

		default:
			ret = append(ret, r[0])
			r = r[1:]
		}
	}
}

func single(n []int) []int {
	if len(n) <= 1 {
		return n
	}

	i := len(n) / 2

	l := single(n[:i])
	r := single(n[i:])

	return merge(l, r)
}

func unlimited(n []int) []int {
	if len(n) <= 1 {
		return n
	}

	i := len(n) / 2

	var l, r []int

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		l = unlimited(n[:i])
		wg.Done()
	}()

	go func() {
		r = unlimited(n[i:])
		wg.Done()
	}()

	wg.Wait()
	return merge(l, r)
}

func numCPU(n []int, lvl int) []int {
	if len(n) <= 1 {
		return n
	}

	i := len(n) / 2

	var l, r []int

	maxLevel := int(math.Log2(float64(runtime.GOMAXPROCS(0))))

	if lvl <= maxLevel {
		lvl++

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			l = numCPU(n[:i], lvl)
			wg.Done()
		}()

		go func() {
			r = numCPU(n[i:], lvl)
			wg.Done()
		}()

		wg.Wait()
		return merge(l, r)
	}

	l = numCPU(n[:i], lvl)
	r = numCPU(n[i:], lvl)

	return merge(l, r)
}
