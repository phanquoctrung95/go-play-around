/**
 * @author trung
 * @created 25/09/2022
 * @project go-play-around
 */
package main

import (
	"fmt"

	fbrecurringoptimize "fibo-recurring-optimize"
)

func main() {
	fmt.Printf("Fibo: %v", fbrecurringoptimize.FibMemoized(50))
	//fmt.Printf("Fibo: %v", fb_recurring_optimize.Fibonacci(50))
}
