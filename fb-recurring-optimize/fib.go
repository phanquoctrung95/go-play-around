/**
 * @author trung
 * @created 25/09/2022
 * @project go-play-around
 */
package fb_recurring_optimize

func Fib(x int) int {
	if x == 0 {
		return 0
	} else if x <= 2 {
		return 1
	} else {
		return Fib(x-2) + Fib(x-1)
	}
}

func fib(x int) int {
	if x == 0 {
		return 0
	} else if x <= 2 {
		return 1
	} else {
		return fib(x-2) + fib(x-1)
	}
}
