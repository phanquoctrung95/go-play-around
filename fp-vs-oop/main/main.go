/**
 * @author trung
 * @created 11/09/2022
 * @project go-play-around
 */
package main

import "strings"

func main() {
	var f1 = func(f2 func()) {
		print("hello")
		f2()
	}
	var f2 = func() {
		print("hello f2")
	}
	strings.Clone()
	f1(f2)
}
