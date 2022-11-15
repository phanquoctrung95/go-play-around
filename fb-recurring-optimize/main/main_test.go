/**
 * @author trung
 * @created 25/09/2022
 * @project go-play-around
 */
package main

import (
	fbrecurringoptimize "fibo-recurring-optimize"
	"fmt"
	"testing"
)

var input = []int{8}

func BenchmarkFib(b *testing.B) {
	for _, v := range input {
		b.Run(fmt.Sprintf("fib: %v\n", v), func(b *testing.B) {
			fbrecurringoptimize.Fib(v)
		})
	}
}

func BenchmarkFibMemoize(b *testing.B) {
	for _, v := range input {
		b.Run(fmt.Sprintf("fib: %v\n", v), func(b *testing.B) {
			fbrecurringoptimize.FibMemoized(v)
		})
	}
}
