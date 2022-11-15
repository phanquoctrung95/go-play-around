/**
 * @author trung
 * @created 25/09/2022
 * @project go-play-around
 */
package fb_recurring_optimize

type Memoized func(int) int

var fibMem = Memoize(fib)

func Memoize(mf Memoized) Memoized {
	cache := make(map[int]int)
	return func(key int) int {
		if val, found := cache[key]; found {
			return val
		}
		temp := mf(key)
		cache[key] = temp
		return temp
	}
}

func FibMemoized(n int) int {
	return fibMem(n)
}
