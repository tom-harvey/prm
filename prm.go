package prm

import (
	"fmt"
	"strings"
)

// just remember all the factorials that can fit in unsigned 64 bits
const maxFactorial = 20

var ftable = [maxFactorial + 1]uint64{1, 1, 2, 6, 24, 120, 720, 5040, 40320,
	362880, 3628800, 39916800, 479001600, 6227020800, 87178291200,
	1307674368000, 20922789888000, 355687428096000, 6402373705728000,
	121645100408832000, 2432902008176640000}

func Factorial(n int) uint64 {
	if n <= maxFactorial {
		return ftable[n]
	}
	return 0 // TODO panic or complain
}

// generate cartesian product in lexicographic order
// assumes 'items' are listed in lexicographic order
func GenProduct(n int, items []string, gen func([]int, []string)) {
	b := len(items)
	top := b - 1
	right_column := n - 1
	o := make([]int, n)
	gen(o, items)
	for {
		for c := right_column; ; c-- {
			o[c]++
			if o[c] <= top {
				break
			}
			// overflow on this column
			if c == 0 {
				return
			}
			o[c] = 0
		}
		gen(o, items)
	}
}

// generatePermutations() by nonrecursive Heap's method
// practical numbers of permutation elements are small (<20)
func GenPermutations(n int, items []string, gen func([]int, []string)) {
	c := make([]int, n)
	a := make([]int, n)
	for i, _ := range a {
		a[i] = i
	}

	gen(a, items)
	for i := 0; i < n; {
		if c[i] < i {
			if i&1 == 0 {
				a[0], a[i] = a[i], a[0]
			} else {
				a[c[i]], a[i] = a[i], a[c[i]]
			}
			gen(a, items)
			c[i]++
			i = 0
		} else {
			c[i] = 0
			i++
		}
	}
}

func showPermutation(perm []byte, items []string) {
	var names []string
	for _, p := range perm {
		names = append(names, items[p])
	}
	fmt.Println(strings.Join(names, " "))
}
