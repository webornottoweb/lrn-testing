package sum_test

import (
	"fmt"

	"github.com/webornottoweb/lrn-testing/sum"
)

func ExampleInts() {
	s := sum.Ints(1, 2, 3, 4, 5)
	fmt.Println("sum of one to five is", s)
	// Output:
	// sum of one to five is 15
}
