package combo_test

import (
	"fmt"

	"github.com/bobg/combo"
)

func ExamplePermutations() {
	slice := []int{1, 2, 3}
	perms := combo.Permutations(slice)
	for p := range perms {
		fmt.Println(p)
	}
	// Output:
	// [1 2 3]
	// [2 1 3]
	// [3 1 2]
	// [1 3 2]
	// [2 3 1]
	// [3 2 1]
}

func ExampleCombinations() {
	slice := []int{1, 2, 3, 4}
	combs := combo.Combinations(slice, 2)
	for c := range combs {
		fmt.Println(c)
	}
	// Output:
	// [1 2]
	// [1 3]
	// [1 4]
	// [2 3]
	// [2 4]
	// [3 4]
}

func ExampleCombinationsWithReplacement() {
	slice := []int{1, 2, 3}
	rcombs := combo.CombinationsWithReplacement(slice, 2)
	for c := range rcombs {
		fmt.Println(c)
	}
	// Output:
	// [1 1]
	// [1 2]
	// [1 3]
	// [2 2]
	// [2 3]
	// [3 3]
}
