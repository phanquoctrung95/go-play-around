/**
 * @author trung
 * @created 24/08/2022
 * @project go-play-around
 */
package main

import "testing"

func FuzzSum(f *testing.F) {
	testcases := [][]int64{{9999999999, 999999999}}
	for _, tc := range testcases {
		f.Add(tc[0], tc[1]) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, number1, number2 int64) {
		result := Divide(number1, number2)
		if result != number1+number2 {
			return
		}
	})
}
