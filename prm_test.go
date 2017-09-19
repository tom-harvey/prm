package prm

import (
	"testing"
)

// const maxFactorial = 20 TODO can this be seen ?

// calculate factorial to check the table-driven version
func altFactorial(n int) uint64 {
	result := uint64(1)
	if n > maxFactorial {
		return 0 // TODO sync with overflow behavior when decided
	}
	for u := uint64(2); u <= uint64(n); u++ {
		result *= u
	}
	return result
}

var pcount uint64
var posint = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k",
	"l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v"}

// count the number of permutations generated,
// TODO and check for uniqueness
func checkP(perm []int, items []string) {
	pcount++
}

func TestFactorial(t *testing.T) {
	for i := 0; i <= maxFactorial; i++ {
		f1 := Factorial(i)
		f2 := altFactorial(i)
		if f1 != f2 {
			t.Errorf("Factorial(%v): expected %v, got %v\n", i, f2, f1)
		}
	}
}

func TestGenPermutations(t *testing.T) {
	for i := 0; i <= 10; i++ {
		pcount = 0
		GenPermutations(i, posint, checkP)
		f := Factorial(i)
		if pcount != Factorial(i) {
			t.Errorf("Generated %v permutations but expected %v\n", pcount, f)
		}
	}
}
